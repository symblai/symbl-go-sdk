// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package interfaces

/*
	Internal package messages
*/
type ConversationInitialization struct {
	ConversationID string `json:"conversationId"`
}

type ConversationTeardown struct {
	ConversationID string `json:"conversationId"`
}

/*
	Shared definitions
*/
type User struct {
	Name   string `json:"name,omitempty" validate:"required"`
	UserID string `json:"userId,omitempty" validate:"required"`
	Email  string `json:"email,omitempty" validate:"required"`
}

type From struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Duration struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type Metric struct {
	Type    string  `json:"type,omitempty"`
	Percent float64 `json:"percent,omitempty"`
	Seconds float64 `json:"seconds,omitempty"`
}

type Sentiment struct {
	Polarity struct {
		Score float64 `json:"score,omitempty"`
	} `json:"polarity,omitempty"`
	Suggested string `json:"suggested,omitempty"`
}

type MessageRef struct {
	ID        string `json:"id,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Text      string `json:"text,omitempty"`
	Offset    int    `json:"offset,omitempty"`
}

type ParentRef struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type InsightRef struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type TrackerMatch struct {
	Type        string       `json:"type,omitempty"`
	Value       string       `json:"value,omitempty"`
	MessageRefs []MessageRef `json:"messageRefs,omitempty"`
	InsightRefs []InsightRef `json:"insightRefs,omitempty"`
}

type EntityMatch struct {
	DetectedValue string       `json:"detectedValue,omitempty"`
	MessageRefs   []MessageRef `json:"messageRefs,omitempty"`
}

type EntityInsight struct {
	Type   string `json:"type,omitempty"`
	Text   string `json:"text,omitempty"`
	Offset int    `json:"offset,omitempty"`
	End    string `json:"end,omitempty"`
}

type Entity struct {
	Type     string        `json:"type,omitempty"`
	SubType  string        `json:"subType,omitempty"`
	Category string        `json:"category,omitempty"`
	Matches  []EntityMatch `json:"matches,omitempty"`
}

/*
	When exercising the API and description is blank...

	HTTP Code: 400
	{
		"message":"\"description\" is not allowed to be empty"
	}
*/
type Bookmark struct {
	ID              string       `json:"id,omitempty"`
	Label           string       `json:"label,omitempty" validate:"required"`
	Description     string       `json:"description,omitempty" validate:"required"` // please see note above
	User            User         `json:"user,omitempty" validate:"required"`
	BeginTimeOffset int          `json:"beginTimeOffset,omitempty" validate:"required"`
	Duration        int          `json:"duration,omitempty" validate:"required"`
	MessageRefs     []MessageRef `json:"messageRefs,omitempty" validate:"required"`
}

type Topic struct {
	Text       string      `json:"text,omitempty"`
	Type       string      `json:"type,omitempty"`
	Score      float64     `json:"score,omitempty"`
	MessageIds []string    `json:"messageIds,omitempty"`
	Sentiment  Sentiment   `json:"sentiment,omitempty"`
	ParentRefs []ParentRef `json:"parentRefs,omitempty"`
}

type Question struct {
	ID         string   `json:"id,omitempty"`
	Text       string   `json:"text,omitempty"`
	Type       string   `json:"type,omitempty"`
	Score      float64  `json:"score,omitempty"`
	MessageIds []string `json:"messageIds,omitempty"`
	From       From     `json:"from,omitempty"`
}

type FollowUp struct {
	ID         string          `json:"id,omitempty"`
	Text       string          `json:"text,omitempty"`
	Type       string          `json:"type,omitempty"`
	Score      int             `json:"score,omitempty"`
	MessageIds []string        `json:"messageIds,omitempty"`
	Entities   []EntityInsight `json:"entities,omitempty"`
	Phrases    []string        `json:"phrases,omitempty"` // TODO: I believe this is []string. Need to validate.
	From       From            `json:"from,omitempty"`
	Definitive bool            `json:"definitive,omitempty"`
	Assignee   struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"assignee,omitempty"`
}

type ActionItem struct {
	ID         string          `json:"id,omitempty"`
	Text       string          `json:"text,omitempty"`
	Type       string          `json:"type,omitempty"`
	Score      float64         `json:"score,omitempty"`
	MessageIds []string        `json:"messageIds,omitempty"`
	Entities   []EntityInsight `json:"entities,omitempty"`
	Phrases    []string        `json:"phrases,omitempty"` // TODO: I believe this is []string. Need to validate.
	From       From            `json:"from,omitempty"`
	Definitive bool            `json:"definitive,omitempty"`
	Assignee   struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"assignee,omitempty"`
	DueBy string `json:"dueBy,omitempty"`
}

type Message struct {
	ID             string    `json:"id,omitempty"`
	Text           string    `json:"text,omitempty"`
	From           From      `json:"from,omitempty"`
	StartTime      string    `json:"startTime,omitempty"`
	EndTime        string    `json:"endTime,omitempty"`
	TimeOffset     float64   `json:"timeOffset,omitempty"`
	Duration       float64   `json:"duration,omitempty"`
	ConversationID string    `json:"conversationId,omitempty"`
	Phrases        []string  `json:"phrases,omitempty"` // TODO: I believe this is []string. Need to validate.
	Sentiment      Sentiment `json:"sentiment,omitempty"`
	Words          []struct {
		Word       string  `json:"word,omitempty"`
		StartTime  string  `json:"startTime,omitempty"`
		EndTime    string  `json:"endTime,omitempty"`
		SpeakerTag int     `json:"speakerTag,omitempty"`
		Score      float64 `json:"score,omitempty"`
		TimeOffset float64 `json:"timeOffset,omitempty"`
		Duration   float64 `json:"duration,omitempty"`
	} `json:"words,omitempty"`
}

type Summary struct {
	ID                string       `json:"id,omitempty"`
	Text              string       `json:"text,omitempty"`
	MessageRefs       []MessageRef `json:"messageRefs,omitempty"`
	StartTime         string       `json:"startTime,omitempty"`
	EndTime           string       `json:"endTime,omitempty"`
	BookmarkReference struct {
		ID string `json:"id"`
	} `json:"bookmarkReference"`
}

type Member struct {
	ID    string `json:"id,omitempty" validate:"required"`
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required"`
	Pace  struct {
		Wpm int `json:"wpm,omitempty"`
	} `json:"pace,omitempty"`
	TalkTime struct {
		Percentage float64 `json:"percentage,omitempty"`
		Seconds    float64 `json:"seconds,omitempty"`
	} `json:"talkTime,omitempty"`
	ListenTime struct {
		Percentage float64 `json:"percentage,omitempty"`
		Seconds    float64 `json:"seconds,omitempty"`
	} `json:"listenTime,omitempty"`
	Overlap struct {
	} `json:"overlap,omitempty"` // TODO: need to revisit this
}

type Conversation struct {
	ID        string   `json:"id,omitempty"`
	Type      string   `json:"type,omitempty"`
	Name      string   `json:"name,omitempty"`
	StartTime string   `json:"startTime,omitempty"`
	EndTime   string   `json:"endTime,omitempty"`
	Members   []Member `json:"members,omitempty"`
	Metadata  struct {
	} `json:"metadata,omitempty"` // TODO: need to revisit this
}

type SpeakerEvent struct {
	Type   string `json:"type,omitempty"`
	User   Member `json:"user,omitempty"`
	Offset Offset `json:"offset,omitempty"`
}

type Offset struct {
	Seconds int `json:"seconds,omitempty"`
	Nanos   int `json:"nanos,omitempty"`
}

type BookmarksSummary struct {
	BookmarkID string    `json:"bookmarkId,omitempty"`
	Summaries  []Summary `json:"summary,omitempty"`
}

type Payload struct {
	Content string `json:"content,omitempty" validate:"required"`
}

type TextMessage struct {
	Payload  Payload   `json:"payload,omitempty"`
	From     *From     `json:"from,omitempty"`
	Duration *Duration `json:"duration,omitempty"`
}

/*
	Input parameters for Async API calls
*/
// AsyncURLFileRequest for PostURL PostFile
type AsyncURLFileRequest struct {
	CustomVocabulary                    []string          `json:"customVocabulary,omitempty"`
	ChannelMetadata                     []ChannelMetadata `json:"channelMetadata,omitempty"`
	URL                                 string            `json:"url,omitempty"`
	Name                                string            `json:"name,omitempty"`
	ConfidenceThreshold                 float64           `json:"confidenceThreshold,omitempty"`
	DetectPhrases                       bool              `json:"detectPhrases,omitempty"`
	WebhookURL                          string            `json:"webhookUrl,omitempty"`
	DetectEntities                      bool              `json:"detectEntities,omitempty"`
	LanguageCode                        string            `json:"languageCode,omitempty"`
	Mode                                string            `json:"mode,omitempty"`
	EnableSeparateRecognitionPerChannel bool              `json:"enableSeparateRecognitionPerChannel,omitempty"`
	EnableSpeakerDiarization            bool              `json:"enableSpeakerDiarization,omitempty"`
	DiarizationSpeakerCount             int               `json:"diarizationSpeakerCount,omitempty"`
	ParentRefs                          bool              `json:"parentRefs,omitempty"`
	Sentiment                           bool              `json:"sentiment,omitempty"`
}

type ChannelMetadata struct {
	Speaker Speaker `json:"speaker,omitempty"`
	Channel int     `json:"channel,omitempty"`
}

type Speaker struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type AsyncTextRequest struct {
	Messages            []TextMessage `json:"messages,omitempty" validate:"required"`
	Name                string        `json:"name,omitempty"`
	ConfidenceThreshold float64       `json:"confidenceThreshold,omitempty"`
	DetectPhrases       bool          `json:"detectPhrases,omitempty"`
	WebhookURL          string        `json:"webhookUrl,omitempty"`
	DetectEntities      bool          `json:"detectEntities,omitempty"`
	EnableSummary       bool          `json:"enableSummary,omitempty"`
}

// WaitForJobStatusOpts parameter needed for Wait call
type WaitForJobStatusOpts struct {
	JobId         string `validate:"required"`
	WaitInSeconds int64
}

// MessageRefRequest for BookmarkRequest
type MessageRefRequest struct {
	ID string `json:"id,omitempty"`
}

// BookmarkRequest for creating bookmarks
type BookmarkRequest struct {
	Label           string              `json:"label,omitempty" validate:"required"`
	Description     string              `json:"description,omitempty" validate:"required"`
	User            User                `json:"user,omitempty" validate:"required"`
	BeginTimeOffset int                 `json:"beginTimeOffset,omitempty"`
	Duration        int                 `json:"duration,omitempty"`
	MessageRefs     []MessageRefRequest `json:"messageRefs,omitempty"`
}

type TextSummaryRequest struct {
	Name string `json:"name"`
}

type AudioSummaryRequest struct {
	Name     string `json:"name"`
	AudioURL string `json:"audioUrl"`
}

type VideoSummaryRequest struct {
	Name     string `json:"name"`
	VideoURL string `json:"videoUrl"`
}

type UpdateSpeakerRequest struct {
	SpeakerEvents []SpeakerEvent `json:"speakerEvents"`
}

/*
	Output parameters for Async API calls
*/
type TopicResult struct {
	Topics []Topic `json:"topics"`
}

type QuestionResult struct {
	Questions []Question `json:"questions"`
}

type FollowUpResult struct {
	FollowUps []FollowUp `json:"followUps"`
}

type EntityResult struct {
	Entities []Entity `json:"entities"`
}

type ActionItemResult struct {
	ActionItems []ActionItem `json:"actionItems"`
}

type MessageResult struct {
	Messages []Message `json:"messages"`
}

type SummaryResult struct {
	Summaries []Summary `json:"summary"`
}

type AnalyticsResult struct {
	Metrics []Metric `json:"metrics"`
	Members []Member `json:"members"`
}

type TrackerResult struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Matches []TrackerMatch `json:"matches"`
}

type BookmarksResult struct {
	Bookmarks []Bookmark `json:"bookmarks"`
}

type SummaryUIResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ConversationsResult struct {
	Conversations []Conversation `json:"conversations"`
}

type MembersResult struct {
	Members []Member `json:"members"`
}

type BookmarkSummaryResult struct {
	Summaries []Summary `json:"summary"`
}

type BookmarksSummaryResult struct {
	BookmarksSummary []BookmarksSummary `json:"bookmarksSummary"`
}
