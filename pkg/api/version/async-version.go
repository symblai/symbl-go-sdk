// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package version

import (
	"fmt"
)

const (
	AsyncAPIVersion string = "v1"

	// processing
	ProcessAudioURI string = "https://api.symbl.ai/%s/process/audio?name=%s"
	ProcessURLURI   string = "https://api.symbl.ai/%s/process/audio/url"

	// job status
	JobStatusURI string = "https://api.symbl.ai/%s/job/%s"

	// intelligence
	TopicsURI      string = "https://api.symbl.ai/%s/conversations/%s/topics?parentRefs=true&sentiment=true"
	QuestionsURI   string = "https://api.symbl.ai/%s/conversations/%s/questions?parentRefs=true&sentiment=true"
	FollowUpsURI   string = "https://api.symbl.ai/%s/conversations/%s/follow-ups?parentRefs=true&sentiment=true"
	EntitiesURI    string = "https://api.symbl.ai/%s/conversations/%s/entities?parentRefs=true&sentiment=true"
	ActionItemsURI string = "https://api.symbl.ai/%s/conversations/%s/action-items?parentRefs=true&sentiment=true"
	MessagesURI    string = "https://api.symbl.ai/%s/conversations/%s/messages?parentRefs=true&sentiment=true"
	AnalyticsURI   string = "https://api.symbl.ai/%s/conversations/%s/analytics?parentRefs=true&sentiment=true"
	TrackersURI    string = "https://api.symbl.ai/%s/conversations/%s/trackers?parentRefs=true&sentiment=true"

	// bookmarks
	BookmarksURI     string = "https://api.symbl.ai/%s/conversations/%s/bookmarks"
	BookmarksByIdURI string = "https://api.symbl.ai/%s/conversations/%s/bookmarks/%s"
	// BookmarkSummaryURI      string = "https://api.symbl.ai/%s/conversations/%s/bookmarks/%s/summary"
	// SummariesOfBookmarksURI string = "https://api.symbl.ai/%s/conversations/%s/bookmarks-summary"

	// summary ui
	SummaryURI string = "https://api.symbl.ai/%s/conversations/%s/experiences"

	// Conversations
	ConversationsURI string = "https://api.symbl.ai/%s/conversations"
	ConversationURI  string = "https://api.symbl.ai/%s/conversations/%s"

	// Members
	MembersURI  string = "https://api.symbl.ai/%s/conversations/%s/members"
	MemberURI   string = "https://api.symbl.ai/%s/conversations/%s/members/%s"
	SpeakersURI string = "https://api.symbl.ai/%s/conversations/%s/speakers"
)

func GetAsyncAPI(URI string, args ...interface{}) string {
	return fmt.Sprintf(URI, append([]interface{}{AsyncAPIVersion}, args...)...)
}
