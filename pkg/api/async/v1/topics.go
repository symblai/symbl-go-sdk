// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"context"
	"net/http"

	klog "k8s.io/klog/v2"

	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

type TopicResult struct {
	Topics []struct {
		Text       string   `json:"text"`
		Type       string   `json:"type"`
		Score      float64  `json:"score"`
		MessageIds []string `json:"messageIds"`
		Sentiment  struct {
			Polarity struct {
				Score float64 `json:"score"`
			} `json:"polarity"`
			Suggested string `json:"suggested"`
		} `json:"sentiment"`
		ParentRefs []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"parentRefs"`
	} `json:"topics"`
}

func (c *Client) GetTopics(ctx context.Context, conversationId string) (*TopicResult, error) {
	klog.V(6).Infof("async.GetTopics ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
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
	var topicResults TopicResult

	err = c.Client.Do(ctx, req, &topicResults)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("WaitForJobComplete failed. HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetTopics LEAVE\n")
			return nil, e
		}
	}

	klog.V(4).Infof("GET Topics succeeded\n")
	klog.V(6).Infof("async.GetTopics LEAVE\n")
	return &topicResults, nil
}
