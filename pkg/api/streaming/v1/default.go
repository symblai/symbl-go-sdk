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
		klog.Errorf("RecognitionResult json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("RecognitionResult Object DUMP:\n%v\n", string(data))
	return nil
}

func (dmr *DefaultMessageRouter) MessageResponseMessage(mr *interfaces.MessageResponse) error {
	data, err := json.Marshal(mr)
	if err != nil {
		klog.Errorf("MessageResponse json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("MessageResponse Object DUMP:\n%v\n", string(data))
	return nil
}

func (dmr *DefaultMessageRouter) InsightResponseMessage(ir *interfaces.InsightResponse) error {
	data, err := json.Marshal(ir)
	if err != nil {
		klog.Errorf("InsightResponseMessage json.Marshal failed. Err: %v\n", err)
		return err
	}

	klog.Infof("InsightResponseMessage Object DUMP:\n%v\n", string(data))
	return nil
}

func (dmr *DefaultMessageRouter) UnhandledMessage(byMsg []byte) error {
	klog.Infof("UnhandledMessage Object DUMP:\n%v\n", string(byMsg))
	return nil
}
