// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
Async package for processing Async conversations
*/
package async

import (
	"context"
	"fmt"
	"net/http"

	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/symblai/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
)

// Get Call Score Status By Id
func (c *Client) GetCallScoreStatusById(ctx context.Context, conversationId string) (*asyncinterfaces.CallScoreStatusResult, error) {
	klog.V(6).Infof("async.GetCallScoreStatusById ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetCallScoreStatusById LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.CallScoreStatusURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetCallScoreStatusById LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.CallScoreStatusResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetCallScoreStatusById LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetCallScoreStatusById LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET CallScoreStatus succeeded\n")
	klog.V(6).Infof("async.GetCallScoreStatusById LEAVE\n")
	return &result, nil
}

// Get Insight Status By Id
func (c *Client) GetInsightStatusById(ctx context.Context, conversationId string) (*asyncinterfaces.InsightStatusResult, error) {
	klog.V(6).Infof("async.GetInsightStatusById ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetInsightStatusById LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.InsightStatusURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetInsightStatusById LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.InsightStatusResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetInsightStatusById LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetInsightStatusById LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET InsightStatus succeeded\n")
	klog.V(6).Infof("async.GetInsightStatusById LEAVE\n")
	return &result, nil
}

// Get Call Score
func (c *Client) GetCallScore(ctx context.Context, conversationId string) (*asyncinterfaces.CallScoreResult, error) {
	klog.V(6).Infof("async.GetCallScore ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetCallScore LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.CallScoreURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetCallScore LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.CallScoreResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetCallScore LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetCallScore LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET CallScore succeeded\n")
	klog.V(6).Infof("async.GetCallScore LEAVE\n")
	return &result, nil
}
