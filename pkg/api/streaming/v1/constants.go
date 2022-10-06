// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import "errors"

var (
	// ErrInvalidMessageType invalid message type
	ErrInvalidMessageType = errors.New("invalid message type")
)

const (
	SymblPlatformHost string = "api.symbl.ai"

	TypeRequestStart string = "start_request"
	TypeRequestStop  string = "stop_request"
)

const (
	MessageTypeInitListening    string = "started_listening"
	MessageTypeInitConversation string = "conversation_created"
	MessageTypeInitRecognition  string = "recognition_started"

	MessageTypeError   string = "error"
	MessageTypeMessage string = "message"
)
