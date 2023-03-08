// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package stream

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dvonthenen/websocket"
	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

// Send pings to peer with this period
const (
	pingPeriod = 30 * time.Second
)

// NewWebSocketClient create new websocket connection
func NewWebSocketClient(ctx context.Context, creds Credentials, callback WebSocketMessageCallback) (*WebSocketClient, error) {
	klog.V(6).Infof("NewWebSocketClient ENTER\n")

	if callback == nil {
		klog.V(3).Infof("NewWebSocketClient callback is nil. Will not process messages. Will print only.\n")
	}

	// validate input
	v := validator.New()
	err := v.Struct(creds)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("NewWebSocketClient validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("NewWebSocketClient LEAVE\n")
		return nil, err
	}

	// init
	conn := WebSocketClient{
		sendBuf:    make(chan []byte, 1),
		creds:      &creds,
		callback:   callback,
		stopListen: make(chan struct{}),
		stopPing:   make(chan struct{}),
	}
	conn.ctx, conn.ctxCancel = context.WithCancel(ctx)

	u := url.URL{Scheme: "wss", Host: creds.Host, Path: creds.Channel}
	conn.configStr = u.String()

	klog.V(3).Infof("NewWebSocketClient Succeeded\n")
	klog.V(6).Infof("NewWebSocketClient LEAVE\n")
	return &conn, nil
}

func (conn *WebSocketClient) Connect() *websocket.Conn {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.wsconn != nil {
		return conn.wsconn
	}

	// TODO: Disable the Hostname validation for now
	dialer := websocket.Dialer{
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
		RedirectService:  conn.creds.RedirectService,
		SkipServerAuth:   conn.creds.SkipServerAuth,
	}

	// access key for Symbl Platfom
	myHeader := http.Header{}

	// restore application options to HTTP header
	if headers, ok := conn.ctx.Value(interfaces.HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				klog.V(5).Infof("Connect() RESTORE Header: %s = %s\n", k, v)
				myHeader.Add(k, v)
			}
		}
	}

	// sets the API key
	myHeader.Set("X-API-KEY", conn.creds.AccessKey)

	// wait for handshake
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		select {
		case <-conn.ctx.Done():
			return nil
		default:
			ws, _, err := dialer.DialContext(conn.ctx, conn.configStr, myHeader)
			if err != nil {
				klog.V(1).Infof("Cannot connect to websocket: %s\n", conn.configStr)
				continue
			}

			// set the object to allow threads to function
			conn.wsconn = ws

			// kick off threads
			go conn.listen(conn.stopListen)
			go conn.listenWrite()
			go conn.ping(conn.stopPing)

			return conn.wsconn
		}
	}
}

func (conn *WebSocketClient) listen(stopChan chan struct{}) {
	klog.V(6).Infof("WebSocketClient::listen ENTER\n")
	klog.V(3).Infof("listen for the messages: %s\n", conn.configStr)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-stopChan:
			return
		case <-ticker.C:
			for {
				ws := conn.Connect()
				if ws == nil {
					klog.V(1).Infof("WebSocketClient::listen Connect is not valid\n")
					klog.V(6).Infof("WebSocketClient::listen LEAVE\n")
					return
				}
				msgType, bytMsg, err := ws.ReadMessage()
				if err != nil {
					klog.V(3).Infof("Cannot read websocket message. Err: %v\n", err)
					klog.V(6).Infof("WebSocketClient::listen LEAVE\n")
					return
				}

				if conn.callback != nil {
					conn.callback.Message(bytMsg)
				} else {
					klog.V(3).Infof("WebSocketClient msg recv (type %d): %s\n", msgType, string(bytMsg))
				}
			}
		}
	}
}

// Write struct to the websocket server
func (conn *WebSocketClient) WriteBinary(byData []byte) error {
	ed := &EncapsulatedMessage{
		Type: websocket.BinaryMessage,
		Data: byData,
	}
	data, err := json.Marshal(ed)
	if err != nil {
		klog.V(1).Infof("WebSocketClient::Write json.Marshal failed. Err: %v\n", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	for {
		select {
		case conn.sendBuf <- data:
			return nil
		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		}
	}
}

// WriteJSON struct to the websocket server
func (conn *WebSocketClient) WriteJSON(payload interface{}) error {
	dataStruct, err := json.Marshal(payload)
	if err != nil {
		klog.V(1).Infof("WebSocketClient::Write json.Marshal failed. Err: %v\n", err)
		return err
	}

	ed := &EncapsulatedMessage{
		Type: websocket.TextMessage,
		Data: dataStruct,
	}
	data, err := json.Marshal(ed)
	if err != nil {
		klog.V(1).Infof("WebSocketClient::Write json.Marshal failed. Err: %v\n", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	for {
		select {
		case conn.sendBuf <- data:
			return nil
		case <-ctx.Done():
			return fmt.Errorf("context canceled")
		}
	}
}

func (conn *WebSocketClient) Write(p []byte) (int, error) {
	byteLen := len(p)
	err := conn.WriteBinary(p)
	if err != nil {
		klog.V(1).Infof("WebSocketClient::WriteBinary failed. Err: %v\n", err)
		return 0, err
	}
	return byteLen, nil
}

func (conn *WebSocketClient) listenWrite() {
	klog.V(6).Infof("WebSocketClient::listen ENTER\n")

	for data := range conn.sendBuf {
		ws := conn.Connect()
		if ws == nil {
			klog.V(1).Infof("WebSocketClient::listenWrite Connect is not valid\n")
			return
		}

		var em EncapsulatedMessage
		err := json.Unmarshal([]byte(data), &em)
		if err != nil {
			klog.V(1).Infof("WebSocketClient::listenWrite json.Unmarshal failed. Err: %v\n", err)
			continue
		}

		if err := ws.WriteMessage(
			em.Type,
			em.Data,
		); err != nil {
			klog.V(1).Infof("WebSocketClient::listenWrite Write failed. Err: %v\n", err)
		}
	}

	klog.V(6).Infof("WebSocketClient::listen LEAVE\n")
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) Stop() {
	klog.V(3).Infof("WebSocketClient::Stop Stopping...\n")
	conn.ctxCancel() // TODO: is this really needed?
	conn.closeWs()

	// stop threads
	close(conn.stopListen)
	<-conn.stopListen
	close(conn.stopPing)
	<-conn.stopPing
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) closeWs() {
	klog.V(3).Infof("WebSocketClient::closeWs closing channels...\n")

	conn.mu.Lock()

	if conn.wsconn != nil {
		conn.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		conn.wsconn.Close()
	}
	conn.wsconn = nil

	conn.mu.Unlock()
}

func (conn *WebSocketClient) ping(stopChan chan struct{}) {
	klog.V(6).Infof("WebSocketClient::ping ENTER\n")

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-stopChan:
			return
		case <-ticker.C:
			ws := conn.Connect()
			if ws == nil {
				continue
			}
			if err := conn.wsconn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pingPeriod/2)); err != nil {
				conn.closeWs()
			}
		}
	}

	klog.V(6).Infof("WebSocketClient::ping LEAVE\n")
}
