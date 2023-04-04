// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

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

	asyncinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	common "github.com/dvonthenen/symbl-go-sdk/pkg/api/common"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

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
	case common.AudioTypeMpeg:
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
	URI := fmt.Sprintf("%s?%s",
		version.GetAsyncAPI(version.SummaryURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

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

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET TextSummaryUI succeeded\n")
	klog.V(6).Infof("async.GetTextSummaryUI LEAVE\n")
	return &result, nil
}

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
	URI := fmt.Sprintf("%s?%s",
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

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET AudioSummaryUI succeeded\n")
	klog.V(6).Infof("async.GetAudioSummaryUI LEAVE\n")
	return &result, nil
}

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
	URI := fmt.Sprintf("%s?%s",
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

	if e, ok := err.(*interfaces.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET VideoSummaryUI succeeded\n")
	klog.V(6).Infof("async.GetVideoSummaryUI LEAVE\n")
	return &result, nil
}
