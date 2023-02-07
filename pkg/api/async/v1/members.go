// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

func (c *Client) GetMembers(ctx context.Context, conversationId string) (*asyncinterfaces.MembersResult, error) {
	klog.V(6).Infof("async.GetMembers ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetMembers LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.MembersURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetMembers LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.MembersResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetMembers LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Members succeeded\n")
	klog.V(6).Infof("async.GetMembers LEAVE\n")
	return &result, nil
}

func (c *Client) UpdateMember(ctx context.Context, conversationId string, member asyncinterfaces.Member) error {
	klog.V(6).Infof("async.UpdateMember ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.UpdateMember LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.MemberURI, conversationId, member.ID)
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(member)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.UpdateMember LEAVE\n")
		return err
	}

	// check the status
	err = c.Client.Do(ctx, req, nil)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.UpdateMember LEAVE\n")
			return err
		}
	}

	klog.V(3).Infof("PUT Member succeeded\n")
	klog.V(6).Infof("async.UpdateMember LEAVE\n")
	return nil
}

func (c *Client) UpdateSpeakers(ctx context.Context, conversationId string, speakers asyncinterfaces.UpdateSpeakerRequest) error {
	klog.V(6).Infof("async.UpdateSpeakers ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.SpeakersURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(speakers)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
		return err
	}

	// check the status
	err = c.Client.Do(ctx, req, nil)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
			return err
		}
	}

	klog.V(3).Infof("PUT UpdateSpeakers succeeded\n")
	klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
	return nil
}
