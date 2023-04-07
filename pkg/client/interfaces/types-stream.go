// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package interfaces

type Tracker struct {
	Name       string   `json:"name,omitempty"`
	Vocabulary []string `json:"vocabulary,omitempty"`
}

type Trackers struct {
	EnableAllTrackers bool `json:"enableAllTrackers,omitempty"`
	InterimResults    bool `json:"interimResults,omitempty"`
}

type SpeechRecognition struct {
	Encoding        string `json:"encoding,omitempty"`
	SampleRateHertz int    `json:"sampleRateHertz,omitempty"`
}

type Config struct {
	ConfidenceThreshold float64           `json:"confidenceThreshold,omitempty"`
	DetectEntities      bool              `json:"detectEntities,omitempty"`
	LanguageCode        string            `json:"languageCode,omitempty"`
	MeetingTitle        string            `json:"meetingTitle,omitempty"`
	Sentiment           bool              `json:"sentiment,omitempty"`
	SpeechRecognition   SpeechRecognition `json:"speechRecognition,omitempty"`
	Trackers            Trackers          `json:"trackers,omitempty"`
}

type Speaker struct {
	UserID string `json:"userId,omitempty"`
	Name   string `json:"name,omitempty"`
}

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
