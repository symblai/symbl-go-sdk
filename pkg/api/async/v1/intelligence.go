// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

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

func (c *Client) GetTopics(ctx context.Context, conversationId string) (*asyncinterfaces.TopicResult, error) {
	klog.V(6).Infof("async.GetTopics ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetTopics LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.TopicsURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetTopics LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.TopicResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetTopics LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Topics succeeded\n")
	klog.V(6).Infof("async.GetTopics LEAVE\n")
	return &result, nil
}

func (c *Client) GetQuestions(ctx context.Context, conversationId string) (*asyncinterfaces.QuestionResult, error) {
	klog.V(6).Infof("async.GetQuestions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetQuestions LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.QuestionsURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetQuestions LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.QuestionResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetQuestions LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Questions succeeded\n")
	klog.V(6).Infof("async.GetQuestions LEAVE\n")
	return &result, nil
}

func (c *Client) GetFollowUps(ctx context.Context, conversationId string) (*asyncinterfaces.FollowUpResult, error) {
	klog.V(6).Infof("async.GetFollowUps ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetFollowUps LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.FollowUpsURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetFollowUps LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.FollowUpResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetFollowUps LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Follow Ups succeeded\n")
	klog.V(6).Infof("async.GetFollowUps LEAVE\n")
	return &result, nil
}

func (c *Client) GetEntities(ctx context.Context, conversationId string) (*asyncinterfaces.EntityResult, error) {
	klog.V(6).Infof("async.GetEntities ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetEntities LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.EntitiesURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetEntities LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.EntityResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetEntities LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Entities succeeded\n")
	klog.V(6).Infof("async.GetEntities LEAVE\n")
	return &result, nil
}

func (c *Client) GetActionItems(ctx context.Context, conversationId string) (*asyncinterfaces.ActionItemResult, error) {
	klog.V(6).Infof("async.GetActionItems ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetActionItems LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.ActionItemsURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetActionItems LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.ActionItemResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetActionItems LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Action Items succeeded\n")
	klog.V(6).Infof("async.GetActionItems LEAVE\n")
	return &result, nil
}

func (c *Client) GetMessages(ctx context.Context, conversationId string) (*asyncinterfaces.MessageResult, error) {
	klog.V(6).Infof("async.GetMessages ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetMessages LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.MessagesURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetMessages LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.MessageResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetMessages LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Messages succeeded\n")
	klog.V(6).Infof("async.GetMessages LEAVE\n")
	return &result, nil
}

func (c *Client) GetSummary(ctx context.Context, conversationId string) (*asyncinterfaces.SummaryResult, error) {
	klog.V(6).Infof("async.GetSummary ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetSummary LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.SummaryURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetSummary LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.SummaryResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetSummary LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Summary succeeded\n")
	klog.V(6).Infof("async.GetSummary LEAVE\n")
	return &result, nil
}

func (c *Client) GetAnalytics(ctx context.Context, conversationId string) (*asyncinterfaces.AnalyticsResult, error) {
	klog.V(6).Infof("async.GetAnalytics ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetAnalytics LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.AnalyticsURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetAnalytics LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.AnalyticsResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetAnalytics LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Analytics succeeded\n")
	klog.V(6).Infof("async.GetAnalytics LEAVE\n")
	return &result, nil
}

func (c *Client) GetTracker(ctx context.Context, conversationId string) (*asyncinterfaces.TrackerResult, error) {
	klog.V(6).Infof("async.GetTracker ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetTracker LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.TrackersURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetTracker LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.TrackerResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetTracker LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Tracker succeeded\n")
	klog.V(6).Infof("async.GetTracker LEAVE\n")
	return &result, nil
}
