// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Defines everything that makes up the Async API interface
*/
package interfaces

/*
	Input parameters for Nebula Async API calls
*/
type Conversation struct {
	Text string `json:"text,omitempty" validate:"required"`
}

type Prompt struct {
	Instruction  string       `json:"instruction,omitempty" validate:"required"`
	Conversation Conversation `json:"conversation,omitempty" validate:"required"`
}

type AskNebulaRequest struct {
	Prompt       Prompt  `json:"prompt,omitempty" validate:"required"`
	ReturnScores bool    `json:"return_scores,omitempty"` // TODO: IMPORTANT: Always set to false. Yields "Internal Server Error"
	MaxNewTokens int     `json:"max_new_tokens,omitempty"`
	TopK         int     `json:"top_k,omitempty"`
	PenaltyAlpha float64 `json:"penalty_alpha,omitempty"`
}

/*
	Output parameters for Nebula Async API calls
*/
type Output struct {
	Text   string      `json:"text,omitempty"`
	Scores interface{} `json:"scores,omitempty"` // TODO: IMPORTANT: Always empty. Yields "Internal Server Error". See above.
}

type Stats struct {
	InstructionTokens  int `json:"instruction_tokens,omitempty"`
	ConversationTokens int `json:"conversation_tokens,omitempty"`
	OutputTokens       int `json:"output_tokens,omitempty"`
	OutputWords        int `json:"output_words,omitempty"`
	OutputSentences    int `json:"output_sentences,omitempty"`
	TotalInputTokens   int `json:"total_input_tokens,omitempty"`
	TotalTokens        int `json:"total_tokens,omitempty"`
}

type AskNebulaResponse struct {
	Model  string `json:"model,omitempty"`
	Output Output `json:"output,omitempty"`
	Stats  Stats  `json:"stats,omitempty"`
}
