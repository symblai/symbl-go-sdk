// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package interfaces

// Tracker captures a dynamic tracker definition
type Tracker struct {
	Name       string   `json:"name,omitempty"`
	Vocabulary []string `json:"vocabulary,omitempty"`
}

// Trackers provides options for the Tracker functionality
type Trackers struct {
	EnableAllTrackers bool `json:"enableAllTrackers,omitempty"`
	InterimResults    bool `json:"interimResults,omitempty"`
}

// SpeechRecognition provides stream configuration options for real-time conversations
type SpeechRecognition struct {
	Encoding        string `json:"encoding,omitempty"`
	SampleRateHertz int    `json:"sampleRateHertz,omitempty"`
}

// Config captures the general options available for a given conversation
type Config struct {
	ConfidenceThreshold float64           `json:"confidenceThreshold,omitempty"`
	DetectEntities      bool              `json:"detectEntities,omitempty"`
	LanguageCode        string            `json:"languageCode,omitempty"`
	MeetingTitle        string            `json:"meetingTitle,omitempty"`
	Sentiment           bool              `json:"sentiment,omitempty"`
	SpeechRecognition   SpeechRecognition `json:"speechRecognition,omitempty"`
	Trackers            Trackers          `json:"trackers,omitempty"`
}

// Speaker identifies a participant in a conversation
type Speaker struct {
	UserID string `json:"userId,omitempty"`
	Name   string `json:"name,omitempty"`
}

// StreamingConfig captures the options for a real-time conversation
type StreamingConfig struct {
	Type                           string    `json:"type,omitempty"`
	Config                         Config    `json:"config,omitempty"`
	CustomVocabulary               []string  `json:"customVocabulary,omitempty"`
	DisconnectOnStopRequest        bool      `json:"disconnectOnStopRequest,omitempty"`
	DisconnectOnStopRequestTimeout int       `json:"disconnectOnStopRequestTimeout,omitempty"`
	EnableAllTrackers              bool      `json:"enableAllTrackers,omitempty"`
	InsightTypes                   []string  `json:"insightTypes,omitempty"`
	NoConnectionTimeout            bool      `json:"noConnectionTimeout,omitempty"`
	Speaker                        Speaker   `json:"speaker,omitempty"`
	Trackers                       []Tracker `json:"trackers,omitempty"`
}
