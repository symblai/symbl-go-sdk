// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
This package handles the versioning in the API both async and streaming
*/
package version

import (
	"fmt"
)

const (
	AsyncAPIVersion string = "v1"

	// processing audio
	ProcessAudioURI    string = "https://api.symbl.ai/%s/process/audio?name=%s"
	ProcessAudioURLURI string = "https://api.symbl.ai/%s/process/audio/url"

	// processing video
	ProcessVideoURI    string = "https://api.symbl.ai/%s/process/video?name=%s"
	ProcessVideoURLURI string = "https://api.symbl.ai/%s/process/video/url"

	// processing text
	ProcessTextURI       string = "https://api.symbl.ai/%s/process/text"
	ProcessAppendTextURI string = "https://api.symbl.ai/%s/process/text/%s"

	// job status
	JobStatusURI string = "https://api.symbl.ai/%s/job/%s"

	// intelligence
	TopicsURI      string = "https://api.symbl.ai/%s/conversations/%s/topics"
	QuestionsURI   string = "https://api.symbl.ai/%s/conversations/%s/questions"
	FollowUpsURI   string = "https://api.symbl.ai/%s/conversations/%s/follow-ups"
	EntitiesURI    string = "https://api.symbl.ai/%s/conversations/%s/entities"
	ActionItemsURI string = "https://api.symbl.ai/%s/conversations/%s/action-items"
	MessagesURI    string = "https://api.symbl.ai/%s/conversations/%s/messages"
	AnalyticsURI   string = "https://api.symbl.ai/%s/conversations/%s/analytics"
	TrackersURI    string = "https://api.symbl.ai/%s/conversations/%s/trackers"
	TranscriptURI  string = "https://api.symbl.ai/%s/conversations/%s/transcript"

	// bookmarks
	BookmarksURI            string = "https://api.symbl.ai/%s/conversations/%s/bookmarks"
	BookmarksByIdURI        string = "https://api.symbl.ai/%s/conversations/%s/bookmarks/%s"
	BookmarkSummaryURI      string = "https://api.symbl.ai/%s/conversations/%s/bookmarks/%s/summary"
	SummariesOfBookmarksURI string = "https://api.symbl.ai/%s/conversations/%s/bookmarks-summary"

	// summary ui
	SummaryURI string = "https://api.symbl.ai/%s/conversations/%s/summary"

	// Insights Ui
	InsightsListUiURI    string = "https://api.symbl.ai/%s/conversations/experiences/insights/list?includeCallScore=true"
	InsightsDetailsUiURI string = "https://api.symbl.ai/%s/conversations/experiences/insights/details/%s?includeCallScore=true"
	UpdateMediaURI string = "https://api.symbl.ai/%s/conversations/%s/experience/url"

	// Conversations
	ConversationsURI string = "https://api.symbl.ai/%s/conversations"
	ConversationURI  string = "https://api.symbl.ai/%s/conversations/%s"

	// Members
	MembersURI  string = "https://api.symbl.ai/%s/conversations/%s/members"
	MemberURI   string = "https://api.symbl.ai/%s/conversations/%s/members/%s"
	SpeakersURI string = "https://api.symbl.ai/%s/conversations/%s/speakers"

	// Call Score
	CallScoreStatusURI string = "https://api.symbl.ai/%s/conversations/%s/callscore/status"
	InsightStatusURI   string = "https://api.symbl.ai/%s/conversations/%s/lm-insights/status"
	CallScoreURI       string = "https://api.symbl.ai/%s/conversations/%s/callscore"
)

func GetAsyncAPI(URI string, args ...interface{}) string {
	return fmt.Sprintf(URI, append([]interface{}{AsyncAPIVersion}, args...)...)
}
