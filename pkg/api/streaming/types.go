// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

// MessageType is the header to bootstrap you way unmarshalling other messages
/*
	Example: {"type":"message", ...}
*/
type MessageType struct {
	Type string `json:"type"`
}

// SybmlMessageType is the header to bootstrap you way unmarshalling other messages
/*
	Example: {"type":"message","message":{"type":"started_listening"}}
*/
type SybmlMessageType struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
	} `json:"message"`
}

// SymblInitializationMessage the init message when mt.Type == "conversation_created"
/*
	Example: {"type":"message","message":{"type":"conversation_created","data":{"conversationId":"5751229838262272"}}}
*/
type SymblInitializationMessage struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
		Data struct {
			ConversationID string `json:"conversationId"`
		} `json:"data"`
	} `json:"message"`
}

// SymblError when mt.Type == "error"
type SymblError struct {
	Type    string `json:"type"`
	Details string `json:"details"`
	Message string `json:"message"`
}
