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
	"net/url"
	"strings"

	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	common "github.com/symblai/symbl-go-sdk/pkg/api/common"
	version "github.com/symblai/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
)

// GetSummaryUI obtains a summary ui for conversation
func (c *Client) GetSummaryUI(ctx context.Context, conversationId string, uri string) (*asyncinterfaces.SummaryUIResult, error) {
	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		return nil, ErrInvalidInput
	}

	// text
	if len(uri) == 0 {
		request := asyncinterfaces.TextSummaryRequest{
			Name: "verbose-text-summary",
		}
		return c.GetTextSummaryUI(ctx, conversationId, request)
	}

	// url
	u, err := url.Parse(uri)
	if err != nil {
		klog.V(1).Infof("uri is invalid. Err: %v\n", err)
		return nil, err
	}

	pos := strings.LastIndex(u.Path, ".")
	if pos == -1 {
		err := ErrInvalidURIExtension
		klog.V(1).Infof("uri is invalid. Err: %v\n", err)
		return nil, err
	}

	extension := u.Path[pos+1:]
	klog.V(3).Infof("extension: %s\n", extension)

	// is audio?
	switch extension {
	case common.AudioTypeMP3:
		request := asyncinterfaces.AudioSummaryRequest{
			Name:     "audio-summary",
			AudioURL: uri,
		}
		return c.GetAudioSummaryUI(ctx, conversationId, request)
	case common.AudioTypeWav:
		request := asyncinterfaces.AudioSummaryRequest{
			Name:     "audio-summary",
			AudioURL: uri,
		}
		return c.GetAudioSummaryUI(ctx, conversationId, request)
	}

	// assume video
	request := asyncinterfaces.VideoSummaryRequest{
		Name:     "video-summary",
		VideoURL: uri,
	}
	return c.GetVideoSummaryUI(ctx, conversationId, request)
}

// GetSummaryUI obtains a summary ui for a text conversation
func (c *Client) GetTextSummaryUI(ctx context.Context, conversationId string, request asyncinterfaces.TextSummaryRequest) (*asyncinterfaces.SummaryUIResult, error) {
	klog.V(6).Infof("async.GetTextSummaryUI ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.SummaryURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	request.Name = "audio-summary"
	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.SummaryUIResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET TextSummaryUI succeeded\n")
	klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
	return &result, nil
}

// GetSummaryUI obtains a summary ui for an audio conversation
func (c *Client) GetAudioSummaryUI(ctx context.Context, conversationId string, request asyncinterfaces.AudioSummaryRequest) (*asyncinterfaces.SummaryUIResult, error) {
	klog.V(6).Infof("async.GetAudioSummaryUI ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.SummaryURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.SummaryUIResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET AudioSummaryUI succeeded\n")
	klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
	return &result, nil
}

// GetSummaryUI obtains a summary ui for a video conversation
func (c *Client) GetVideoSummaryUI(ctx context.Context, conversationId string, request asyncinterfaces.VideoSummaryRequest) (*asyncinterfaces.SummaryUIResult, error) {
	klog.V(6).Infof("async.GetVideoSummaryUI ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.SummaryURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.SummaryUIResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET VideoSummaryUI succeeded\n")
	klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
	return &result, nil
}

// GetInsightsListUiURI - Get insights list url for the logged in user
func (c *Client) GetInsightsListUiURI(ctx context.Context) (*asyncinterfaces.InsightsListUiResult, error) {
	klog.V(6).Infof("async.GetInsightsListUiUrl ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.InsightsListUiURI),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetInsightsListUiUrl LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.InsightsListUiResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetInsightsListUiUrl LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetInsightsListUiUrl LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET InsightsListUiUrl succeeded\n")
	klog.V(6).Infof("async.GetInsightsListUiUrl LEAVE\n")
	return &result, nil
}

// GetInsightsDetailsUiURI - Get insights details url for the logged in user by conversationId
func (c *Client) GetInsightsDetailsUiURI(ctx context.Context, conversationId string) (*asyncinterfaces.InsightsDetailsUiResult, error) {
	klog.V(6).Infof("async.GetInsightsDetailsUiURI ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetInsightsDetailsUiURI LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.InsightsDetailsUiURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetInsightsDetailsUiURI LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.InsightsDetailsUiResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetInsightsDetailsUiURI LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetInsightsDetailsUiURI LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET InsightsDetailsUiURI succeeded\n")
	klog.V(6).Infof("async.GetInsightsDetailsUiURI LEAVE\n")
	return &result, nil
}

// UpdateMediaUrlForInsightsDetailsUI updates the audio/video URL that will be played in Insights UI
func (c *Client) UpdateMediaUrlForInsightsDetailsUI(ctx context.Context, conversationId string, mediaUrl string) error {
	klog.V(6).Info("async.UpdateMediaUrlForInsightsDetailsUI ENTER")
	defer klog.V(6).Info("async.UpdateMediaUrlForInsightsDetailsUI LEAVE")

	// checks
	if conversationId == "" || mediaUrl == "" {
		klog.V(1).Info("conversationId or mediaUrl is empty")
		return ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.UpdateMediaURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s", URI)

	requestBody, err := json.Marshal(map[string]string{
		"url": mediaUrl,
	})
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v", err)
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", URI, bytes.NewBuffer(requestBody))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v", err)
		return err
	}

	// execute request
	err = c.Do(ctx, req, nil) // we don't need the response body
	if err != nil {
		klog.V(1).Infof("Request execution failed. Err: %v", err)
		return err
	}

	klog.V(3).Info("Update MediaUrl For InsightsUI succeeded")
	return nil
}
