// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

import (
	"context"

	"github.com/google/uuid"
	klog "k8s.io/klog/v2"

	streaming "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	common "github.com/dvonthenen/symbl-go-sdk/pkg/client/common"
	stream "github.com/dvonthenen/symbl-go-sdk/pkg/client/stream"
)

const (
	defaultConfidenceThreshold float64 = 0.7
	defaultSampleRateHertz     int     = 16000
	defaultUserID              string  = "user@email.com"
	defaultUserName            string  = "Jane Doe"
)

// type StreamingConfig struct {
// 	Type         string   `json:"type"`
// 	InsightTypes []string `json:"insightTypes"`
// 	Config       struct {
// 		MeetingTitle        string  `json:"meetingTitle"`
// 		ConfidenceThreshold float64 `json:"confidenceThreshold"`
// 		TimezoneOffset      int     `json:"timezoneOffset"`
// 		LanguageCode        string  `json:"languageCode"`
// 		SampleRateHertz     int     `json:"sampleRateHertz"`
// 	} `json:"config"`
// 	Speaker struct {
// 		UserID string `json:"userId"`
// 		Name   string `json:"name"`
// 	} `json:"speaker"`
// }

/*
	Example Config:
	{
	"type": "start_request",
	"insightTypes": ["question", "action_item"],
	"config": {
		"confidenceThreshold": 0.9,
		"timezoneOffset": 480,
		"speechRecognition": {
		"encoding": "LINEAR16",
		"sampleRateHertz": 44100
		},
		"meetingTitle": "Client Meeting"
	},
	"speaker": {
		"userId": "jane.doe@example.com",
		"name": "Jane"
	}
	}
*/

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

func getDefaultConfig() *StreamingConfig {
	config := &StreamingConfig{}

	config.Type = streaming.TypeRequestStart
	config.InsightTypes = []string{"topic", "question", "action_item", "follow_up"}
	config.Config.MeetingTitle = "my-meeting"
	config.Config.ConfidenceThreshold = defaultConfidenceThreshold
	// config.Config.TimezoneOffset = 480
	// config.Config.LanguageCode = "en-US"
	config.Config.SpeechRecognition.Encoding = "LINEAR16"
	config.Config.SpeechRecognition.SampleRateHertz = defaultSampleRateHertz
	config.Speaker.Name = defaultUserName
	config.Speaker.UserID = defaultUserID

	// TODO remove... the above seems to work
	// config.Type = streaming.TypeRequestStart
	// config.InsightTypes = []string{"topic", "question", "action_item"}
	// config.Config.ConfidenceThreshold = defaultConfidenceThreshold
	// config.Config.SpeechRecognition.SampleRateHertz = defaultSampleRateHertz
	// config.Speaker.Name = defaultUserName
	// config.Speaker.UserID = defaultUserID

	return config
}

// TODO remove... the above seems to work
// type StreamingConfig struct {
// 	Type         string   `json:"type"`
// 	InsightTypes []string `json:"insightTypes"`
// 	Config       struct {
// 		ConfidenceThreshold float64 `json:"confidenceThreshold"`
// 		SpeechRecognition   struct {
// 			Encoding        string `json:"encoding"`
// 			SampleRateHertz int    `json:"sampleRateHertz"`
// 		} `json:"speechRecognition"`
// 	} `json:"config"`
// 	Speaker struct {
// 		UserID string `json:"userId"`
// 		Name   string `json:"name"`
// 	} `json:"speaker"`
// }

type StreamClient struct {
	*stream.WebSocketClient

	restClient     *RestClient
	symblStreaming stream.WebSocketMessageCallback
}

// NewClient creates a new client on the Symbl.ai platform. The client authenticates with the
// server with APP_ID/APP_SECRET.
func NewStreamClient(ctx context.Context) (*StreamClient, error) {
	klog.V(6).Infof("NewStreamClient ENTER\n")

	config := getDefaultConfig()

	// create rest client
	restClient, err := NewRestClient(ctx)
	if err != nil {
		klog.V(2).Infof("NewRestClient failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, err
	}
	klog.V(6).Infof("IMPORTANT: Never print in production\n")
	klog.V(6).Infof("AppId: %s\n", restClient.creds.AppId)
	klog.V(6).Infof("AppSecret: %s\n", restClient.creds.AppSecret)

	// generate unique id... not even sure why this is needed, but hey
	id := uuid.New()
	klog.V(2).Infof("UUID: %s\n", id.String())

	streamPath := version.GetStreamingAPI(version.StreamPath, id)
	klog.V(6).Infof("IMPORTANT: Never print in production\n")
	klog.V(6).Infof("streamPath: %s\n", streamPath)

	// init symbl websocket message router
	symblStreaming := streaming.New()

	// create client
	creds := stream.Credentials{
		Host:      streaming.SymblPlatformHost,
		Channel:   streamPath,
		AccessKey: restClient.auth.AccessToken,
	}
	wsClient, err := stream.NewWebSocketClient(creds, symblStreaming)
	if err != nil {
		klog.V(2).Infof("stream.NewWebSocketClient failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, err
	}

	// establish connection
	wsConnection := wsClient.Connect()
	if wsConnection == nil {
		klog.V(2).Infof("stream.NewWebSocketClient failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, common.ErrWebSocketInitializationFailed
	}

	// write Symbl config to Platform
	err = wsClient.WriteJSON(config)
	if err != nil {
		klog.V(2).Infof("wsClient.WriteJSON failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, err
	}

	// save client for return
	streamClient := &StreamClient{
		wsClient,
		restClient,
		symblStreaming,
	}

	klog.V(2).Infof("NewStreamClientWithCreds Succeeded\n")
	klog.V(6).Infof("NewStreamClient LEAVE\n")
	return streamClient, nil
}

// TODO: Delete this... redundant
// // NewStreamClientWithCreds creates a new client on the Symbl.ai platform. The client authenticates with the
// // server with APP_ID/APP_SECRET.
// func NewStreamClientWithCreds(ctx context.Context, creds Credentials) (*StreamClient, error) {
// 	klog.V(6).Infof("NewStreamClientWithCreds ENTER\n")

// 	config := getDefaultConfig()

// 	restClient, err := NewRestClientWithCreds(ctx, creds)
// 	if err != nil {
// 		klog.V(2).Infof("NewRestClientWithCreds failed. Err: %v\n", err)
// 		klog.V(6).Infof("NewStreamClientWithCreds LEAVE\n")
// 		return nil, err
// 	}

// 	klog.V(6).Infof("IMPORTANT: Never print in production\n")
// 	klog.V(6).Infof("AppId: %s\n", restClient.creds.AppId)
// 	klog.V(6).Infof("AppSecret: %s\n", restClient.creds.AppSecret)

// 	// TODO: conversationID???
// 	id := uuid.New()
// 	klog.V(2).Infof("UUID: %s\n", id.String())

// 	streamPath := version.GetStreamingAPI(version.StreamPath, id)
// 	klog.V(6).Infof("IMPORTANT: Never print in production\n")
// 	klog.V(6).Infof("streamPath: %s\n", streamPath)

// 	wsClient, err := stream.NewWebSocketClient(symblPlatformHost, streamPath, restClient.auth.AccessToken)
// 	if err != nil {
// 		klog.V(2).Infof("stream.NewWebSocketClient failed. Err: %v\n", err)
// 		klog.V(6).Infof("NewStreamClient LEAVE\n")
// 		return nil, err
// 	}

// 	wsConnection := wsClient.Connect()
// 	if wsConnection == nil {
// 		klog.V(2).Infof("stream.NewWebSocketClient failed. Err: %v\n", err)
// 		klog.V(6).Infof("NewStreamClient LEAVE\n")
// 		return nil, common.ErrWebSocketInitializationFailed
// 	}

// 	err = wsClient.WriteJSON(config)
// 	if err != nil {
// 		klog.V(2).Infof("wsClient.WriteJSON failed. Err: %v\n", err)
// 		klog.V(6).Infof("NewStreamClient LEAVE\n")
// 		return nil, err
// 	}

// 	streamClient := &StreamClient{
// 		wsClient,
// 		restClient,
// 	}

// 	klog.V(2).Infof("NewStreamClientWithCreds Succeeded\n")
// 	klog.V(6).Infof("NewStreamClientWithCreds LEAVE\n")
// 	return streamClient, nil
// }
