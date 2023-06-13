// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Streaming package for processing real-time conversations
*/
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
