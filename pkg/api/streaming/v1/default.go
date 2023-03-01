// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import (
	"encoding/json"
	"os"
	"strings"

	prettyjson "github.com/hokaccha/go-prettyjson"
	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1/interfaces"
)

type DefaultMessageRouter struct {
	TranscriptionDemo    bool
	TranscriptionDisable bool

	ChatmessageDemo    bool
	ChatmessageDisable bool
}

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

	transcriptionDemo := strings.EqualFold(strings.ToLower(transcriptionDemoStr), "true")
	transcriptionDisable := strings.EqualFold(strings.ToLower(transcriptionDisableStr), "true")
	chatmessageDemo := strings.EqualFold(strings.ToLower(chatmessageDemoStr), "true")
	chatmessageDisable := strings.EqualFold(strings.ToLower(chatmessageDisableStr), "true")

	return &DefaultMessageRouter{
		TranscriptionDemo:    transcriptionDemo,
		TranscriptionDisable: transcriptionDisable,
		ChatmessageDemo:      chatmessageDemo,
		ChatmessageDisable:   chatmessageDisable,
	}
}

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

func (dmr *DefaultMessageRouter) RecognitionResultMessage(rr *interfaces.RecognitionResult) error {
	if dmr.TranscriptionDisable {
		return nil // disable all output
	}

	if dmr.TranscriptionDemo {
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

func (dmr *DefaultMessageRouter) InsightResponseMessage(ir *interfaces.InsightResponse) error {
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

func (dmr *DefaultMessageRouter) TopicResponseMessage(tr *interfaces.TopicResponse) error {
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
func (dmr *DefaultMessageRouter) TrackerResponseMessage(tr *interfaces.TrackerResponse) error {
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

func (dmr *DefaultMessageRouter) EntityResponseMessage(tr *interfaces.EntityResponse) error {
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

func (dmr *DefaultMessageRouter) UserDefinedMessage(byMsg []byte) error {
	prettyJson, err := prettyjson.Format(byMsg)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nUserDefinedMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}

func (dmr *DefaultMessageRouter) UnhandledMessage(byMsg []byte) error {
	prettyJson, err := prettyjson.Format(byMsg)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nUnhandledMessage Object DUMP:\n%s\n\n", prettyJson)
	return nil
}
