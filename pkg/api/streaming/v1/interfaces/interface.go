// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package interfaces

type InsightCallback interface {
	InitializedConversation(im *InitializationMessage) error
	RecognitionResultMessage(rr *RecognitionResult) error
	MessageResponseMessage(mr *MessageResponse) error
	InsightResponseMessage(ir *InsightResponse) error
	TopicResponseMessage(tr *TopicResponse) error
	TrackerResponseMessage(tr *TrackerResponse) error
	EntityResponseMessage(er *EntityResponse) error
	TeardownConversation(tm *TeardownMessage) error
	UserDefinedMessage(data []byte) error
	UnhandledMessage(byMsg []byte) error
}
