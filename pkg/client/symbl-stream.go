// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

import (
	"context"

	"github.com/google/uuid"
	klog "k8s.io/klog/v2"

	streaming "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	stream "github.com/dvonthenen/symbl-go-sdk/pkg/client/stream"
)

const (
	defaultConfidenceThreshold float64 = 0.7
	defaultSampleRateHertz     int     = 16000
	defaultUserID              string  = "user@email.com"
	defaultUserName            string  = "Jane Doe"
)

type StreamClient struct {
	*stream.WebSocketClient

	restClient     *RestClient
	symblStreaming stream.WebSocketMessageCallback
}

func getDefaultConfig() *StreamingConfig {
	config := &StreamingConfig{}

	config.Type = streaming.TypeRequestStart
	config.InsightTypes = []string{"topic", "question", "action_item", "follow_up"}
	config.Config.MeetingTitle = "my-meeting"
	config.Config.ConfidenceThreshold = defaultConfidenceThreshold
	// config.Config.TimezoneOffset = 480
	config.Config.SpeechRecognition.Encoding = "LINEAR16"
	config.Config.SpeechRecognition.SampleRateHertz = defaultSampleRateHertz
	config.Speaker.Name = defaultUserName
	config.Speaker.UserID = defaultUserID

	return config
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

	// klog.V(6).Infof("IMPORTANT: Never print in production\n")
	// klog.V(6).Infof("AppId: %s\n", restClient.creds.AppId)
	// klog.V(6).Infof("AppSecret: %s\n", restClient.creds.AppSecret)

	// generate unique id... not even sure why this is needed, but hey
	id := uuid.New()
	klog.V(4).Infof("UUID: %s\n", id.String())

	streamPath := version.GetStreamingAPI(version.StreamPath, id)
	klog.V(4).Infof("streamPath: %s\n", streamPath)

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
		return nil, ErrWebSocketInitializationFailed
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

func (sc *StreamClient) Stop() {
	// signal stop to Symbl Platform
	stopMsg := &streaming.MessageType{
		Type: streaming.TypeRequestStop,
	}

	err := sc.WriteJSON(stopMsg)
	if err != nil {
		klog.Errorf("wsClient.WriteJSON failed. Err: %v\n", err)
	}

	// stop websocket
	sc.WebSocketClient.Stop()
}
