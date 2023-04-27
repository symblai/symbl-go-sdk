// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package interfaces

type InsightCallback interface {
	InitializedConversation(ci *InitializationMessage) error
	MessageResult(mr *MessageResult) error
	QuestionResult(qr *QuestionResult) error
	FollowUpResult(fr *FollowUpResult) error
	ActionItemResult(air *ActionItemResult) error
	TopicResult(tr *TopicResult) error
	SummaryResult(tr *SummaryResult) error
	TrackerResult(tr *TrackerResult) error
	EntityResult(er *EntityResult) error
	SummaryUIResult(er *SummaryUIResult) error
	TeardownConversation(ct *TeardownMessage) error

	// TODO: need to make the case for this callback. Violates immutable rule.
	// MembersResult(er *MembersResult) error
	// BookmarksResult(er *BookmarksResult) error
	// BookmarkSummaryResult(er *BookmarkSummaryResult) error
	// BookmarksSummaryResult(er *BookmarksSummaryResult) error
}
