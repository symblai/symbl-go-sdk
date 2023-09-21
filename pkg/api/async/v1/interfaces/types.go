// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
Defines everything that makes up the Async API interface
*/
package interfaces

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
TODO: When exercising the API and description is blank...

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
// shared structs in request
type Phrases struct {
	HighlightOnlyInsightKeyPhrases bool `json:"highlightOnlyInsightKeyPhrases,omitempty"`
	HighlightAllKeyPhrases         bool `json:"highlightAllKeyPhrases,omitempty"`
}

type Speaker struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type ChannelMetadata struct {
	Speaker Speaker `json:"speaker,omitempty"`
	Channel int     `json:"channel,omitempty"`
}

type Features struct {
	FeatureList []string `json:"featureList"`
}

type Metadata struct {
	SalesStage   string `json:"salesStage"`
	ProspectName string `json:"prospectName"`
}

// AsyncURLFileRequest for PostURL to post a file to the platform
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
	Features                            Features          `json:"features,omitempty"`
	ConversationType                    string            `json:"conversationType,omitempty"`
	Metadata                            Metadata          `json:"metadata,omitempty"`
}

// AsyncTextRequest for PostText to post text to the platform
type AsyncTextRequest struct {
	Messages            []TextMessage `json:"messages,omitempty" validate:"required"`
	Name                string        `json:"name,omitempty"`
	ConfidenceThreshold float64       `json:"confidenceThreshold,omitempty"`
	DetectPhrases       bool          `json:"detectPhrases,omitempty"`
	WebhookURL          string        `json:"webhookUrl,omitempty"`
	DetectEntities      bool          `json:"detectEntities,omitempty"`
	EnableSummary       bool          `json:"enableSummary,omitempty"`
	Features            Features      `json:"features,omitempty"`
	ConversationType    string        `json:"conversationType,omitempty"`
	Metadata            Metadata      `json:"metadata,omitempty"`
}

// WaitForJobStatusOpts parameter needed for Wait call
type WaitForJobStatusOpts struct {
	JobId              string `validate:"required"`
	TotalWaitInSeconds int64
	WaitInSeconds      int64
}

type MessageRefRequest struct {
	ID string `json:"id,omitempty"`
}

type BookmarkRequest struct {
	Label           string              `json:"label,omitempty" validate:"required"`
	Description     string              `json:"description,omitempty" validate:"required"`
	User            User                `json:"user,omitempty" validate:"required"`
	BeginTimeOffset int                 `json:"beginTimeOffset,omitempty"`
	Duration        int                 `json:"duration,omitempty"`
	MessageRefs     []MessageRefRequest `json:"messageRefs,omitempty"`
}

type TextSummaryRequest struct {
	Name string `json:"name,omitempty"`
}

type AudioSummaryRequest struct {
	Name     string `json:"name,omitempty"`
	AudioURL string `json:"audioUrl,omitempty"`
}

type VideoSummaryRequest struct {
	Name     string `json:"name,omitempty"`
	VideoURL string `json:"videoUrl,omitempty"`
}

type UpdateSpeakerRequest struct {
	SpeakerEvents []SpeakerEvent `json:"speakerEvents,omitempty"`
}

type TranscriptRequest struct {
	ContentType           string  `json:"contentType"`
	Phrases               Phrases `json:"phrases,omitempty"`
	CreateParagraphs      bool    `json:"createParagraphs,omitempty"`
	ShowSpeakerSeparation bool    `json:"showSpeakerSeparation,omitempty"`
}

/*
	Output parameters for Async API calls
*/
// shared structs in result
type Transcript struct {
	Payload     string `json:"payload,omitempty"`
	ContentType string `json:"contentType,omitempty"`
}

// TopicResult provides Async API results for topics
type TopicResult struct {
	Topics []Topic `json:"topics,omitempty"`
}

// QuestionResult provides Async API results for questions
type QuestionResult struct {
	Questions []Question `json:"questions,omitempty"`
}

// FollowUpResult provides Async API results for follow ups
type FollowUpResult struct {
	FollowUps []FollowUp `json:"followUps,omitempty"`
}

// EntityResult provides Async API results for entities
type EntityResult struct {
	Entities []Entity `json:"entities,omitempty"`
}

// ActionItemResult provides Async API results for action items
type ActionItemResult struct {
	ActionItems []ActionItem `json:"actionItems,omitempty"`
}

// MessageResult provides Async API results for message results
type MessageResult struct {
	Messages []Message `json:"messages,omitempty"`
}

// SummaryResult provides Async API results for summary results
type SummaryResult struct {
	Summaries []Summary `json:"summary,omitempty"`
}

// AnalyticsResult provides Async API results for analytics results
type AnalyticsResult struct {
	Metrics []Metric `json:"metrics,omitempty"`
	Members []Member `json:"members,omitempty"`
}

// TrackerResult provides Async API results for tracker results
type TrackerResult struct {
	ID      string         `json:"id,omitempty"`
	Name    string         `json:"name,omitempty"`
	Matches []TrackerMatch `json:"matches,omitempty"`
}

// BookmarksResult provides Async API results for bookmarks results
type BookmarksResult struct {
	Bookmarks []Bookmark `json:"bookmarks,omitempty"`
}

// SummaryUIResult provides Async API results for summary ui results
type SummaryUIResult struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// ConversationsResult provides Async API results for conversation results
type ConversationsResult struct {
	Conversations []Conversation `json:"conversations,omitempty"`
}

// MembersResult provides Async API results for member results
type MembersResult struct {
	Members []Member `json:"members,omitempty"`
}

// BookmarkSummaryResult provides Async API results for bookmark summary results
type BookmarkSummaryResult struct {
	Summaries []Summary `json:"summary,omitempty"`
}

// BookmarksSummaryResult provides Async API results for bookmarks summary results
type BookmarksSummaryResult struct {
	BookmarksSummary []BookmarksSummary `json:"bookmarksSummary,omitempty"`
}

type TranscriptResult struct {
	Transcript Transcript `json:"transcript,omitempty"`
}

/*
	Internal package messages
*/
// InitializationMessage is an internal representation for an Async conversation start event
type InitializationMessage struct {
	ConversationID string `json:"conversationId,omitempty"`
}

// TeardownMessage is an internal representation for an Async conversation stop event
type TeardownMessage struct {
	ConversationID string `json:"conversationId,omitempty"`
}

// CallScoreResult
type CallScoreResult struct{}

// CallScoreStatusResult
type CallScoreStatusResult struct{}

// InsightStatusResult
type InsightStatusResult struct{}
