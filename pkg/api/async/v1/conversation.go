// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"context"
	"fmt"
	"net/http"

	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

func (c *Client) GetConversations(ctx context.Context) (*asyncinterfaces.ConversationsResult, error) {
	klog.V(6).Infof("async.GetConversations ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.ConversationsURI),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetConversations LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.ConversationsResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
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

func (c *Client) GetConversation(ctx context.Context, conversationId string) (*asyncinterfaces.Conversation, error) {
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
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.ConversationURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetConversation LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.Conversation

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
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
