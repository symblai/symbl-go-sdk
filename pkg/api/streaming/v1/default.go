// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Streaming package for processing real-time conversations
*/
package streaming

import (
	"encoding/json"
	"os"
	"strings"

	prettyjson "github.com/hokaccha/go-prettyjson"
	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1/interfaces"
)

// DefaultMessageRouter is a sample implementation that just prints insights to the console
type DefaultMessageRouter struct {
	TranscriptionDemo    bool
	TranscriptionDisable bool

	ChatmessageDemo    bool
	ChatmessageDisable bool

	AllDisable     bool
	InsightDisable bool
	EntityDisable  bool
	TopicDisable   bool
	TrackerDisable bool
	UserDisable    bool
}

// NewDefaultMessageRouter creates a new DefaultMessageRouter
func NewDefaultMessageRouter() *DefaultMessageRouter {
	var transcriptionDemoStr string
	if v := os.Getenv("SYMBL_TRANSCRIPTION_DEMO"); v != "" {
		klog.V(4).Info("SYMBL_TRANSCRIPTION_DEMO found")
		transcriptionDemoStr = v
	}
	var transcriptionDisableStr string
	if v := os.Getenv("SYMBL_TRANSCRIPTION_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_TRANSCRIPTION_DISABLE found")
		transcriptionDisableStr = v
	}
	var chatmessageDemoStr string
	if v := os.Getenv("SYMBL_CHAT_MESSAGE_DEMO"); v != "" {
		klog.V(4).Info("SYMBL_CHAT_MESSAGE_DEMO found")
		chatmessageDemoStr = v
	}
	var chatmessageDisableStr string
	if v := os.Getenv("SYMBL_CHAT_MESSAGE_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_CHAT_MESSAGE_DISABLE found")
		chatmessageDisableStr = v
	}

	var allDisableStr string
	if v := os.Getenv("SYMBL_ALL_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_ALL_DISABLE found")
		allDisableStr = v
	}
	var insightDisableStr string
	if v := os.Getenv("SYMBL_INSIGHT_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_INSIGHT_DISABLE found")
		insightDisableStr = v
	}
	var entityDisableStr string
	if v := os.Getenv("SYMBL_ENTITY_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_ENTITY_DISABLE found")
		entityDisableStr = v
	}
	var topicDisableStr string
	if v := os.Getenv("SYMBL_TOPIC_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_TOPIC_DISABLE found")
		topicDisableStr = v
	}
	var trackerDisableStr string
	if v := os.Getenv("SYMBL_TRACKER_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_TRACKER_DISABLE found")
		trackerDisableStr = v
	}
	var userDisableStr string
	if v := os.Getenv("SYMBL_USER_DISABLE"); v != "" {
		klog.V(4).Info("SYMBL_USER_DISABLE found")
		userDisableStr = v
	}

	transcriptionDemo := strings.EqualFold(strings.ToLower(transcriptionDemoStr), "true")
	transcriptionDisable := strings.EqualFold(strings.ToLower(transcriptionDisableStr), "true")
	chatmessageDemo := strings.EqualFold(strings.ToLower(chatmessageDemoStr), "true")
	chatmessageDisable := strings.EqualFold(strings.ToLower(chatmessageDisableStr), "true")

	allDisable := strings.EqualFold(strings.ToLower(allDisableStr), "true")
	insightDisable := strings.EqualFold(strings.ToLower(insightDisableStr), "true")
	entityDisable := strings.EqualFold(strings.ToLower(entityDisableStr), "true")
	topicDisable := strings.EqualFold(strings.ToLower(topicDisableStr), "true")
	trackerDisable := strings.EqualFold(strings.ToLower(trackerDisableStr), "true")
	userDisable := strings.EqualFold(strings.ToLower(userDisableStr), "true")

	return &DefaultMessageRouter{
		TranscriptionDemo:    transcriptionDemo,
		TranscriptionDisable: transcriptionDisable,
		ChatmessageDemo:      chatmessageDemo,
		ChatmessageDisable:   chatmessageDisable,
		AllDisable:           allDisable,
		InsightDisable:       insightDisable,
		EntityDisable:        entityDisable,
		TopicDisable:         topicDisable,
		TrackerDisable:       trackerDisable,
		UserDisable:          userDisable,
	}
}

// InitializedConversation implements the interface
func (dmr *DefaultMessageRouter) InitializedConversation(im *interfaces.InitializationMessage) error {
	data, err := json.Marshal(im)
	if err != nil {
		klog.V(1).Infof("InitializationMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nInitializationMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

// RecognitionResultMessage implements the streaming interface
func (dmr *DefaultMessageRouter) RecognitionResultMessage(rr *interfaces.RecognitionResult) error {
	if dmr.TranscriptionDisable {
		return nil // disable all output
	}

	if dmr.TranscriptionDemo {
		// if rr.Message.IsFinal {
		// 	klog.Infof("TRANSCRIPTION (FINAL): %s\n", rr.Message.Punctuated.Transcript)
		// } else {
		// 	for cnt, alternative := range rr.Message.Payload.Raw.Alternatives {
		// 		klog.Infof("TRANSCRIPTION (Alt: %d, %f): %s\n", cnt, alternative.Confidence, alternative.Transcript)
		// 	}
		// }
		klog.Infof("TRANSCRIPTION: %s\n", rr.Message.Punctuated.Transcript)

		return nil
	}

	data, err := json.Marshal(rr)
	if err != nil {
		klog.V(1).Infof("RecognitionResult json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nRecognitionResult Object DUMP:\n%s\n\n", prettyJson)

	return nil
}

// MessageResponseMessage implements the streaming interface
func (dmr *DefaultMessageRouter) MessageResponseMessage(mr *interfaces.MessageResponse) error {
	if dmr.ChatmessageDisable {
		return nil // disable chat output
	}

	if dmr.ChatmessageDemo {
		for _, msg := range mr.Messages {
			klog.Infof("\n\nChat Message [%s]: %s\n\n", msg.From.Name, msg.Payload.Content)
		}
		return nil
	}

	data, err := json.Marshal(mr)
	if err != nil {
		klog.V(1).Infof("MessageResponse json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nMessageResponse Object DUMP:\n%s\n\n", prettyJson)

	return nil
}

// InsightResponseMessage implements the streaming interface
func (dmr *DefaultMessageRouter) InsightResponseMessage(ir *interfaces.InsightResponse) error {
	if dmr.AllDisable || dmr.InsightDisable {
		return nil // disable all output
	}

	data, err := json.Marshal(ir)
	if err != nil {
		klog.V(1).Infof("InsightResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nInsightResponseMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

// TopicResponseMessage implements the streaming interface
func (dmr *DefaultMessageRouter) TopicResponseMessage(tr *interfaces.TopicResponse) error {
	if dmr.AllDisable || dmr.TopicDisable {
		return nil // disable all output
	}

	data, err := json.Marshal(tr)
	if err != nil {
		klog.V(1).Infof("TopicResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nTopicResponseMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

// TrackerResponseMessage implements the streaming interface
func (dmr *DefaultMessageRouter) TrackerResponseMessage(tr *interfaces.TrackerResponse) error {
	if dmr.AllDisable || dmr.TrackerDisable {
		return nil // disable all output
	}

	data, err := json.Marshal(tr)
	if err != nil {
		klog.V(1).Infof("TrackerResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nTrackerResponseMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

// EntityResponseMessage implements the streaming interface
func (dmr *DefaultMessageRouter) EntityResponseMessage(tr *interfaces.EntityResponse) error {
	if dmr.AllDisable || dmr.EntityDisable {
		return nil // disable all output
	}

	data, err := json.Marshal(tr)
	if err != nil {
		klog.V(1).Infof("EntityResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nEntityResponseMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

// TeardownConversation implements the streaming interface
func (dmr *DefaultMessageRouter) TeardownConversation(tm *interfaces.TeardownMessage) error {
	data, err := json.Marshal(tm)
	if err != nil {
		klog.V(1).Infof("TeardownConversation json.Marshal failed. Err: %v\n", err)
		return err
	}

	prettyJson, err := prettyjson.Format(data)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nTeardownConversation Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

// UserDefinedMessage implements the streaming interface
func (dmr *DefaultMessageRouter) UserDefinedMessage(byMsg []byte) error {
	if dmr.AllDisable || dmr.UserDisable {
		return nil // disable all output
	}

	prettyJson, err := prettyjson.Format(byMsg)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nUserDefinedMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

// UnhandledMessage implements the streaming interface
func (dmr *DefaultMessageRouter) UnhandledMessage(byMsg []byte) error {
	prettyJson, err := prettyjson.Format(byMsg)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nUnhandledMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}
