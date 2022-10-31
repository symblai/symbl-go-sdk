// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import "time"

/*
	Shared definitions
*/
type User struct {
	Name   string `json:"name,omitempty" validate:"required"`
	UserID string `json:"userId,omitempty" validate:"required"`
	Email  string `json:"email,omitempty" validate:"required"`
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
	ID        string    `json:"id,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Text      string    `json:"text,omitempty"`
	Offset    int       `json:"offset,omitempty"`
}

type ParentRef []struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type InsightRef []struct {
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

type Bookmark struct {
	ID              string       `json:"id,omitempty"`
	Label           string       `json:"label,omitempty"  validate:"required"`
	Description     string       `json:"description,omitempty"`
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
	From       struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"from,omitempty"`
}

type FollowUp struct {
	ID         string          `json:"id,omitempty"`
	Text       string          `json:"text,omitempty"`
	Type       string          `json:"type,omitempty"`
	Score      int             `json:"score,omitempty"`
	MessageIds []string        `json:"messageIds,omitempty"`
	Entities   []EntityInsight `json:"entities,omitempty"`
	Phrases    []string        `json:"phrases,omitempty"` // TODO: I believe this is []string. Need to validate.
	From       struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"from,omitempty"`
	Definitive bool `json:"definitive,omitempty"`
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
	From       struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"from,omitempty"`
	Definitive bool `json:"definitive,omitempty"`
	Assignee   struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"assignee,omitempty"`
	DueBy time.Time `json:"dueBy,omitempty,omitempty"`
}

type Message struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	From struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"from,omitempty"`
	StartTime      time.Time `json:"startTime,omitempty"`
	EndTime        time.Time `json:"endTime,omitempty"`
	TimeOffset     float64   `json:"timeOffset,omitempty"`
	Duration       float64   `json:"duration,omitempty"`
	ConversationID string    `json:"conversationId,omitempty"`
	Phrases        []string  `json:"phrases,omitempty"` // TODO: I believe this is []string. Need to validate.
	Sentiment      Sentiment `json:"sentiment,omitempty"`
	Words          []struct {
		Word       string    `json:"word,omitempty"`
		StartTime  time.Time `json:"startTime,omitempty"`
		EndTime    time.Time `json:"endTime,omitempty"`
		SpeakerTag int       `json:"speakerTag,omitempty"`
		Score      float64   `json:"score,omitempty"`
		TimeOffset float64   `json:"timeOffset,omitempty"`
		Duration   float64   `json:"duration,omitempty"`
	} `json:"words,omitempty"`
}

type Summary struct {
	ID          string `json:"id,omitempty"`
	Text        string `json:"text,omitempty"`
	MessageRefs []struct {
		ID string `json:"id,omitempty"`
	} `json:"messageRefs,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

type Member struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Pace struct {
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
	} `json:"overlap,omitempty"`
}

/*
	Input parameters for Async API calls
*/
// WaitForJobStatusOpts parameter needed for Wait call
type WaitForJobStatusOpts struct {
	JobId         string `validate:"required"`
	WaitInSeconds int
}

// MessageRefRequest for BookmarkByMessageRefsRequest
type MessageRefRequest struct {
	ID string `json:"id" validate:"required"`
}

// BookmarkByMessageRefsRequest for creating bookmarks
type BookmarkByMessageRefsRequest struct {
	Label       string `json:"label" validate:"required"`
	Description string `json:"description" validate:"required"`
	User        User   `json:"user" validate:"required"`
	// BeginTimeOffset int          `json:"beginTimeOffset"`
	// Duration        int          `json:"duration"`
	MessageRefs []MessageRefRequest `json:"messageRefs" validate:"required"`
}

// BookmarkByMessageRefsRequest for creating bookmarks
type BookmarkByTimeDurationsRequest struct {
	Label           string `json:"label" validate:"required"`
	Description     string `json:"description" validate:"required"`
	User            User   `json:"user" validate:"required"`
	BeginTimeOffset int    `json:"beginTimeOffset" validate:"required"`
	Duration        int    `json:"duration" validate:"required"`
	// MessageRefs []MessageRef `json:"messageRefs"`
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

type TrackerResult []struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Matches []TrackerMatch `json:"matches"`
}

type BookmarksResult struct {
	Bookmarks []Bookmark `json:"bookmarks"`
}
