// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	"context"
	"crypto/tls"
	"encoding/json"
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
		sendBuf:  make(chan []byte, 1),
		org:      ctx,
		creds:    &creds,
		callback: callback,
		retry:    true,
	}
	conn.ctx, conn.ctxCancel = context.WithCancel(ctx)

	u := url.URL{Scheme: "wss", Host: creds.Host, Path: creds.Channel}
	conn.configStr = u.String()

	klog.V(3).Infof("NewWebSocketClient Succeeded\n")
	klog.V(6).Infof("NewWebSocketClient LEAVE\n")
	return &conn, nil
}

func (conn *WebSocketClient) Connect() *websocket.Conn {
	return conn.ConnectWithRetry(defaultConnectRetry)
}

func (conn *WebSocketClient) AttemptReconnect(retries int64) *websocket.Conn {
	conn.retry = true
	return conn.ConnectWithRetry(retries)
}

func (conn *WebSocketClient) ConnectWithRetry(retries int64) *websocket.Conn {
	// we explicitly stopped and should not attempt to reconnect
	if !conn.retry {
		klog.V(5).Infof("This connection has been terminated. Please either call with AttemptReconnect or create a new Client object using NewWebSocketClient.")
		return nil
	}

	// if the connection is good, return it
	// otherwise, attempt reconnect
	if conn.wsconn != nil {
		select {
		case <-conn.ctx.Done():
			// continue through to reconnect by recreating the wsconn object
			klog.V(6).Infof("Connection is broken. Will attempt reconnect.")
			conn.ctx, conn.ctxCancel = context.WithCancel(conn.org)
		default:
			klog.V(7).Infof("Connection is good. Return object.")
			return conn.wsconn
		}
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
				klog.V(4).Infof("Connect() RESTORE Header: %s = %s\n", k, v)
				myHeader.Add(k, v)
			}
		}
	}

	// sets the API key
	myHeader.Set("X-API-KEY", conn.creds.AccessKey)

	// attempt to establish connection
	i := int64(0)
	for {
		if retries != connectionRetryInfinite && i >= retries {
			klog.V(1).Infof("Connect timeout... exiting!\n")
			break
		}

		// delay on subsequent calls
		if i > 0 {
			klog.V(4).Infof("Sleep for retry #%d...\n", i)
			time.Sleep(time.Second * time.Duration(defaultDelayBetweenRetry))
		}

		i++

		// create new connection
		ws, _, err := dialer.DialContext(conn.ctx, conn.configStr, myHeader)
		if err != nil {
			klog.V(1).Infof("Cannot connect to websocket: %s\n", conn.configStr)
			continue
		}

		// set the object to allow threads to function
		klog.V(4).Infof("WebSocket Connection Successful!")
		conn.wsconn = ws
		conn.retry = true

		// kick off threads
		go conn.listen()
		go conn.ping()

		return conn.wsconn
	}

	return nil
}

func (conn *WebSocketClient) listen() {
	klog.V(6).Infof("WebSocketClient::listen ENTER\n")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-conn.ctx.Done():
			return
		case <-ticker.C:
			for {
				ws := conn.Connect()
				if ws == nil {
					klog.V(3).Infof("WebSocketClient::listen Connection is not valid\n")
					break
				}

				msgType, byMsg, err := ws.ReadMessage()
				if err != nil {
					klog.V(3).Infof("Cannot read websocket message. Err: %v\n", err)
					break
				}

				if conn.callback != nil {
					conn.callback.Message(byMsg)
				} else {
					klog.V(3).Infof("WebSocketClient msg recv (type %d): %s\n", msgType, string(byMsg))
				}
			}
		}
	}

	klog.V(6).Infof("WebSocketClient::listen LEAVE\n")
}

// Write struct to the websocket server
func (conn *WebSocketClient) WriteBinary(byData []byte) error {
	// doing a write, need to lock
	conn.mu.Lock()
	defer conn.mu.Unlock()

	ws := conn.Connect()
	if ws == nil {
		klog.V(1).Infof("WebSocketClient::WriteBinary Connection is not valid\n")
		return ErrInvalidConnection
	}

	if err := ws.WriteMessage(
		websocket.BinaryMessage,
		byData,
	); err != nil {
		klog.V(1).Infof("WebSocketClient::WriteBinary WriteMessage failed. Err: %v\n", err)
		return err
	}

	klog.V(7).Infof("WriteBinary Successful\n")
	klog.V(7).Infof("WriteBinary payload:\nData: %x\n", byData)

	return nil
}

// WriteJSON struct to the websocket server
func (conn *WebSocketClient) WriteJSON(payload interface{}) error {
	// doing a write, need to lock
	conn.mu.Lock()
	defer conn.mu.Unlock()

	ws := conn.Connect()
	if ws == nil {
		klog.V(1).Infof("WebSocketClient::WriteJSON Connection is not valid\n")
		return ErrInvalidConnection
	}

	dataStruct, err := json.Marshal(payload)
	if err != nil {
		klog.V(1).Infof("WebSocketClient::WriteJSON json.Marshal failed. Err: %v\n", err)
		return err
	}

	if err := ws.WriteMessage(
		websocket.TextMessage,
		dataStruct,
	); err != nil {
		klog.V(1).Infof("WebSocketClient::WriteJSON WriteMessage failed. Err: %v\n", err)
		return err
	}

	klog.V(6).Infof("WriteJSON Successful\n")
	klog.V(7).Infof("WriteJSON payload:\nData: %s\n", string(dataStruct))

	return nil
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

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) Stop() {
	klog.V(3).Infof("WebSocketClient::Stop Stopping...\n")
	conn.retry = false
	conn.ctxCancel()
	conn.closeWs()
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) closeWs() {
	klog.V(3).Infof("WebSocketClient::closeWs closing channels...\n")

	// doing a write, need to lockx
	conn.mu.Lock()
	defer conn.mu.Unlock()

	if conn.wsconn != nil {
		err := conn.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			klog.V(1).Infof("Failed to send CloseNormalClosure. Err: %v\n", err)
		}
		time.Sleep(time.Millisecond * time.Duration(100)) // allow time for server to register closure
		conn.wsconn.Close()
	}
}

func (conn *WebSocketClient) ping() {
	klog.V(6).Infof("WebSocketClient::ping ENTER\n")

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-conn.ctx.Done():
			return
		case <-ticker.C:
			klog.V(6).Infof("Starting ping...")

			ws := conn.Connect()
			if ws == nil {
				klog.V(1).Infof("WebSocketClient::ping Connect is not valid\n")
				break
			}

			// doing a write, need to lock
			conn.mu.Lock()
			klog.V(6).Infof("Sending ping... need reply in %d\n", (pingPeriod / 2))
			err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pingPeriod/2))
			conn.mu.Unlock()

			if err != nil {
				klog.V(1).Infof("WebSocketClient::ping failed\n")
				conn.closeWs()
			} else {
				klog.V(4).Infof("Ping sent!")
			}
		}
	}

	klog.V(6).Infof("WebSocketClient::ping LEAVE\n")
}
