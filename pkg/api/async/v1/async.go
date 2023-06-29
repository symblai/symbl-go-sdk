// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Async package for processing Async conversations
*/
package async

import (
	"context"
	"net/http"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	client "github.com/dvonthenen/symbl-go-sdk/pkg/client"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

const (
	defaultWaitForCompletion int64 = 120
	defaultDelayBetweenCheck int64 = 2
)

// Context switch for processing Async functionality
type Client struct {
	*client.RestClient
}

// New changes the context of the REST client to an Async client
func New(client *client.RestClient) *Client {
	return &Client{client}
}

// PostText posts text conversations to the platform
func (c *Client) PostText(ctx context.Context, messages []string) (*JobConversation, error) {
	textRequest := asyncinterfaces.AsyncTextRequest{}

	for _, message := range messages {
		textRequest.Messages = append(textRequest.Messages, asyncinterfaces.TextMessage{
			Payload: asyncinterfaces.Payload{
				Content: message,
			},
			From:     nil,
			Duration: nil,
		})
	}

	return c.PostTextWithOptions(ctx, textRequest)
}

// PostAppendText appends text conversations to the platform
func (c *Client) PostAppendText(ctx context.Context, conversationId string, messages []string) (*JobConversation, error) {
	textRequest := asyncinterfaces.AsyncTextRequest{}

	for _, message := range messages {
		textRequest.Messages = append(textRequest.Messages, asyncinterfaces.TextMessage{
			Payload: asyncinterfaces.Payload{
				Content: message,
			},
			From:     nil,
			Duration: nil,
		})
	}

	return c.PostAppendTextWithOptions(ctx, conversationId, textRequest)
}

// PostFile posts a file containing a conversations to the platform
func (c *Client) PostFile(ctx context.Context, filePath string) (*JobConversation, error) {
	ufRequest := asyncinterfaces.AsyncURLFileRequest{}
	return c.PostFileWithOptions(ctx, filePath, ufRequest)
}

// PostURL posts a URL pointing to a conversations to the platform
func (c *Client) PostURL(ctx context.Context, url string) (*JobConversation, error) {
	ufRequest := asyncinterfaces.AsyncURLFileRequest{
		URL: url,
	}
	return c.PostURLWithOptions(ctx, ufRequest)
}

// PostURLWithOptions posts a URL pointing to a conversations to the platform with given options
func (c *Client) PostURLWithOptions(ctx context.Context, ufRequest asyncinterfaces.AsyncURLFileRequest) (*JobConversation, error) {
	klog.V(6).Infof("async.PostURLWithOptions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	klog.V(3).Infof("url: %s\n", ufRequest.URL)

	// send the URL!
	var jobConvo JobConversation

	err := c.DoURLWithOptions(ctx, ufRequest, &jobConvo)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.PostURLWithOptions LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.PostURLWithOptions LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("async.PostURLWithOptions Succeeded\n")
	klog.V(6).Infof("async.PostURLWithOptions LEAVE\n")
	return &jobConvo, nil
}

// PostFileWithOptions posts a File pointing to a conversations to the platform with given options
func (c *Client) PostFileWithOptions(ctx context.Context, filePath string, ufRequest asyncinterfaces.AsyncURLFileRequest) (*JobConversation, error) {
	klog.V(6).Infof("async.PostFileWithOptions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	klog.V(3).Infof("filePath: %s\n", filePath)

	// send the file!
	var jobConvo JobConversation

	err := c.DoFileWithOptions(ctx, filePath, ufRequest, &jobConvo)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.PostFileWithOptions LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.PostFileWithOptions LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("async.PostFileWithOptions Succeeded\n")
	klog.V(6).Infof("async.PostFileWithOptions LEAVE\n")
	return &jobConvo, nil
}

// WaitForJobCompleteOnce is a convenience wrapper for checking if the platform is finished processing a conversation
func (c *Client) WaitForJobCompleteOnce(ctx context.Context, jobId string) (bool, error) {
	klog.V(6).Infof("async.WaitForJobCompleteOnce ENTER\n")

	// checks
	if jobId == "" {
		return false, ErrInvalidInput
	}
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := version.GetAsyncAPI(version.JobStatusURI, jobId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.WaitForJobCompleteOnce ENTER\n")
		return false, err
	}

	// check the status
	var jobStatus JobStatus

	err = c.Client.Do(ctx, req, &jobStatus)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.WaitForJobCompleteOnce LEAVE\n")
				return false, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.WaitForJobCompleteOnce LEAVE\n")
		return false, err
	}

	complete := (jobStatus.Status == JobStatusComplete)

	klog.V(3).Infof("%s: %t", URI, complete)
	klog.V(6).Infof("async.WaitForJobCompleteOnce LEAVE\n")
	return complete, nil
}

// PostTextWithOptions posts text conversation to the platform with given options
func (c *Client) PostTextWithOptions(ctx context.Context, textRequest asyncinterfaces.AsyncTextRequest) (*JobConversation, error) {
	klog.V(6).Infof("async.PostTextWithOptions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// send the URL!
	var jobConvo JobConversation

	err := c.DoTextWithOptions(ctx, textRequest, &jobConvo)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.PostTextWithOptions LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.PostTextWithOptions LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("async.PostTextWithOptions Succeeded\n")
	klog.V(6).Infof("async.PostTextWithOptions LEAVE\n")
	return &jobConvo, nil
}

// PostTextWithOptions appends text conversation to the platform with given options
func (c *Client) PostAppendTextWithOptions(ctx context.Context, conversationId string, textRequest asyncinterfaces.AsyncTextRequest) (*JobConversation, error) {
	klog.V(6).Infof("async.PostAppendTextWithOptions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// send the URL!
	var jobConvo JobConversation

	err := c.DoAppendTextWithOptions(ctx, conversationId, textRequest, &jobConvo)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.PostAppendTextWithOptions LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.PostAppendTextWithOptions LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("async.PostAppendTextWithOptions Succeeded\n")
	klog.V(6).Infof("async.PostAppendTextWithOptions LEAVE\n")
	return &jobConvo, nil
}

// WaitForJobComplete is a loop wrapping the WaitForJobCompleteOnce call
func (c *Client) WaitForJobComplete(ctx context.Context, jobStatusOpts asyncinterfaces.WaitForJobStatusOpts) (bool, error) {
	klog.V(6).Infof("async.WaitForJobComplete ENTER\n")

	// validate input
	v := validator.New()
	err := v.Struct(jobStatusOpts)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("WaitForJobComplete validation failed: %v\n", e)
		}
		klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
		return false, err
	}

	// is valid?
	if jobStatusOpts.TotalWaitInSeconds <= 0 {
		jobStatusOpts.TotalWaitInSeconds = defaultWaitForCompletion
		klog.V(3).Infof("Use fauled wait interval. Input: %d\n", jobStatusOpts.TotalWaitInSeconds)
	}

	// is valid?
	if jobStatusOpts.WaitInSeconds <= 0 {
		jobStatusOpts.WaitInSeconds = defaultDelayBetweenCheck
		klog.V(3).Infof("Use fauled wait interval. Input: %d\n", jobStatusOpts.WaitInSeconds)
	}

	// do only 1 retry in the loop
	if jobStatusOpts.WaitInSeconds >= jobStatusOpts.TotalWaitInSeconds {
		jobStatusOpts.WaitInSeconds = jobStatusOpts.TotalWaitInSeconds
		jobStatusOpts.TotalWaitInSeconds = jobStatusOpts.TotalWaitInSeconds * 2
		klog.V(3).Infof("Invalid Input. Changing to TotalWaitInSeconds: %d, : %d\n", jobStatusOpts.TotalWaitInSeconds, jobStatusOpts.WaitInSeconds)
	}

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	numOfLoops := jobStatusOpts.TotalWaitInSeconds / jobStatusOpts.WaitInSeconds
	klog.V(4).Infof("numOfLoops: %d\n", numOfLoops)
	klog.V(4).Infof("WaitInSeconds: %d\n", jobStatusOpts.WaitInSeconds)

	for i := 0; i < int(numOfLoops); i++ {
		// delay on subsequent calls
		if i > 0 {
			klog.V(4).Infof("Sleep for retry #%d...\n", i)
			time.Sleep(time.Second * time.Duration(jobStatusOpts.WaitInSeconds))
		}

		// check for completion
		completed, err := c.WaitForJobCompleteOnce(ctx, jobStatusOpts.JobId)
		if err != nil {
			klog.V(1).Infof("WaitForJobCompleteOnce failed. Err: %v\n", err)
			klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
			return false, err
		}
		if completed {
			klog.V(3).Info("WaitForJobCompleteOnce completed!\n")
			klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
			return true, nil
		}
	}

	klog.V(1).Infof("job status timed out\n")
	klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
	return false, ErrJobStatusTimeout
}
