// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package stream

import (
	"context"
	"sync"

	"github.com/dvonthenen/websocket"
)

type WebSocketMessageCallback interface {
	Message(byMsg []byte) error
}

// Credentials is the input needed to login to the Symbl.ai platform
type Credentials struct {
	Host            string `validate:"required"`
	Channel         string `validate:"required"`
	AccessKey       string `validate:"required"`
	RedirectService bool
	SkipServerAuth  bool
}

// BinaryData format for sending audio
type EncapsulatedMessage struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
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

	stopListen chan struct{}
	stopPing   chan struct{}
}
