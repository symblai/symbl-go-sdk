// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Async package for processing Async conversations
*/
package async

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

// GetMembers obtains members in a conversation
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
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.MembersURI, conversationId),
		c.getQueryParamFromContext(ctx))
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

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetMembers LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetMembers LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET Members succeeded\n")
	klog.V(6).Infof("async.GetMembers LEAVE\n")
	return &result, nil
}

// UpdateMember updates a member in a conversation
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
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.MemberURI, conversationId, member.ID),
		c.getQueryParamFromContext(ctx))
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

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.UpdateMember LEAVE\n")
				return err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.UpdateMember LEAVE\n")
		return err
	}

	klog.V(3).Infof("PUT Member succeeded\n")
	klog.V(6).Infof("async.UpdateMember LEAVE\n")
	return nil
}

// UpdateSpeakers updates a speaker in a conversation
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
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.SpeakersURI, conversationId),
		c.getQueryParamFromContext(ctx))
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

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
				return err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
		return err
	}

	klog.V(3).Infof("PUT UpdateSpeakers succeeded\n")
	klog.V(6).Infof("async.UpdateSpeakers LEAVE\n")
	return nil
}
