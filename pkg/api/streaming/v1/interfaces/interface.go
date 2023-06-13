// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Defines everything that makes up the Streaming API interface
*/
package interfaces

type InsightCallback interface {
	// InitializedConversation signals a conversation started
	InitializedConversation(im *InitializationMessage) error

	// RecognitionResultMessage signals a word is recognized
	RecognitionResultMessage(rr *RecognitionResult) error

	// MessageResponseMessage signals a sentence is completed
	MessageResponseMessage(mr *MessageResponse) error

	// InsightResponseMessage signals a question, action item, follow up is detected within the conversation
	InsightResponseMessage(ir *InsightResponse) error

	// TopicResponseMessage signals a topic is detected within the conversation
	TopicResponseMessage(tr *TopicResponse) error

	// TrackerResponseMessage signals a tracker is detected within the conversation
	TrackerResponseMessage(tr *TrackerResponse) error

	// EntityResponseMessage signals an entity is detected within the conversation
	EntityResponseMessage(er *EntityResponse) error

	// InitializedConversation signals a conversation ended
	TeardownConversation(tm *TeardownMessage) error

	// UserDefinedMessage signals this is a user defined message
	UserDefinedMessage(data []byte) error

	// UnhandledMessage signals an unknown message has been received. this usually indicates a new
	// capability, message, or insight is available on the platform
	UnhandledMessage(byMsg []byte) error
}
