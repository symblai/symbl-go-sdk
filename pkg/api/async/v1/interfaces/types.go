// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import "time"

/*
	Shared definitions
*/
type User struct {
	Name   string `json:"name" validate:"required"`
	UserID string `json:"userId" validate:"required"`
	Email  string `json:"email" validate:"required"`
}

type Metric struct {
	Type    string  `json:"type"`
	Percent float64 `json:"percent"`
	Seconds float64 `json:"seconds"`
}

type Sentiment struct {
	Polarity struct {
		Score float64 `json:"score"`
	} `json:"polarity"`
	Suggested string `json:"suggested"`
}

type MessageRef struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Text      string    `json:"text"`
	Offset    int       `json:"offset"`
}

type ParentRef []struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type InsightRef []struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type TrackerMatch struct {
	Type        string       `json:"type"`
	Value       string       `json:"value"`
	MessageRefs []MessageRef `json:"messageRefs"`
	InsightRefs []InsightRef `json:"insightRefs"`
}

type EntityMatch struct {
	DetectedValue string       `json:"detectedValue"`
	MessageRefs   []MessageRef `json:"messageRefs"`
}

type EntityInsight struct {
	Type   string `json:"type"`
	Text   string `json:"text"`
	Offset int    `json:"offset"`
	End    string `json:"end"`
}

type Entity struct {
	Type     string        `json:"type"`
	SubType  string        `json:"subType"`
	Category string        `json:"category"`
	Matches  []EntityMatch `json:"matches"`
}

type Bookmark struct {
	ID              string       `json:"id"`
	Label           string       `json:"label"  validate:"required"`
	Description     string       `json:"description"`
	User            User         `json:"user" validate:"required"`
	BeginTimeOffset int          `json:"beginTimeOffset" validate:"required"`
	Duration        int          `json:"duration" validate:"required"`
	MessageRefs     []MessageRef `json:"messageRefs" validate:"required"`
}

type Topic struct {
	Text       string      `json:"text"`
	Type       string      `json:"type"`
	Score      float64     `json:"score"`
	MessageIds []string    `json:"messageIds"`
	Sentiment  Sentiment   `json:"sentiment"`
	ParentRefs []ParentRef `json:"parentRefs"`
}

type Question struct {
	ID         string   `json:"id"`
	Text       string   `json:"text"`
	Type       string   `json:"type"`
	Score      float64  `json:"score"`
	MessageIds []string `json:"messageIds"`
	From       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"from"`
}

type FollowUp struct {
	ID         string          `json:"id"`
	Text       string          `json:"text"`
	Type       string          `json:"type"`
	Score      int             `json:"score"`
	MessageIds []string        `json:"messageIds"`
	Entities   []EntityInsight `json:"entities"`
	Phrases    []string        `json:"phrases"` // TODO: I believe this is []string. Need to validate.
	From       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"from"`
	Definitive bool `json:"definitive"`
	Assignee   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"assignee"`
}

type ActionItem struct {
	ID         string          `json:"id"`
	Text       string          `json:"text"`
	Type       string          `json:"type"`
	Score      float64         `json:"score"`
	MessageIds []string        `json:"messageIds"`
	Entities   []EntityInsight `json:"entities"`
	Phrases    []string        `json:"phrases"` // TODO: I believe this is []string. Need to validate.
	From       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"from"`
	Definitive bool `json:"definitive"`
	Assignee   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"assignee"`
	DueBy time.Time `json:"dueBy,omitempty"`
}

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	From struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"from"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	TimeOffset     float64   `json:"timeOffset"`
	Duration       float64   `json:"duration"`
	ConversationID string    `json:"conversationId"`
	Phrases        []string  `json:"phrases"` // TODO: I believe this is []string. Need to validate.
	Sentiment      Sentiment `json:"sentiment"`
	Words          []struct {
		Word       string    `json:"word"`
		StartTime  time.Time `json:"startTime"`
		EndTime    time.Time `json:"endTime"`
		SpeakerTag int       `json:"speakerTag"`
		Score      float64   `json:"score"`
		TimeOffset float64   `json:"timeOffset"`
		Duration   float64   `json:"duration"`
	} `json:"words"`
}

type Summary struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	MessageRefs []struct {
		ID string `json:"id"`
	} `json:"messageRefs"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pace struct {
		Wpm int `json:"wpm"`
	} `json:"pace"`
	TalkTime struct {
		Percentage float64 `json:"percentage"`
		Seconds    float64 `json:"seconds"`
	} `json:"talkTime"`
	ListenTime struct {
		Percentage float64 `json:"percentage"`
		Seconds    float64 `json:"seconds"`
	} `json:"listenTime"`
	Overlap struct {
	} `json:"overlap"`
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
