// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Defines everything that makes up the Async API interface
*/
package interfaces

type InsightCallback interface {
	// InitializedConversation is called when an async conversation is starting to be processed
	InitializedConversation(ci *InitializationMessage) error

	// MessageResult is called for processing messages for transcription
	MessageResult(mr *MessageResult) error

	// QuestionResult is called for processing questions
	QuestionResult(qr *QuestionResult) error

	// FollowUpResult is called for processing follow ups
	FollowUpResult(fr *FollowUpResult) error

	// ActionItemResult is called for processing action items
	ActionItemResult(air *ActionItemResult) error

	// TopicResult is called for processing topics
	TopicResult(tr *TopicResult) error

	// TrackerResult is called for processing trackers
	TrackerResult(tr *TrackerResult) error

	// EntityResult is called for processing entities
	EntityResult(er *EntityResult) error

	// TeardownConversation is called when all insights are processed and we are done with the given conversation
	TeardownConversation(ct *TeardownMessage) error

	// TODO: need to make the case for this callback. Violates immutable rule.
	// MembersResult(er *MembersResult) error
	// SummaryResult(tr *SummaryResult) error
	// BookmarksResult(er *BookmarksResult) error
	// BookmarkSummaryResult(er *BookmarkSummaryResult) error
	// BookmarksSummaryResult(er *BookmarksSummaryResult) error
	// SummaryUIResult(er *SummaryUIResult) error
}
