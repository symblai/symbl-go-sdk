// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import "time"

/*
	Shared definitions
*/

type MessageRecognition struct {
	Type    string `json:"type"`
	IsFinal bool   `json:"isFinal"`
	Payload struct {
		Raw struct {
			Alternatives []struct {
				Words []struct {
					Word      string `json:"word"`
					StartTime struct {
						Seconds string `json:"seconds"`
						Nanos   string `json:"nanos"`
					} `json:"startTime"`
					EndTime struct {
						Seconds string `json:"seconds"`
						Nanos   string `json:"nanos"`
					} `json:"endTime"`
				} `json:"words"`
				Transcript string  `json:"transcript"`
				Confidence float64 `json:"confidence"`
			} `json:"alternatives"`
		} `json:"raw"`
	} `json:"payload"`
	Punctuated struct {
		Transcript string `json:"transcript"`
	} `json:"punctuated"`
	User struct {
		UserID string `json:"userId"`
		Name   string `json:"name"`
		ID     string `json:"id"`
	} `json:"user"`
}

type Message struct {
	From struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		UserID string `json:"userId"`
	} `json:"from"`
	Payload struct {
		Content     string `json:"content"`
		ContentType string `json:"contentType"`
	} `json:"payload"`
	ID      string `json:"id"`
	Channel struct {
		ID string `json:"id"`
	} `json:"channel"`
	Metadata struct {
		DisablePunctuation bool      `json:"disablePunctuation"`
		TimezoneOffset     int       `json:"timezoneOffset"`
		OriginalContent    string    `json:"originalContent"`
		Words              time.Time `json:"words"`
		OriginalMessageID  string    `json:"originalMessageId"`
	} `json:"metadata"`
	Dismissed bool `json:"dismissed"`
	Duration  struct {
		StartTime  time.Time `json:"startTime"`
		EndTime    time.Time `json:"endTime"`
		TimeOffset float64   `json:"timeOffset"`
		Duration   float64   `json:"duration"`
	} `json:"duration"`
}

type Insight struct {
	ID         string  `json:"id"`
	Confidence float64 `json:"confidence"`
	Hints      []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"hints"`
	Type     string `json:"type"`
	Assignee struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		UserID string `json:"userId"`
	} `json:"assignee"`
	Tags []struct {
		Type        string `json:"type"`
		Text        string `json:"text"`
		BeginOffset int    `json:"beginOffset"`
		Value       struct {
			Value struct {
				Name   string `json:"name"`
				Alias  string `json:"alias"`
				UserID string `json:"userId"`
			} `json:"value"`
		} `json:"value"`
	} `json:"tags"`
	Dismissed bool `json:"dismissed"`
	Payload   struct {
		Content     string `json:"content"`
		ContentType string `json:"contentType"`
	} `json:"payload"`
	From struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		UserID string `json:"userId"`
	} `json:"from"`
	Entities         interface{} `json:"entities"` // TODO needs to be defined. Need an example.
	MessageReference struct {
		ID string `json:"id"`
	} `json:"messageReference"`
}

type Topic struct {
	ID                string `json:"id"`
	MessageReferences []struct {
		ID string `json:"id"`
	} `json:"messageReferences"`
	Phrases   string `json:"phrases"`
	RootWords []struct {
		Text string `json:"text"`
	} `json:"rootWords"`
	Score        float64 `json:"score"`
	Type         string  `json:"type"`
	MessageIndex int     `json:"messageIndex"`
}

type Tracker struct {
	ID                string `json:"id"`
	MessageReferences []struct {
		ID       string `json:"id"`
		Relation string `json:"relation"`
	} `json:"messageReferences"`
	Phrases   string `json:"phrases"`
	RootWords []struct {
		Text string `json:"text"`
	} `json:"rootWords"`
	Score float64 `json:"score"`
	Type  string  `json:"type"`
}

/*
	Conversation Insights
*/
type RecognitionResult struct {
	Type       string             `json:"type"`
	Message    MessageRecognition `json:"message"`
	TimeOffset int                `json:"timeOffset"`
}

type MessageResponse struct {
	Type           string    `json:"type"`
	Messages       []Message `json:"messages"`
	SequenceNumber int       `json:"sequenceNumber"`
}

type InsightResponse struct {
	Type           string    `json:"type"`
	Insights       []Insight `json:"insights"`
	SequenceNumber int       `json:"sequenceNumber"`
}

type TopicResponse struct {
	Type   string  `json:"type"`
	Topics []Topic `json:"topics"`
}

type TrackerResponse struct {
	Type     string    `json:"type"`
	Trackers []Tracker `json:"trackers"`
}
