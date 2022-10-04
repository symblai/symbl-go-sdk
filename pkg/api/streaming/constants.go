// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

const (
	SymblPlatformHost string = "api.symbl.ai"

	TypeRequestStart string = "start_request"
	TypeBinaryData   string = "binary"
)

const (
	MessageTypeInitListening    string = "started_listening"
	MessageTypeInitConversation string = "conversation_created"
	MessageTypeInitRecognition  string = "recognition_started"
	MessageTypeError            string = "error"
)
