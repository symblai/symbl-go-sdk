// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package symbl

import (
	"context"

	"github.com/google/uuid"
	klog "k8s.io/klog/v2"

	streaming "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	stream "github.com/dvonthenen/symbl-go-sdk/pkg/client/stream"
)

const (
	defaultUserID   string = "user@email.com"
	defaultUserName string = "Jane Doe"
)

// GetDefaultConfig returns a minimal Symbl Config for the Websocket Interface
func GetDefaultConfig() *interfaces.StreamingConfig {
	config := &interfaces.StreamingConfig{}

	config.Type = streaming.TypeRequestStart
	config.InsightTypes = []string{"topic", "question", "action_item", "follow_up"}
	config.Speaker.Name = defaultUserName
	config.Speaker.UserID = defaultUserID

	return config
}

// NewStreamClientWithDefaults creates a Symbl Streaming Client with the default config and
// uses the APP_ID/APP_SECRET environment variables for authentication.
func NewStreamClientWithDefaults(ctx context.Context) (*StreamClient, error) {
	options := StreamingOptions{
		SymblConfig: GetDefaultConfig(),
		Callback:    streaming.NewDefaultMessageRouter(),
	}
	return NewStreamClient(ctx, options)
}

// NewStreamClient creates a Symbl Streaming Client with the provided StreamingOptions and
// uses the APP_ID/APP_SECRET environment variables for authentication.
func NewStreamClient(ctx context.Context, options StreamingOptions) (*StreamClient, error) {
	klog.V(6).Infof("NewStreamClient ENTER\n")

	if options.SymblConfig == nil {
		klog.V(1).Infof("Config is null\n")
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, ErrInvalidInput
	}

	// create rest client
	restClient, err := NewRestClient(ctx)
	if err != nil {
		klog.V(1).Infof("NewRestClient failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, err
	}

	// is there a proxy?
	streamingAddress := streaming.SymblPlatformHost
	if len(options.SymblEndpoint) > 0 {
		streamingAddress = options.SymblEndpoint
		klog.V(3).Infof("[OVERRIDE] Symbl Address: %s\n", streamingAddress)
	}

	// generate unique conversationId
	conversationId := options.UUID
	if len(options.UUID) == 0 {
		conversationId = uuid.New().String()
	}
	klog.V(4).Infof("UUID: %s\n", conversationId)

	streamPath := version.GetStreamingAPI(version.StreamPath, conversationId)
	klog.V(4).Infof("streamPath: %s\n", streamPath)

	// init symbl websocket message router
	symblStreaming := streaming.New(options.Callback)

	// create client
	creds := stream.Credentials{
		Host:            streamingAddress,
		Channel:         streamPath,
		AccessKey:       restClient.auth.AccessToken,
		RedirectService: options.RedirectService,
		SkipServerAuth:  options.SkipServerAuth,
	}
	wsClient, err := stream.NewWebSocketClient(ctx, creds, symblStreaming)
	if err != nil {
		klog.V(1).Infof("stream.NewWebSocketClient failed. Err: %v\n", err)
		klog.V(6).Infof("NewStreamClient LEAVE\n")
		return nil, err
	}

	// save client for return
	streamClient := &StreamClient{
		wsClient,
		conversationId,
		restClient,
		symblStreaming,
		&options,
	}

	klog.V(3).Infof("NewStreamClient Succeeded\n")
	klog.V(6).Infof("NewStreamClient LEAVE\n")
	return streamClient, nil
}

// Start begins the Symbl Platform Websocket Protocol by sending the "start_request" message
func (sc *StreamClient) Start() error {
	klog.V(6).Infof("Start ENTER\n")

	// set streaming type
	if sc.options.SymblConfig == nil {
		klog.V(1).Infof("Config is null\n")
		klog.V(6).Infof("Start LEAVE\n")
		return ErrInvalidInput
	}
	sc.options.SymblConfig.Type = streaming.TypeRequestStart

	// establish connection
	wsConnection := sc.Connect()
	if wsConnection == nil {
		klog.V(1).Infof("wsClient.Connect failed\n")
		klog.V(6).Infof("Start LEAVE\n")
		return ErrWebSocketInitializationFailed
	}

	// write Symbl config to Platform
	err := sc.WriteJSON(sc.options.SymblConfig)
	if err != nil {
		klog.V(1).Infof("wsClient.WriteJSON failed. Err: %v\n", err)
		klog.V(6).Infof("Start LEAVE\n")
		return err
	}

	klog.V(3).Infof("Start Succeeded\n")
	klog.V(6).Infof("Start LEAVE\n")
	return nil
}

// GetConversationId returns the Symbl Conversation ID for this Real-Time Streaming session.
func (sc *StreamClient) GetConversationId() string {
	return sc.uuid
}

// Stop closes the Websocket connection cleanly by sending "stop_request" message to the Symbl Platform.
func (sc *StreamClient) Stop() {
	// signal stop to Symbl Platform
	stopMsg := &streaming.MessageType{
		Type: streaming.TypeRequestStop,
	}

	err := sc.WriteJSON(stopMsg)
	if err != nil {
		klog.V(1).Infof("wsClient.WriteJSON failed. Err: %v\n", err)
	}

	// stop websocket
	sc.WebSocketClient.Stop()
}
