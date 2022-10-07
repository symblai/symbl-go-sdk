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

func (c *Client) GetTopics(ctx context.Context, conversationId string) (*interfaces.TopicResult, error) {
	klog.V(6).Infof("async.GetTopics ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetTopics LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.TopicsURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetTopics LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.TopicResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetTopics LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Topics succeeded\n")
	klog.V(6).Infof("async.GetTopics LEAVE\n")
	return &result, nil
}

func (c *Client) GetQuestions(ctx context.Context, conversationId string) (*interfaces.QuestionResult, error) {
	klog.V(6).Infof("async.GetQuestions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetQuestions LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.QuestionsURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetQuestions LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.QuestionResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetQuestions LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Questions succeeded\n")
	klog.V(6).Infof("async.GetQuestions LEAVE\n")
	return &result, nil
}

func (c *Client) GetFollowUps(ctx context.Context, conversationId string) (*interfaces.FollowUpResult, error) {
	klog.V(6).Infof("async.GetFollowUps ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetFollowUps LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.FollowUpsURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetFollowUps LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.FollowUpResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetFollowUps LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Follow Ups succeeded\n")
	klog.V(6).Infof("async.GetFollowUps LEAVE\n")
	return &result, nil
}

func (c *Client) GetEntities(ctx context.Context, conversationId string) (*interfaces.EntityResult, error) {
	klog.V(6).Infof("async.GetEntities ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetEntities LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.EntitiesURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetEntities LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.EntityResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetEntities LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Entities succeeded\n")
	klog.V(6).Infof("async.GetEntities LEAVE\n")
	return &result, nil
}

func (c *Client) GetActionItems(ctx context.Context, conversationId string) (*interfaces.ActionItemResult, error) {
	klog.V(6).Infof("async.GetActionItems ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetActionItems LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.ActionItemsURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetActionItems LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.ActionItemResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetActionItems LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Action Items succeeded\n")
	klog.V(6).Infof("async.GetActionItems LEAVE\n")
	return &result, nil
}

func (c *Client) GetMessages(ctx context.Context, conversationId string) (*interfaces.MessageResult, error) {
	klog.V(6).Infof("async.GetMessages ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetMessages LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.MessagesURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetMessages LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.MessageResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetMessages LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Messages succeeded\n")
	klog.V(6).Infof("async.GetMessages LEAVE\n")
	return &result, nil
}

func (c *Client) GetSummary(ctx context.Context, conversationId string) (*interfaces.SummaryResult, error) {
	klog.V(6).Infof("async.GetSummary ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetSummary LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.SummaryURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetSummary LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.SummaryResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetSummary LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Summary succeeded\n")
	klog.V(6).Infof("async.GetSummary LEAVE\n")
	return &result, nil
}

func (c *Client) GetAnalytics(ctx context.Context, conversationId string) (*interfaces.AnalytiscResult, error) {
	klog.V(6).Infof("async.GetAnalytics ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetAnalytics LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.AnalyticsURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetAnalytics LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.AnalytiscResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetAnalytics LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Analytics succeeded\n")
	klog.V(6).Infof("async.GetAnalytics LEAVE\n")
	return &result, nil
}

func (c *Client) GetTracker(ctx context.Context, conversationId string) (*interfaces.TrackerResult, error) {
	klog.V(6).Infof("async.GetTracker ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetTracker LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetAsyncAPI(version.TrackersURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetTracker LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.TrackerResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetTracker LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Tracker succeeded\n")
	klog.V(6).Infof("async.GetTracker LEAVE\n")
	return &result, nil
}
