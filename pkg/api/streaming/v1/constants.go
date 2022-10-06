// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import "errors"

var (
	// ErrInvalidMessageType invalid message type
	ErrInvalidMessageType = errors.New("invalid message type")

	// ErrUserCallbackNotDefined user callback object not defined
	ErrUserCallbackNotDefined = errors.New("user callback object not defined")
)

// Handshake Related
const (
	SymblPlatformHost string = "api.symbl.ai"

	TypeRequestStart string = "start_request"
	TypeRequestStop  string = "stop_request"
)

// Message Types
const (
	MessageTypeInitListening        string = "started_listening"
	MessageTypeInitConversation     string = "conversation_created"
	MessageTypeInitRecognition      string = "recognition_started"
	MessageTypeSessionModified      string = "session_modified"
	MessageTypeTeardownConversation string = "conversation_completed"
	MessageTypeTeardownRecognition  string = "recognition_stopped"

	MessageTypeError   string = "error"
	MessageTypeMessage string = "message"
)
