// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	"context"
	"sync"

	"github.com/dvonthenen/websocket"
)

// WebSocketMessageCallback is a callback used to write a message on websocket without
// exposing the entire struct to the user
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

// WebSocketClient return websocket client connection
type WebSocketClient struct {
	configStr string
	sendBuf   chan []byte

	org       context.Context
	ctx       context.Context
	ctxCancel context.CancelFunc

	mu     sync.RWMutex
	wsconn *websocket.Conn
	retry  bool

	creds    *Credentials
	callback WebSocketMessageCallback
}
