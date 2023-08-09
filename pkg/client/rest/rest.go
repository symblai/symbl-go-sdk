// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	common "github.com/symblai/symbl-go-sdk/pkg/api/common"
	version "github.com/symblai/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
	simple "github.com/symblai/symbl-go-sdk/pkg/client/simple"
)

// New allocated a REST client
func New() *Client {
	c := Client{
		Client: simple.New(),
	}
	return &c
}

// SetAuthorization sets an authorization token to make API calls to a given platform
func (c *Client) SetAuthorization(auth *AccessToken) {
	c.auth = auth
}

// TODO: Multipart file upload is not supported by the platform
// func (c *Client) DoMultiPartFile(ctx context.Context, filePath string, resBody interface{}) error {
// 	klog.V(6).Infof("rest.DoMultiPartFile ENTER\n")

// 	// checks
// 	fileInfo, err := os.Stat(filePath)
// 	if err != nil || errors.Is(err, os.ErrNotExist) {
// 		klog.V(1).Infof("File %s does not exist. Err : %v\n", filePath, err)
// 		klog.V(6).Infof("rest.DoMultiPartFile LEAVE\n")
// 		return err
// 	}

// 	if fileInfo.IsDir() && fileInfo.Size() > 0 {
// 		klog.V(1).Infof("%sis a directory not a file\n", filePath)
// 		klog.V(6).Infof("rest.DoMultiPartFile LEAVE\n")
// 		return ErrInvalidInput
// 	}

// 	baseName := filepath.Base(strings.TrimSpace(filePath))
// 	klog.V(4).Infof("Filename: %s\n", baseName)

// 	r, w := io.Pipe()
// 	m := multipart.NewWriter(w)

// 	go func() {
// 		defer w.Close()
// 		defer m.Close()
// 		part, err := m.CreateFormFile("name", baseName)
// 		if err != nil {
// 			return
// 		}
// 		file, err := os.Open(filePath)
// 		if err != nil {
// 			return
// 		}
// 		defer file.Close()
// 		if _, err = io.Copy(part, file); err != nil {
// 			return
// 		}
// 	}()

// 	URI := version.GetAsyncAPI(version.ProcessAudioURI, baseName)
// 	klog.V(6).Infof("URI: %s\n", URI)

// 	req, err := http.NewRequestWithContext(ctx, "POST", URI, r)
// 	if err != nil {
// 		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
// 		klog.V(6).Infof("rest.DoMultiPartFile LEAVE\n")
// 		return err
// 	}

// 	if headers, ok := ctx.Value(interfaces.HeadersContext{}).(http.Header); ok {
// 		for k, v := range headers {
// 			for _, v := range v {
//				klog.V(3).Infof("DoMultiPartFile() Custom Header: %s = %s\n", k, v)
// 				req.Header.Add(k, v)
// 			}
// 		}
// 	}

// 	req.Header.Set("Accept", "application/json")
// if c.auth != nil && c.auth.NebulaToken != "" {
// 	req.Header.Set("ApiKey", c.auth.NebulaToken)
// } else if c.auth != nil && c.auth.AccessToken != "" {
// 	req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
// }

// 	contentType := m.FormDataContentType()
// 	req.Header.Set("Content-Type", contentType)
// 	klog.V(4).Infof("Content-Type; %s\n", contentType)

// 	err = c.Client.Do(ctx, req, func(res *http.Response) error {
// 		switch res.StatusCode {
// 		case http.StatusOK:
// 		case http.StatusCreated:
// 		case http.StatusNoContent:
// 		case http.StatusBadRequest:
// 			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
// 			detail, err := io.ReadAll(res.Body)
// 			if err != nil {
// 				klog.V(4).Infof("io.ReadAll failed. Err: %e\n", err)
// 				klog.V(6).Infof("rest.DoFile LEAVE\n")
// 				return err
// 			}
// 			klog.V(6).Infof("rest.DoFile LEAVE\n")
// 			return fmt.V(1).Infof("%s: %s", res.Status, bytes.TrimSpace(detail))
// 		default:
// 			return &interfaces.StatusError{res}
// 		}

// 		if resBody == nil {
// 			klog.V(4).Infof("resBody == nil\n")
// 			klog.V(6).Infof("rest.DoFile LEAVE\n")
// 			return nil
// 		}

// 		switch b := resBody.(type) {
// 		case *interfaces.RawResponse:
// 			klog.V(4).Infof("RawResponse\n")
// 			klog.V(6).Infof("rest.DoFile LEAVE\n")
// 			return res.Write(b)
// 		case io.Writer:
// 			klog.V(4).Infof("io.Writer\n")
// 			klog.V(6).Infof("rest.DoFile LEAVE\n")
// 			_, err := io.Copy(b, res.Body)
// 			return err
// 		default:
// 			klog.V(4).Infof("json.NewDecoder\n")
// 			d := json.NewDecoder(res.Body)
// 			klog.V(6).Infof("rest.DoFile LEAVE\n")
// 			return d.Decode(resBody)
// 		}
// 	})

// 	if err != nil {
// 		klog.V(1).Infof("err = c.Client.Do failed. Err: %v\n", err)
// 		return err
// 	}

// 	klog.V(3).Infof("rest.DoMultiPartFile Succeeded\n"))
// 	klog.V(6).Infof("rest.DoMultiPartFile LEAVE\n")
// 	return nil
// }

// DoAppendText appends Text to a given conversation ID
func (c *Client) DoAppendText(ctx context.Context, conversationId string, text asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	if len(conversationId) == 0 {
		klog.V(1).Infof("ConversationID is not valid\n")
		return ErrInvalidInput
	}

	return c.doCommonText(ctx, conversationId, text, resBody)
}

// DoAppendText initializes Text for a given conversation ID
func (c *Client) DoText(ctx context.Context, text asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	return c.doCommonText(ctx, "", text, resBody)
}

func (c *Client) doCommonText(ctx context.Context, conversationId string, text asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	klog.V(6).Infof("rest.doCommonText ENTER\n")

	// validate input
	v := validator.New()
	err := v.Struct(text)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("NewWithCreds validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("rest.doCommonText LEAVE\n")
		return err
	}

	verb := "POST"
	URI := version.GetAsyncAPI(version.ProcessTextURI)
	if len(conversationId) > 0 {
		verb = "PUT"
		URI = version.GetAsyncAPI(version.ProcessAppendTextURI, conversationId)
	}
	klog.V(6).Infof("verb: %s\n", verb)
	klog.V(6).Infof("URI: %s\n", URI)

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(text)
	if err != nil {
		klog.V(1).Infof("json.NewEncoder().Encode() failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonText LEAVE\n")
		return err
	}

	req, err := http.NewRequestWithContext(ctx, verb, URI, &buf)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonText LEAVE\n")
		return err
	}

	if headers, ok := ctx.Value(interfaces.HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				klog.V(3).Infof("doCommonText() Custom Header: %s = %s\n", k, v)
				req.Header.Add(k, v)
			}
		}
	}

	switch req.Method {
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		klog.V(3).Infof("Content-Type = application/json\n")
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	if c.auth != nil && c.auth.NebulaToken != "" {
		req.Header.Set("ApiKey", c.auth.NebulaToken)
	} else if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err = c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(1).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, errBody := io.ReadAll(res.Body)
			if err != nil {
				klog.V(1).Infof("io.ReadAll failed. Err: %e\n", errBody)
				klog.V(6).Infof("rest.doCommonText LEAVE\n")
				return &interfaces.StatusError{res}
			}
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(1).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(3).Infof("RawResponse\n")
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(3).Infof("io.Writer\n")
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(3).Infof("json.NewDecoder\n")
			d := json.NewDecoder(res.Body)
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			return d.Decode(resBody)
		}
	})

	if err != nil {
		klog.V(1).Infof("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonText LEAVE\n")
		return err
	}

	klog.V(3).Infof("rest.doCommonText Succeeded\n")
	klog.V(6).Infof("rest.doCommonText LEAVE\n")
	return nil
}

// DoFile posts a file capturing a conversation to a given REST endpoint
func (c *Client) DoFile(ctx context.Context, filePath string, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	// file?
	fileInfo, err := os.Stat(filePath)
	if err != nil || errors.Is(err, os.ErrNotExist) {
		klog.V(1).Infof("File %s does not exist. Err : %v\n", filePath, err)
		return err
	}

	if fileInfo.IsDir() || fileInfo.Size() == 0 {
		klog.V(1).Infof("%s is a directory not a file\n", filePath)
		return ErrInvalidInput
	}

	baseName := filepath.Base(strings.TrimSpace(filePath))
	klog.V(4).Infof("filePath: %s\n", filePath)
	klog.V(4).Infof("baseName: %s\n", baseName)

	// file
	pos := strings.LastIndex(filePath, ".")
	if pos == -1 {
		err := ErrInvalidURIExtension
		klog.V(1).Infof("uri is invalid. Err: %v\n", err)
		return err
	}

	extension := filePath[pos+1:]
	klog.V(3).Infof("extension: %s\n", extension)

	// is audio?
	switch extension {
	case common.AudioTypeMP3:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioFile(ctx, filePath, ufRequest, resBody)
	case common.AudioTypeMpeg:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioFile(ctx, filePath, ufRequest, resBody)
	case common.AudioTypeWav:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioFile(ctx, filePath, ufRequest, resBody)
	}

	// assume video
	klog.V(3).Infof("Defaulting IsVideo = TRUE\n")
	return c.doVideoFile(ctx, filePath, ufRequest, resBody)
}

func (c *Client) doAudioFile(ctx context.Context, filePath string, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	return c.doCommonFile(ctx, version.ProcessAudioURI, filePath, ufRequest, resBody)
}

func (c *Client) doVideoFile(ctx context.Context, filePath string, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	return c.doCommonFile(ctx, version.ProcessVideoURI, filePath, ufRequest, resBody)
}

func (c *Client) doCommonFile(ctx context.Context, apiURI, filePath string, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	klog.V(6).Infof("rest.doCommonFile ENTER\n")
	klog.V(4).Infof("rest.doCommonFile apiURI: %s\n", apiURI)

	// checks
	fileInfo, err := os.Stat(filePath)
	if err != nil || errors.Is(err, os.ErrNotExist) {
		klog.V(1).Infof("File %s does not exist. Err : %v\n", filePath, err)
		klog.V(6).Infof("rest.doCommonFile LEAVE\n")
		return err
	}

	if fileInfo.IsDir() || fileInfo.Size() == 0 {
		klog.V(1).Infof("%s is a directory not a file\n", filePath)
		klog.V(6).Infof("rest.doCommonFile LEAVE\n")
		return ErrInvalidInput
	}

	baseName := filepath.Base(strings.TrimSpace(filePath))
	klog.V(4).Infof("filePath: %s\n", filePath)
	klog.V(4).Infof("baseName: %s\n", baseName)

	file, err := os.Open(filePath)
	if err != nil {
		klog.V(1).Infof("os.Open failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonFile LEAVE\n")
		return err
	}
	defer file.Close()

	// start: until multipart post is supported, options must be used as a query string
	params := make(map[string][]string, 0)

	if len(ufRequest.Name) > 0 {
		params["name"] = []string{ufRequest.Name}
	}
	if ufRequest.ConfidenceThreshold > 0 {
		params["confidenceThreshold"] = []string{fmt.Sprintf("%f", ufRequest.ConfidenceThreshold)}
	}
	// TODO: channelMetadata... need to see what that looks like as a query string param
	if ufRequest.DetectPhrases {
		params["detectPhrases"] = []string{fmt.Sprintf("%t", ufRequest.DetectPhrases)}
	}
	if ufRequest.DetectEntities {
		params["detectEntities"] = []string{fmt.Sprintf("%t", ufRequest.DetectEntities)}
	}
	if len(ufRequest.LanguageCode) > 0 {
		params["languageCode"] = []string{ufRequest.LanguageCode}
	}
	if len(ufRequest.CustomVocabulary) > 0 {
		params["customVocabulary"] = ufRequest.CustomVocabulary
	}
	if ufRequest.Sentiment {
		params["mode"] = []string{ufRequest.Mode}
	}
	if ufRequest.ParentRefs {
		params["parentRefs"] = []string{fmt.Sprintf("%t", ufRequest.ParentRefs)}
	}
	if ufRequest.Sentiment {
		params["sentiment"] = []string{fmt.Sprintf("%t", ufRequest.Sentiment)}
	}
	if ufRequest.Sentiment {
		params["enableSeparateRecognitionPerChannel"] = []string{fmt.Sprintf("%t", ufRequest.EnableSeparateRecognitionPerChannel)}
	}
	if ufRequest.Sentiment {
		params["enableSpeakerDiarization"] = []string{fmt.Sprintf("%t", ufRequest.EnableSpeakerDiarization)}
	}
	if ufRequest.Sentiment {
		params["diarizationSpeakerCount"] = []string{fmt.Sprintf("%d", ufRequest.DiarizationSpeakerCount)}
	}
	if ufRequest.Sentiment {
		params["webhookUrl"] = []string{ufRequest.WebhookURL}
	}
	// end

	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(apiURI, baseName),
		c.getQueryParamFromContext(ctx, &params))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "POST", URI, file)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonFile LEAVE\n")
		return err
	}

	if headers, ok := ctx.Value(interfaces.HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				klog.V(3).Infof("doCommonFile() Custom Header: %s = %s\n", k, v)
				req.Header.Add(k, v)
			}
		}
	}

	req.Header.Set("Accept", "application/json")
	if c.auth != nil && c.auth.NebulaToken != "" {
		req.Header.Set("ApiKey", c.auth.NebulaToken)
	} else if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err = c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, errBody := io.ReadAll(res.Body)
			if err != nil {
				klog.V(4).Infof("io.ReadAll failed. Err: %e\n", errBody)
				klog.V(6).Infof("rest.doCommonFile LEAVE\n")
				return &interfaces.StatusError{res}
			}
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(1).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(3).Infof("RawResponse\n")
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(3).Infof("io.Writer\n")
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(3).Infof("json.NewDecoder\n")
			d := json.NewDecoder(res.Body)
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			return d.Decode(resBody)
		}
	})

	if err != nil {
		klog.V(1).Infof("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonFile LEAVE\n")
		return err
	}

	klog.V(3).Infof("rest.doCommonFile Succeeded\n")
	klog.V(6).Infof("rest.doCommonFile LEAVE\n")
	return nil
}

// IsUrl returns true if a string is of a URL format
func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// DoURL performs a REST call using a URL conversation source
func (c *Client) DoURL(ctx context.Context, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	// url
	u, err := url.Parse(ufRequest.URL)
	if err != nil {
		klog.V(1).Infof("uri is invalid. Err: %v\n", err)
		return err
	}

	pos := strings.LastIndex(u.Path, ".")
	if pos == -1 {
		err := ErrInvalidURIExtension
		klog.V(1).Infof("uri is invalid. Err: %v\n", err)
		return err
	}

	extension := u.Path[pos+1:]
	klog.V(3).Infof("extension: %s\n", extension)

	// is audio?
	switch extension {
	case common.AudioTypeMP3:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioURL(ctx, ufRequest, resBody)
	case common.AudioTypeMpeg:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioURL(ctx, ufRequest, resBody)
	case common.AudioTypeWav:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioURL(ctx, ufRequest, resBody)
	}

	// assume video
	klog.V(3).Infof("Default IsVideo = TRUE\n")
	return c.doVideoURL(ctx, ufRequest, resBody)
}

func (c *Client) doAudioURL(ctx context.Context, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	return c.doCommonURL(ctx, version.ProcessAudioURLURI, ufRequest, resBody)
}

func (c *Client) doVideoURL(ctx context.Context, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	return c.doCommonURL(ctx, version.ProcessVideoURLURI, ufRequest, resBody)
}

func (c *Client) doCommonURL(ctx context.Context, apiURI string, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	klog.V(6).Infof("rest.DoURL ENTER\n")
	klog.V(4).Infof("rest.doCommonURL apiURI: %s\n", apiURI)

	// checks
	validURL := IsUrl(ufRequest.URL)
	if !validURL {
		klog.V(1).Infof("Invalid URL: %s\n", ufRequest.URL)
		klog.V(6).Infof("rest.doCommonURL LEAVE\n")
		return ErrInvalidInput
	}

	baseName := filepath.Base(strings.TrimSpace(ufRequest.URL))
	klog.V(4).Infof("url: %s\n", ufRequest.URL)
	klog.V(4).Infof("baseName: %s\n", baseName)

	if len(ufRequest.Name) == 0 {
		ufRequest.Name = baseName
	}

	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(apiURI),
		c.getQueryParamFromContext(ctx, nil))
	klog.V(6).Infof("Calling %s\n", URI)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(ufRequest)
	if err != nil {
		klog.V(1).Infof("json.NewEncoder().Encode() failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonURL LEAVE\n")
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, &buf)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonURL LEAVE\n")
		return err
	}

	if headers, ok := ctx.Value(interfaces.HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				klog.V(3).Infof("doCommonURL() Custom Header: %s = %s\n", k, v)
				req.Header.Add(k, v)
			}
		}
	}

	switch req.Method {
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		klog.V(3).Infof("Content-Type = application/json\n")
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	if c.auth != nil && c.auth.NebulaToken != "" {
		req.Header.Set("ApiKey", c.auth.NebulaToken)
	} else if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err = c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, errBody := io.ReadAll(res.Body)
			if err != nil {
				klog.V(1).Infof("io.ReadAll failed. Err: %e\n", errBody)
				klog.V(6).Infof("rest.doCommonURL LEAVE\n")
				return &interfaces.StatusError{res}
			}
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(1).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(3).Infof("RawResponse\n")
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(3).Infof("io.Writer\n")
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(3).Infof("json.NewDecoder\n")
			d := json.NewDecoder(res.Body)
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			return d.Decode(resBody)
		}
	})

	if err != nil {
		klog.V(1).Infof("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.doCommonURL LEAVE\n")
		return err
	}

	klog.V(3).Infof("rest.doCommonURL Succeeded\n")
	klog.V(6).Infof("rest.doCommonURL LEAVE\n")
	return nil
}

// Do is a generic REST API call to the platform
func (c *Client) Do(ctx context.Context, req *http.Request, resBody interface{}) error {
	klog.V(6).Infof("rest.Do ENTER\n")

	if headers, ok := ctx.Value(interfaces.HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				klog.V(3).Infof("Do() Custom Header: %s = %s\n", k, v)
				req.Header.Add(k, v)
			}
		}
	}

	switch req.Method {
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		klog.V(3).Infof("Content-Type = application/json\n")
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	if c.auth != nil && c.auth.NebulaToken != "" {
		req.Header.Set("ApiKey", c.auth.NebulaToken)
	} else if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err := c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(1).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, errBody := io.ReadAll(res.Body)
			if errBody != nil {
				klog.V(1).Infof("io.ReadAll failed. Err: %e\n", errBody)
				klog.V(6).Infof("rest.DoFile LEAVE\n")
				return &interfaces.StatusError{res}
			}
			klog.V(6).Infof("rest.Do LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(1).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.Do LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(3).Infof("RawResponse\n")
			klog.V(6).Infof("rest.Do LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(3).Infof("io.Writer\n")
			klog.V(6).Infof("rest.Do LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(3).Infof("json.NewDecoder\n")
			d := json.NewDecoder(res.Body)
			klog.V(6).Infof("rest.Do LEAVE\n")
			return d.Decode(resBody)
		}
	})

	if err != nil {
		klog.V(1).Infof("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.Do LEAVE\n")
		return err
	}

	klog.V(3).Infof("rest.Do Succeeded\n")
	klog.V(6).Infof("rest.Do LEAVE\n")
	return nil
}

func (c *Client) getQueryParamFromContext(ctx context.Context, input *map[string][]string) string {
	if input == nil {
		tmp := make(map[string][]string, 0)
		input = &tmp
	}

	if parameters, ok := ctx.Value(interfaces.ParametersContext{}).(map[string][]string); ok {
		for k, vs := range parameters {
			klog.V(5).Infof("Key/Value: %s = %v\n", k, vs)
			(*input)[k] = vs
		}
	}

	// TODO: replace with https://github.com/google/go-querystring
	// API differs from how go-querystring works
	//
	//	go-querystring : []vals -> key=vals[0]&key=vals[1]
	//	symbl API: []vals -> key=["$vals[0]", "$vals[1]"]
	//
	// need to look into switching that behavior in order to use go-querystring lib
	if len(*input) > 1 {
		queryString := "&"
		for k, vs := range *input {
			if len(queryString) > 3 {
				queryString += "&"
			}
			if len(vs) == 1 {
				queryString += fmt.Sprintf("%s=%s", k, vs[0])
			} else {
				appended := false
				for _, v := range vs {
					if !appended {
						queryString += fmt.Sprintf("%s=[", k)
						appended = true
					} else {
						queryString += ","
					}
					queryString += fmt.Sprintf("%s", v)
				}
				if len(vs) > 0 {
					queryString += "]"
				}
			}
		}
		klog.V(5).Infof("Final Query String: %s\n", queryString)
		return queryString
	}

	klog.V(6).Infof("Final Query String is Empty\n")
	return ""
}
