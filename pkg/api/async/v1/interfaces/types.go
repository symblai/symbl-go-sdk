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

type MessageRef struct {
	ID string `json:"id" validate:"required"`
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

/*
	Input parameters for Async API calls
*/
// WaitForJobStatusOpts parameter needed for Wait call
type WaitForJobStatusOpts struct {
	JobId         string `validate:"required"`
	WaitInSeconds int
}

// BookmarkByMessageRefsRequest for creating bookmarks
type BookmarkByMessageRefsRequest struct {
	Label       string `json:"label" validate:"required"`
	Description string `json:"description" validate:"required"`
	User        User   `json:"user" validate:"required"`
	// BeginTimeOffset int          `json:"beginTimeOffset"`
	// Duration        int          `json:"duration"`
	MessageRefs []MessageRef `json:"messageRefs" validate:"required"`
}

// BookmarkByMessageRefsRequest for creating bookmarks
type BookmarkBtTimeDurationsRequest struct {
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
	Topics []struct {
		Text       string   `json:"text"`
		Type       string   `json:"type"`
		Score      float64  `json:"score"`
		MessageIds []string `json:"messageIds"`
		Sentiment  struct {
			Polarity struct {
				Score float64 `json:"score"`
			} `json:"polarity"`
			Suggested string `json:"suggested"`
		} `json:"sentiment"`
		ParentRefs []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"parentRefs"`
	} `json:"topics"`
}

type QuestionResult struct {
	Questions []struct {
		ID         string   `json:"id"`
		Text       string   `json:"text"`
		Type       string   `json:"type"`
		Score      float64  `json:"score"`
		MessageIds []string `json:"messageIds"`
		From       struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
	} `json:"questions"`
}

type FollowUpResult struct {
	FollowUps []struct {
		ID         string        `json:"id"`
		Text       string        `json:"text"`
		Type       string        `json:"type"`
		Score      int           `json:"score"`
		MessageIds []string      `json:"messageIds"`
		Entities   []interface{} `json:"entities"`
		Phrases    []interface{} `json:"phrases"`
		From       struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		Definitive bool `json:"definitive"`
		Assignee   struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"assignee"`
	} `json:"followUps"`
}

type EntityResult struct {
	Entities []struct {
		Type     string `json:"type"`
		SubType  string `json:"subType"`
		Category string `json:"category"`
		Matches  []struct {
			DetectedValue string `json:"detectedValue"`
			MessageRefs   []struct {
				ID        string    `json:"id"`
				StartTime time.Time `json:"startTime"`
				EndTime   time.Time `json:"endTime"`
				Text      string    `json:"text"`
				Offset    int       `json:"offset"`
			} `json:"messageRefs"`
		} `json:"matches"`
	} `json:"entities"`
}

type ActionItemResult struct {
	ActionItems []struct {
		ID         string   `json:"id"`
		Text       string   `json:"text"`
		Type       string   `json:"type"`
		Score      float64  `json:"score"`
		MessageIds []string `json:"messageIds"`
		Entities   []struct {
			Type   string `json:"type"`
			Text   string `json:"text"`
			Offset int    `json:"offset"`
			End    string `json:"end"`
		} `json:"entities"`
		Phrases []interface{} `json:"phrases"`
		From    struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		Definitive bool `json:"definitive"`
		Assignee   struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"assignee"`
		DueBy time.Time `json:"dueBy,omitempty"`
	} `json:"actionItems"`
}

type MessageResult struct {
	Messages []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
		From struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		StartTime      time.Time     `json:"startTime"`
		EndTime        time.Time     `json:"endTime"`
		TimeOffset     float64       `json:"timeOffset"`
		Duration       float64       `json:"duration"`
		ConversationID string        `json:"conversationId"`
		Phrases        []interface{} `json:"phrases"`
		Sentiment      struct {
			Polarity struct {
				Score float64 `json:"score"`
			} `json:"polarity"`
			Suggested string `json:"suggested"`
		} `json:"sentiment"`
		Words []struct {
			Word       string    `json:"word"`
			StartTime  time.Time `json:"startTime"`
			EndTime    time.Time `json:"endTime"`
			SpeakerTag int       `json:"speakerTag"`
			Score      float64   `json:"score"`
			TimeOffset float64   `json:"timeOffset"`
			Duration   float64   `json:"duration"`
		} `json:"words"`
	} `json:"messages"`
}

type SummaryResult struct {
	Summary []struct {
		ID          string `json:"id"`
		Text        string `json:"text"`
		MessageRefs []struct {
			ID string `json:"id"`
		} `json:"messageRefs"`
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
	} `json:"summary"`
}

type AnalyticsResult struct {
	Metrics []struct {
		Type    string  `json:"type"`
		Percent float64 `json:"percent"`
		Seconds float64 `json:"seconds"`
	} `json:"metrics"`
	Members []struct {
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
	} `json:"members"`
}

type TrackerResult []struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Matches []struct {
		Type        string `json:"type"`
		Value       string `json:"value"`
		MessageRefs []struct {
			ID     string `json:"id"`
			Text   string `json:"text"`
			Offset int    `json:"offset"`
		} `json:"messageRefs"`
		InsightRefs []interface{} `json:"insightRefs"`
	} `json:"matches"`
}

/*
	Output for Bookmark APIs
*/
type BookmarksResult struct {
	Bookmarks []Bookmark `json:"bookmarks"`
}
