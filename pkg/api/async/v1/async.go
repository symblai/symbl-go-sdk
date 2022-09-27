// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"context"
	"net/http"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

const (
	UseDefaultWaitForCompletion int   = 0
	defaultWaitForCompletion    int   = 60
	defaultDelayBetweenCheck    int64 = 2

	JobStatusInProgress string = "in_progress"
	JobStatusComplete   string = "completed"
)

type Client struct {
	*symbl.Client
}

// Input parameters for calls
type WaitForJobStatusOpts struct {
	JobId         string `validate:"required"`
	WaitInSeconds int
}

// output structs
// JobStatus captures the API for getting status
type JobStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// JobConversation represents processing an Async API request
type JobConversation struct {
	JobID          string `json:"jobId"`
	ConversationID string `json:"conversationId"`
}

func New(client *symbl.Client) *Client {
	return &Client{client}
}

func (c *Client) PostFile(ctx context.Context, filePath string) (*JobConversation, error) {
	klog.V(6).Infof("async.PostFile ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	klog.V(2).Infof("filePath: %s\n", filePath)

	// send the file!
	var jobConvo JobConversation

	err := c.DoFile(ctx, filePath, &jobConvo)
	if e, ok := err.(*symbl.StatusError); ok {
		klog.V(2).Infof("DoFile failed. HTTP Code: %v\n", e.Resp.StatusCode)
		klog.V(6).Infof("async.PostFile LEAVE\n")
		return nil, e
	}

	klog.V(6).Infof("------------------------\n")
	klog.V(6).Infof("jobConvo:\n%v\n", jobConvo)
	klog.V(6).Infof("------------------------\n")

	klog.V(2).Infof("async.PostFile Succeeded\n")
	klog.V(6).Infof("async.PostFile LEAVE\n")
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
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.WaitForJobCompleteOnce ENTER\n")
		return false, err
	}

	klog.V(6).Infof("------------------------\n")
	klog.V(6).Infof("req:\n%v\n", req)
	klog.V(6).Infof("------------------------\n")

	// check the status
	var jobStatus JobStatus

	err = c.Client.Do(ctx, req, &jobStatus)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(2).Infof("WaitForJobComplete failed. HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.WaitForJobCompleteOnce LEAVE\n")
			return false, e
		}
	}

	klog.V(6).Infof("------------------------\n")
	klog.V(6).Infof("jobStatus:\n%v\n", jobStatus)
	klog.V(6).Infof("------------------------\n")

	complete := (jobStatus.Status == JobStatusComplete)

	klog.V(2).Infof("%s: %t", URI, complete)
	klog.V(6).Infof("async.WaitForJobCompleteOnce LEAVE\n")
	return complete, nil
}

func (c *Client) WaitForJobComplete(ctx context.Context, jobStatusOpts WaitForJobStatusOpts) (bool, error) {
	klog.V(6).Infof("async.WaitForJobComplete ENTER\n")

	// validate input
	v := validator.New()
	err := v.Struct(jobStatusOpts)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.Errorf("WaitForJobComplete validation failed: %v\n", e)
		}
		klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
		return false, err
	}

	if jobStatusOpts.WaitInSeconds < 0 {
		klog.Errorf("Invalid wait interval. Input: %d\n", jobStatusOpts.WaitInSeconds)
		klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
		return false, ErrInvalidWaitTime
	}

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	numOfLoops := float64(defaultWaitForCompletion) / float64(defaultDelayBetweenCheck)
	if jobStatusOpts.WaitInSeconds != UseDefaultWaitForCompletion {
		numOfLoops = float64(jobStatusOpts.WaitInSeconds) / float64(defaultDelayBetweenCheck)
		klog.V(4).Infof("User provided jobStatusOpts.WaitInSeconds\n")
	}
	klog.V(5).Infof("numOfLoops: %f\n", numOfLoops)
	klog.V(5).Infof("WaitInSeconds: %f\n", jobStatusOpts.WaitInSeconds)

	for i := 1; i <= int(numOfLoops); i++ {
		// delay on subsequent calls
		if i > 1 {
			klog.V(2).Info("Sleep for retry...\n")
			time.Sleep(time.Second * time.Duration(defaultDelayBetweenCheck))
		}

		// check for completion
		completed, err := c.WaitForJobCompleteOnce(ctx, jobStatusOpts.JobId)
		if err != nil {
			klog.Errorf("WaitForJobCompleteOnce failed. Err: %v\n", err)
			klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
			return false, err
		}
		if completed {
			klog.V(2).Info("WaitForJobCompleteOnce completed!\n")
			klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
			return true, nil
		}
	}

	klog.Errorf("job status timed out\n")
	klog.V(6).Infof("async.WaitForJobComplete LEAVE\n")
	return false, ErrJobStatusTimeout
}
