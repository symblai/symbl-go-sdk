// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"context"
	"net/http"

	klog "k8s.io/klog/v2"

	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
)

func (c *Client) GetConversations(ctx context.Context) (*interfaces.ConversationsResult, error) {
	klog.V(6).Infof("async.GetConversations ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := version.GetAsyncAPI(version.ConversationsURI)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetConversations LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.ConversationsResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetConversations LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Conversations succeeded\n")
	klog.V(6).Infof("async.GetConversations LEAVE\n")
	return &result, nil
}

func (c *Client) GetConversation(ctx context.Context, conversationId string) (*interfaces.Conversation, error) {
	klog.V(6).Infof("async.GetConversation ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetConversation LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.ConversationURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetConversation LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.Conversation

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetConversations LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Conversations succeeded\n")
	klog.V(6).Infof("async.GetConversations LEAVE\n")
	return &result, nil
}
