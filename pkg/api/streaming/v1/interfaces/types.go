// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

/*
	Shared definitions
*/
// message recognition
type EndTime struct {
	Nanos   string `json:"nanos,omitempty"`
	Seconds string `json:"seconds,omitempty"`
}
type StartTime struct {
	Nanos   string `json:"nanos,omitempty"`
	Seconds string `json:"seconds,omitempty"`
}
type Words struct {
	EndTime   EndTime   `json:"endTime,omitempty"`
	StartTime StartTime `json:"startTime,omitempty"`
	Word      string    `json:"word,omitempty"`
}
type Alternatives struct {
	Confidence float64 `json:"confidence,omitempty"`
	Transcript string  `json:"transcript,omitempty"`
	Words      []Words `json:"words,omitempty"`
}
type Raw struct {
	Alternatives []Alternatives `json:"alternatives,omitempty"`
}
type RecognitionPayload struct {
	Raw Raw `json:"raw,omitempty"`
}
type Punctuated struct {
	Transcript string `json:"transcript,omitempty"`
}
type User struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID string `json:"userId,omitempty"`
}
type Recognition struct {
	IsFinal    bool               `json:"isFinal,omitempty"`
	Payload    RecognitionPayload `json:"payload,omitempty"`
	Punctuated Punctuated         `json:"punctuated,omitempty"`
	Type       string             `json:"type,omitempty"`
	User       User               `json:"user,omitempty"`
}

// message result
type Channel struct {
	ID string `json:"id,omitempty"`
}
type Duration struct {
	StartTime  string  `json:"startTime,omitempty"`
	EndTime    string  `json:"endTime,omitempty"`
	TimeOffset float64 `json:"timeOffset,omitempty"`
	Duration   float64 `json:"duration,omitempty"`
}
type Entities struct {
	Category      string  `json:"category,omitempty"`
	DetectedValue string  `json:"detectedValue,omitempty"`
	Message       Message `json:"message,omitempty"`
	Offset        int     `json:"offset,omitempty"`
	SubType       string  `json:"subType,omitempty"`
	Type          string  `json:"type,omitempty"`
}
type From User
type Metadata struct {
	DisablePunctuation bool   `json:"disablePunctuation,omitempty"`
	OriginalContent    string `json:"originalContent,omitempty"`
	OriginalMessageID  string `json:"originalMessageId,omitempty"`
	Words              string `json:"words,omitempty"`
}
type Payload struct {
	Content     string `json:"content,omitempty"`
	ContentType string `json:"contentType,omitempty"`
}
type Polarity struct {
	Score float64 `json:"score,omitempty"`
}
type Sentiment struct {
	Polarity  Polarity `json:"polarity,omitempty"`
	Suggested string   `json:"suggested,omitempty"`
}
type Message struct {
	Channel   Channel    `json:"channel,omitempty"`
	Dismissed bool       `json:"dismissed,omitempty"`
	Duration  Duration   `json:"duration,omitempty"`
	Entities  []Entities `json:"entities,omitempty"`
	From      From       `json:"from,omitempty"`
	ID        string     `json:"id,omitempty"`
	Metadata  Metadata   `json:"metadata,omitempty"`
	Payload   Payload    `json:"payload,omitempty"`
	Sentiment Sentiment  `json:"sentiment,omitempty"`
}

// insight
type Assignee User
type MessageReference struct {
	ID string `json:"id,omitempty"`
}
type Hints struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
type Tag struct {
	Type        string `json:"type,omitempty"`
	Text        string `json:"text,omitempty"`
	BeginOffset int    `json:"beginOffset,omitempty"`
	Value       struct {
		Value struct {
			Name   string `json:"name,omitempty"`
			Alias  string `json:"alias,omitempty"`
			UserID string `json:"userId,omitempty"`
		} `json:"value,omitempty"`
	} `json:"value,omitempty"`
}
type MessageReferences struct {
	ID string `json:"id,omitempty"`
}
type Insight struct {
	Assignee         Assignee         `json:"assignee,omitempty"`
	Confidence       float64          `json:"confidence,omitempty"`
	Dismissed        bool             `json:"dismissed,omitempty"`
	Entities         []Entity         `json:"entities,omitempty"`
	From             From             `json:"from,omitempty"`
	Hints            []Hints          `json:"hints,omitempty"`
	ID               string           `json:"id,omitempty"`
	MessageReference MessageReference `json:"messageReference,omitempty"`
	Payload          Payload          `json:"payload,omitempty"`
	Tags             []Tag            `json:"tags,omitempty"`
	Type             string           `json:"type,omitempty"`
}

// topic
type RootWord struct {
	Text string `json:"text,omitempty"`
}
type Topic struct {
	ID                string              `json:"id,omitempty"`
	MessageIndex      int                 `json:"messageIndex,omitempty"`
	MessageReferences []MessageReferences `json:"messageReferences,omitempty"`
	Phrases           string              `json:"phrases,omitempty"`
	RootWords         []RootWord          `json:"rootWords,omitempty"`
	Score             float64             `json:"score,omitempty"`
	Sentiment         Sentiment           `json:"sentiment,omitempty"`
	Type              string              `json:"type,omitempty"`
}

// tracker
type MessageRef struct {
	ID        string `json:"id,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Text      string `json:"text,omitempty"`
	Offset    int    `json:"offset,omitempty"`
}
type InsightRef struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}
type TrackerMatch struct {
	InsightRefs []InsightRef `json:"insightRefs,omitempty"`
	MessageRefs []MessageRef `json:"messageRefs,omitempty"`
	Type        string       `json:"type,omitempty"`
	Value       string       `json:"value,omitempty"`
}
type Tracker struct {
	ID      string         `json:"id,omitempty"`
	Name    string         `json:"name,omitempty"`
	Matches []TrackerMatch `json:"matches,omitempty"`
}

// entity
type EntityMatch struct {
	DetectedValue string       `json:"detectedValue,omitempty"`
	MessageRefs   []MessageRef `json:"messageRefs,omitempty"`
}
type Entity struct {
	Type     string        `json:"type,omitempty"`
	SubType  string        `json:"subType,omitempty"`
	Category string        `json:"category,omitempty"`
	Matches  []EntityMatch `json:"matches,omitempty"`
}

/*
	Conversation Insights
*/
type InitializationMessage struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
		Data struct {
			ConversationID string `json:"conversationId"`
		} `json:"data"`
	} `json:"message"`
}

type RecognitionResult struct {
	Type       string      `json:"type,omitempty"`
	Message    Recognition `json:"message,omitempty"`
	TimeOffset int         `json:"timeOffset,omitempty"`
}

type MessageResponse struct {
	Messages       []Message `json:"messages,omitempty"`
	Sentiment      bool      `json:"sentiment,omitempty"`
	SequenceNumber int       `json:"sequenceNumber,omitempty"`
	Type           string    `json:"type,omitempty"`
}

type InsightResponse struct {
	Type           string    `json:"type,omitempty"`
	Insights       []Insight `json:"insights,omitempty"`
	SequenceNumber int       `json:"sequenceNumber,omitempty"`
}

type TopicResponse struct {
	Type   string  `json:"type,omitempty"`
	Topics []Topic `json:"topics,omitempty"`
}

type TrackerResponse struct {
	IsFinal        bool      `json:"isFinal,omitempty"`
	SequenceNumber int       `json:"sequenceNumber,omitempty"`
	Trackers       []Tracker `json:"trackers,omitempty"`
	Type           string    `json:"type,omitempty"`
}

type EntityResponse struct {
	Type           string   `json:"type,omitempty"`
	Entities       []Entity `json:"entities,omitempty"`
	SequenceNumber int      `json:"sequenceNumber,omitempty"`
}

type TeardownMessage struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
		Data struct {
			ConversationID string `json:"conversationId"`
		} `json:"data"`
	} `json:"message"`
}
