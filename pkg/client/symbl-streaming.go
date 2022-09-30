// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

import (
	"context"

	"github.com/google/uuid"
	klog "k8s.io/klog/v2"

	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	common "github.com/dvonthenen/symbl-go-sdk/pkg/client/common"
	streaming "github.com/dvonthenen/symbl-go-sdk/pkg/client/streaming"
)

const (
	symblPlatformHost string = "api.symbl.ai"

	requestStart string = "start_request"

	defaultConfidenceThreshold float64 = 0.7
	defaultSampleRateHertz     int     = 16000
	defaultUserID              string  = "user@email.com"
	defaultUserName            string  = "Jane Doe"
)

func getDefaultConfig() *StreamingConfig {
	config := &StreamingConfig{}

	config.Type = requestStart
	config.InsightTypes = []string{"topic", "question", "action_item"}
	config.Config.MeetingTitle = "my-meeting"
	config.Config.ConfidenceThreshold = defaultConfidenceThreshold
	config.Config.LanguageCode = "en-US"
	config.Config.SampleRateHertz = defaultSampleRateHertz
	config.Speaker.Name = defaultUserName
	config.Speaker.UserID = defaultUserID

	// TODO remove... the above seems to work
	// config.Type = requestStart
	// config.InsightTypes = []string{"topic", "question", "action_item"}
	// config.Config.ConfidenceThreshold = defaultConfidenceThreshold
	// config.Config.SpeechRecognition.SampleRateHertz = defaultSampleRateHertz
	// config.Speaker.Name = defaultUserName
	// config.Speaker.UserID = defaultUserID

	return config
}

type StreamingConfig struct {
	Type         string   `json:"type"`
	InsightTypes []string `json:"insightTypes"`
	Config       struct {
		MeetingTitle        string  `json:"meetingTitle"`
		ConfidenceThreshold float64 `json:"confidenceThreshold"`
		TimezoneOffset      int     `json:"timezoneOffset"`
		LanguageCode        string  `json:"languageCode"`
		SampleRateHertz     int     `json:"sampleRateHertz"`
	} `json:"config"`
	Speaker struct {
		UserID string `json:"userId"`
		Name   string `json:"name"`
	} `json:"speaker"`
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
	*streaming.WebSocketClient

	restClient *RestClient
}

// NewClient creates a new client on the Symbl.ai platform. The client authenticates with the
// server with APP_ID/APP_SECRET.
func NewStreamClient(ctx context.Context, callback streaming.WebSocketMessageCallback) (*StreamClient, error) {
	klog.V(6).Infof("NewStreamClient ENTER\n")

	config := getDefaultConfig()

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

	// create client
	creds := streaming.Credentials{
		Host:      symblPlatformHost,
		Channel:   streamPath,
		AccessKey: restClient.auth.AccessToken,
	}
	wsClient, err := streaming.NewWebSocketClient(creds, callback)
	if err != nil {
		klog.V(2).Infof("streaming.NewWebSocketClient failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, err
	}

	// establish connection
	wsConnection := wsClient.Connect()
	if wsConnection == nil {
		klog.V(2).Infof("streaming.NewWebSocketClient failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, common.ErrWebSocketInitializationFailed
	}

	// write Symbl config to Platform
	err = wsClient.Write(config)
	if err != nil {
		klog.V(2).Infof("wsClient.Write failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, err
	}

	// save client for return
	streamClient := &StreamClient{
		wsClient,
		restClient,
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

// 	wsClient, err := streaming.NewWebSocketClient(symblPlatformHost, streamPath, restClient.auth.AccessToken)
// 	if err != nil {
// 		klog.V(2).Infof("streaming.NewWebSocketClient failed. Err: %v\n", err)
// 		klog.V(6).Infof("NewStreamClient LEAVE\n")
// 		return nil, err
// 	}

// 	wsConnection := wsClient.Connect()
// 	if wsConnection == nil {
// 		klog.V(2).Infof("streaming.NewWebSocketClient failed. Err: %v\n", err)
// 		klog.V(6).Infof("NewStreamClient LEAVE\n")
// 		return nil, common.ErrWebSocketInitializationFailed
// 	}

// 	err = wsClient.Write(config)
// 	if err != nil {
// 		klog.V(2).Infof("wsClient.Write failed. Err: %v\n", err)
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
