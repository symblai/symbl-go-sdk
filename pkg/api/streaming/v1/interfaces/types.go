// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

/*
	Shared definitions
*/
type MessageReference struct {
	ID string `json:"id,omitempty"`
}

type From struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID string `json:"userId,omitempty"`
}

type Assignee From

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

type MessageRecognition struct {
	Type    string `json:"type,omitempty"`
	IsFinal bool   `json:"isFinal,omitempty"`
	Payload struct {
		Raw struct {
			Alternatives []struct {
				Words []struct {
					Word      string `json:"word,omitempty"`
					StartTime struct {
						Seconds string `json:"seconds,omitempty"`
						Nanos   string `json:"nanos,omitempty"`
					} `json:"startTime,omitempty"`
					EndTime struct {
						Seconds string `json:"seconds,omitempty"`
						Nanos   string `json:"nanos,omitempty"`
					} `json:"endTime,omitempty"`
				} `json:"words,omitempty"`
				Transcript string  `json:"transcript,omitempty"`
				Confidence float64 `json:"confidence,omitempty"`
			} `json:"alternatives,omitempty"`
		} `json:"raw,omitempty"`
	} `json:"payload,omitempty"`
	Punctuated struct {
		Transcript string `json:"transcript,omitempty"`
	} `json:"punctuated,omitempty"`
	User struct {
		UserID string `json:"userId,omitempty"`
		Name   string `json:"name,omitempty"`
		ID     string `json:"id,omitempty"`
	} `json:"user,omitempty"`
}

type Message struct {
	From    From `json:"from,omitempty"`
	Payload struct {
		Content     string `json:"content,omitempty"`
		ContentType string `json:"contentType,omitempty"`
	} `json:"payload,omitempty"`
	ID      string `json:"id,omitempty"`
	Channel struct {
		ID string `json:"id,omitempty"`
	} `json:"channel,omitempty"`
	Metadata struct {
		DisablePunctuation bool   `json:"disablePunctuation,omitempty"`
		TimezoneOffset     int    `json:"timezoneOffset,omitempty"`
		OriginalContent    string `json:"originalContent,omitempty"`
		Words              string `json:"words,omitempty"`
		OriginalMessageID  string `json:"originalMessageId,omitempty"`
	} `json:"metadata,omitempty"`
	Dismissed bool `json:"dismissed,omitempty"`
	Duration  struct {
		StartTime  string  `json:"startTime,omitempty"`
		EndTime    string  `json:"endTime,omitempty"`
		TimeOffset float64 `json:"timeOffset,omitempty"`
		Duration   float64 `json:"duration,omitempty"`
	} `json:"duration,omitempty"`
	Entities []Entity `json:"entities,omitempty"`
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

type Insight struct {
	ID         string  `json:"id,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
	Hints      []struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"hints,omitempty"`
	Type      string   `json:"type,omitempty"`
	Assignee  Assignee `json:"assignee,omitempty"`
	Tags      []Tag    `json:"tags,omitempty"`
	Dismissed bool     `json:"dismissed,omitempty"`
	Payload   struct {
		Content     string `json:"content,omitempty"`
		ContentType string `json:"contentType,omitempty"`
	} `json:"payload,omitempty"`
	From             From             `json:"from,omitempty"`
	Entities         []Entity         `json:"entities,omitempty"`
	MessageReference MessageReference `json:"messageReference,omitempty"`
}

type RootWord struct {
	Text string `json:"text,omitempty"`
}

type Topic struct {
	ID                string             `json:"id,omitempty"`
	MessageReferences []MessageReference `json:"messageReferences,omitempty"`
	Phrases           string             `json:"phrases,omitempty"`
	RootWords         []RootWord         `json:"rootWords,omitempty"`
	Score             float64            `json:"score,omitempty"`
	Type              string             `json:"type,omitempty"`
	MessageIndex      int                `json:"messageIndex,omitempty"`
}

type TrackerMatch struct {
	Value       string       `json:"value,omitempty"`
	MessageRefs []MessageRef `json:"messageRefs,omitempty"`
	InsightRefs []InsightRef `json:"insightRefs,omitempty"`
}

type Tracker struct {
	ID      string         `json:"id,omitempty"`
	Name    string         `json:"name,omitempty"`
	Matches []TrackerMatch `json:"matches,omitempty"`
}

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
	Type       string             `json:"type,omitempty"`
	Message    MessageRecognition `json:"message,omitempty"`
	TimeOffset int                `json:"timeOffset,omitempty"`
}

type MessageResponse struct {
	Type           string    `json:"type,omitempty"`
	Messages       []Message `json:"messages,omitempty"`
	SequenceNumber int       `json:"sequenceNumber,omitempty"`
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
	Type     string    `json:"type,omitempty"`
	Trackers []Tracker `json:"trackers,omitempty"`
}

type EntityResponse struct {
	Type           string   `json:"type,omitempty"`
	Entities       []Entity `json:"entities,omitempty"`
	SequenceNumber int      `json:"sequenceNumber,omitempty"`
}
