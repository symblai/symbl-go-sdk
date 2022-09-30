// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"
)

// Send pings to peer with this period
const (
	pingPeriod = 30 * time.Second
)

type WebSocketMessageCallback interface {
	Message(byMsg []byte) error
}

// Credentials is the input needed to login to the Symbl.ai platform
type Credentials struct {
	Host      string `validate:"required"`
	Channel   string `validate:"required"`
	AccessKey string `validate:"required"`
}

// WebSocketClient return websocket client connection
type WebSocketClient struct {
	configStr string
	sendBuf   chan []byte
	ctx       context.Context
	ctxCancel context.CancelFunc

	mu     sync.RWMutex
	wsconn *websocket.Conn

	creds    *Credentials
	callback WebSocketMessageCallback
}

// NewWebSocketClient create new websocket connection
func NewWebSocketClient(creds Credentials, callback WebSocketMessageCallback) (*WebSocketClient, error) {
	klog.V(6).Infof("NewWebSocketClient ENTER\n")

	if callback == nil {
		klog.V(2).Infof("NewWebSocketClient callback is nil. Will not process messages. Will print only.\n")
	}

	// validate input
	v := validator.New()
	err := v.Struct(creds)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.Errorf("NewWebSocketClient validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("NewWebSocketClient LEAVE\n")
		return nil, err
	}

	// init
	conn := WebSocketClient{
		sendBuf:  make(chan []byte, 1),
		creds:    &creds,
		callback: callback,
	}
	conn.ctx, conn.ctxCancel = context.WithCancel(context.Background())

	u := url.URL{Scheme: "wss", Host: creds.Host, Path: creds.Channel}
	conn.configStr = u.String()

	go conn.listen()
	go conn.listenWrite()
	go conn.ping()

	klog.V(2).Infof("NewWebSocketClient Succeeded\n")
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
	}

	// access key for Symbl Platfom
	myHeader := http.Header{}
	myHeader.Set("X-API-KEY", conn.creds.AccessKey)

	// wait for handshake
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		select {
		case <-conn.ctx.Done():
			return nil
		default:
			ws, _, err := dialer.Dial(conn.configStr, myHeader)
			if err != nil {
				klog.V(2).Infof("Cannot connect to websocket: %s\n", conn.configStr)
				continue
			}

			conn.wsconn = ws
			return conn.wsconn
		}
	}
}

func (conn *WebSocketClient) listen() {
	klog.V(6).Infof("WebSocketClient::listen ENTER\n")
	klog.V(2).Infof("listen for the messages: %s\n", conn.configStr)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-conn.ctx.Done():
			return
		case <-ticker.C:
			for {
				ws := conn.Connect()
				if ws == nil {
					klog.V(2).Infof("WebSocketClient::listen Connect is not valid\n")
					klog.V(6).Infof("WebSocketClient::listen LEAVE\n")
					return
				}
				msgType, bytMsg, err := ws.ReadMessage()
				if err != nil {
					klog.V(2).Infof("Cannot read websocket message. Err: %v\n", err)
					conn.closeWs()
					break
				}

				if conn.callback != nil {
					conn.callback.Message(bytMsg)
				} else {
					// klog.V(2).Infof("WebSocketClient msg string (type %d): %x\n", msgType, bytMsg)
					klog.V(2).Infof("WebSocketClient msg recv (type %d): %s\n", msgType, string(bytMsg))
				}
			}
		}
	}
}

// Write data to the websocket server
func (conn *WebSocketClient) Write(payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		klog.V(2).Infof("WebSocketClient::Write json.Marshal failed. Err: %v\n", err)
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

func (conn *WebSocketClient) listenWrite() {
	for data := range conn.sendBuf {
		ws := conn.Connect()
		if ws == nil {
			klog.V(2).Infof("WebSocketClient::listenWrite Connect is not valid\n")
			continue
		}

		if err := ws.WriteMessage(
			websocket.TextMessage,
			data,
		); err != nil {
			klog.V(2).Infof("WebSocketClient::listenWrite Write failed. Err: %v\n", err)
		}
		klog.V(2).Infof("WebSocketClient::listenWrite Write succeeded.\nSend: %s\n", data)
	}
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) Stop() {
	klog.V(2).Infof("WebSocketClient::Stop Stopping...\n")
	conn.ctxCancel()
	conn.closeWs()
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) closeWs() {
	klog.V(2).Infof("WebSocketClient::closeWs closing channels...\n")

	conn.mu.Lock()
	if conn.wsconn != nil {
		conn.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		conn.wsconn.Close()
		conn.wsconn = nil
	}
	conn.mu.Unlock()
}

func (conn *WebSocketClient) ping() {
	klog.V(2).Infof("WebSocketClient::ping started...\n")

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			ws := conn.Connect()
			if ws == nil {
				continue
			}
			if err := conn.wsconn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pingPeriod/2)); err != nil {
				conn.closeWs()
			}
		case <-conn.ctx.Done():
			return
		}
	}
}
