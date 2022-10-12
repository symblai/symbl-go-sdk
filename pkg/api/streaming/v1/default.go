// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import (
	"encoding/json"

	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1/interfaces"
)

type DefaultMessageRouter struct{}

func NewDefaultMessageRouter() *DefaultMessageRouter {
	return &DefaultMessageRouter{}
}

func (dmr *DefaultMessageRouter) RecognitionResultMessage(rr *interfaces.RecognitionResult) error {
	data, err := json.Marshal(rr)
	if err != nil {
		klog.V(1).Infof("RecognitionResult json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nRecognitionResult Object DUMP:\n%v\n\n", string(data))
	return nil
}

func (dmr *DefaultMessageRouter) MessageResponseMessage(mr *interfaces.MessageResponse) error {
	data, err := json.Marshal(mr)
	if err != nil {
		klog.V(1).Infof("MessageResponse json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nMessageResponse Object DUMP:\n%v\n\n", string(data))
	return nil
}

func (dmr *DefaultMessageRouter) InsightResponseMessage(ir *interfaces.InsightResponse) error {
	data, err := json.Marshal(ir)
	if err != nil {
		klog.V(1).Infof("InsightResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nInsightResponseMessage Object DUMP:\n%v\n\n", string(data))
	return nil
}

func (dmr *DefaultMessageRouter) TopicResponseMessage(tr *interfaces.TopicResponse) error {
	data, err := json.Marshal(tr)
	if err != nil {
		klog.V(1).Infof("TopicResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nTopicResponseMessage Object DUMP:\n%v\n\n", string(data))
	return nil
}
func (dmr *DefaultMessageRouter) TrackerResponseMessage(tr *interfaces.TrackerResponse) error {
	data, err := json.Marshal(tr)
	if err != nil {
		klog.V(1).Infof("TrackerResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("\n\nTrackerResponseMessage Object DUMP:\n%v\n\n", string(data))
	return nil
}

func (dmr *DefaultMessageRouter) UnhandledMessage(byMsg []byte) error {
	klog.Infof("\n\nUnhandledMessage Object DUMP:\n%v\n\n", string(byMsg))
	return nil
}
