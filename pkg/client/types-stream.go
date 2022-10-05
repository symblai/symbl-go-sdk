// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

type StreamingConfig struct {
	Type         string   `json:"type"`
	InsightTypes []string `json:"insightTypes"`
	Config       struct {
		MeetingTitle        string  `json:"meetingTitle"`
		ConfidenceThreshold float64 `json:"confidenceThreshold"`
		TimezoneOffset      int     `json:"timezoneOffset"`
		SpeechRecognition   struct {
			Encoding        string `json:"encoding"`
			SampleRateHertz int    `json:"sampleRateHertz"`
		} `json:"speechRecognition"`
	} `json:"config"`
	Speaker struct {
		UserID string `json:"userId"`
		Name   string `json:"name"`
	} `json:"speaker"`
}
