// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

type Tracker struct {
	Name       string   `json:"name,omitempty"`
	Vocabulary []string `json:"vocabulary,omitempty"`
}

type SpeechRecognition struct {
	Encoding        string `json:"encoding,omitempty"`
	SampleRateHertz int    `json:"sampleRateHertz,omitempty"`
}

type Config struct {
	MeetingTitle        string            `json:"meetingTitle,omitempty"`
	ConfidenceThreshold float64           `json:"confidenceThreshold,omitempty"`
	TimezoneOffset      int               `json:"timezoneOffset,omitempty"`
	SpeechRecognition   SpeechRecognition `json:"speechRecognition,omitempty"`
}

type Speaker struct {
	UserID string `json:"userId,omitempty"`
	Name   string `json:"name,omitempty"`
}

type StreamingConfig struct {
	Type             string    `json:"type,omitempty"`
	InsightTypes     []string  `json:"insightTypes,omitempty"`
	CustomVocabulary []string  `json:"customVocabulary,omitempty"`
	Trackers         []Tracker `json:"trackers,omitempty"`
	Config           Config    `json:"config,omitempty"`
	Speaker          Speaker   `json:"speaker,omitempty"`
}
