// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"context"
	"net/http"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

const (
	defaultWaitForCompletion int   = 60
	defaultDelayBetweenCheck int64 = 2
)

type Client struct {
	*symbl.RestClient
}

func New(client *symbl.RestClient) *Client {
	return &Client{client}
}

func (c *Client) PostFile(ctx context.Context, filePath string) (*JobConversation, error) {
	options := interfaces.AsyncOptions{}
	return c.PostFileWithOptions(ctx, filePath, options)
}

func (c *Client) PostURL(ctx context.Context, url string) (*JobConversation, error) {
	options := interfaces.AsyncOptions{
		URL: url,
	}
	return c.PostURLWithOptions(ctx, options)
}

func (c *Client) PostURLWithOptions(ctx context.Context, options interfaces.AsyncOptions) (*JobConversation, error) {
	klog.V(6).Infof("async.PostURLWithOptions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	klog.V(3).Infof("url: %s\n", options.URL)

	// send the URL!
	var jobConvo JobConversation

	err := c.DoURLWithOptions(ctx, options, &jobConvo)
	if e, ok := err.(*symbl.StatusError); ok {
		klog.V(1).Infof("DoURL failed. HTTP Code: %v\n", e.Resp.StatusCode)
		klog.V(6).Infof("async.PostURLWithOptions LEAVE\n")
		return nil, e
	}

	klog.V(3).Infof("async.PostURLWithOptions Succeeded\n")
	klog.V(6).Infof("async.PostURLWithOptions LEAVE\n")
	return &jobConvo, nil
}

func (c *Client) PostFileWithOptions(ctx context.Context, filePath string, options interfaces.AsyncOptions) (*JobConversation, error) {
	klog.V(6).Infof("async.PostFileWithOptions ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	klog.V(3).Infof("filePath: %s\n", filePath)

	// send the file!
	var jobConvo JobConversation

	err := c.DoFileWithOptions(ctx, filePath, options, &jobConvo)
	if e, ok := err.(*symbl.StatusError); ok {
		klog.V(1).Infof("DoFile failed. HTTP Code: %v\n", e.Resp.StatusCode)
		klog.V(6).Infof("async.PostFileWithOptions LEAVE\n")
		return nil, e
	}

	klog.V(3).Infof("async.PostFileWithOptions Succeeded\n")
	klog.V(6).Infof("async.PostFileWithOptions LEAVE\n")
	return &jobConvo, nil
}

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

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.WaitForJobCompleteOnce LEAVE\n")
			return false, err
		}
	}

	complete := (jobStatus.Status == JobStatusComplete)

	klog.V(3).Infof("%s: %t", URI, complete)
	klog.V(6).Infof("async.WaitForJobCompleteOnce LEAVE\n")
	return complete, nil
}

func (c *Client) WaitForJobComplete(ctx context.Context, jobStatusOpts interfaces.WaitForJobStatusOpts) (bool, error) {
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

	if jobStatusOpts.WaitInSeconds < 0 {
		klog.V(1).Infof("Invalid wait interval. Input: %d\n", jobStatusOpts.WaitInSeconds)
		klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
		return false, ErrInvalidWaitTime
	}

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	numOfLoops := float64(defaultWaitForCompletion) / float64(defaultDelayBetweenCheck)
	if jobStatusOpts.WaitInSeconds != interfaces.UseDefaultWaitForCompletion {
		numOfLoops = float64(jobStatusOpts.WaitInSeconds) / float64(defaultDelayBetweenCheck)
		klog.V(4).Infof("User provided jobStatusOpts.WaitInSeconds\n")
	}
	klog.V(5).Infof("numOfLoops: %f\n", numOfLoops)
	klog.V(5).Infof("WaitInSeconds: %f\n", jobStatusOpts.WaitInSeconds)

	for i := 1; i <= int(numOfLoops); i++ {
		// delay on subsequent calls
		if i > 1 {
			klog.V(4).Info("Sleep for retry...\n")
			time.Sleep(time.Second * time.Duration(defaultDelayBetweenCheck))
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
