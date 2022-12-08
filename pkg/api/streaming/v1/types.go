// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

// MessageType is the header to bootstrap you way unmarshalling other messages
/*
	Example:
	{
		"type": "message",
		"message": {
			"type": "started_listening"
		}
	}
*/
type MessageType struct {
	Type string `json:"type"`
}

// SybmlMessageType is the header to bootstrap you way unmarshalling other messages
/*
	Example:
	{
		"type": "message",
		"message": {
			"type": "started_listening"
		}
	}
*/
type SybmlMessageType struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
	} `json:"message"`
}

// SymblError when mt.Type == "error"
type SymblError struct {
	Type    string `json:"type"`
	Details string `json:"details"`
	Message string `json:"message"`
}
