// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package stream

type WebSocketMessageCallback interface {
	Message(byMsg []byte) error
}

// Credentials is the input needed to login to the Symbl.ai platform
type Credentials struct {
	Host      string `validate:"required"`
	Channel   string `validate:"required"`
	AccessKey string `validate:"required"`
}

// BinaryData format for sending audio
type EncapsulatedMessage struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
}
