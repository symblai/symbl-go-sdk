// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

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

	asyncinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	common "github.com/dvonthenen/symbl-go-sdk/pkg/api/common"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	simple "github.com/dvonthenen/symbl-go-sdk/pkg/client/simple"
)

// Client which extends basic client to support REST
type Client struct {
	*simple.Client

	auth *AccessToken
}

func New() *Client {
	c := Client{
		Client: simple.New(),
	}
	return &c
}

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
// 	if c.auth != nil && c.auth.AccessToken != "" {
// 		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
// 	}

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

func (c *Client) DoAppendText(ctx context.Context, conversationId string, text asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	if len(conversationId) == 0 {
		klog.V(1).Infof("ConversationID is not valid\n")
		return ErrInvalidInput
	}

	return c.doCommonText(ctx, conversationId, text, resBody)
}

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
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err = c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, err := io.ReadAll(res.Body)
			if err != nil {
				klog.V(4).Infof("io.ReadAll failed. Err: %e\n", err)
				klog.V(6).Infof("rest.doCommonText LEAVE\n")
				return err
			}
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(4).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(4).Infof("RawResponse\n")
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(4).Infof("io.Writer\n")
			klog.V(6).Infof("rest.doCommonText LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(4).Infof("json.NewDecoder\n")
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

func (c *Client) DoFile(ctx context.Context, filePath string, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
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
		return c.doAudioFile(ctx, filePath, options, resBody)
	case common.AudioTypeMpeg:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioFile(ctx, filePath, options, resBody)
	case common.AudioTypeWav:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioFile(ctx, filePath, options, resBody)
	}

	// assume video
	klog.V(3).Infof("Defaulting IsVideo = TRUE\n")
	return c.doVideoFile(ctx, filePath, options, resBody)
}

func (c *Client) doAudioFile(ctx context.Context, filePath string, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	return c.doCommonFile(ctx, version.ProcessAudioURI, filePath, options, resBody)
}

func (c *Client) doVideoFile(ctx context.Context, filePath string, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	return c.doCommonFile(ctx, version.ProcessVideoURI, filePath, options, resBody)
}

func (c *Client) doCommonFile(ctx context.Context, apiURI, filePath string, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	klog.V(6).Infof("rest.doCommonFile ENTER\n")

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
	params := make([]string, 0)
	params = append(params, "?")

	if len(options.Name) > 0 {
		params = append(params, fmt.Sprintf("name=%s", options.Name))
	}
	if options.ConfidenceThreshold > 0 {
		params = append(params, fmt.Sprintf("confidenceThreshold=%f", options.ConfidenceThreshold))
	}
	if options.DetectPhrases {
		params = append(params, fmt.Sprintf("detectPhrases=%t", options.DetectPhrases))
	}
	if options.DetectEntities {
		params = append(params, fmt.Sprintf("detectEntities=%t", options.DetectEntities))
	}
	if len(options.LanguageCode) > 0 {
		params = append(params, fmt.Sprintf("languageCode=%s", options.LanguageCode))
	}
	for _, vocab := range options.CustomVocabulary {
		params = append(params, fmt.Sprintf("customVocabulary=%s", vocab))
	}
	if options.ParentRefs {
		params = append(params, fmt.Sprintf("parentRefs=%t", options.ParentRefs))
	}
	if options.Sentiment {
		params = append(params, fmt.Sprintf("sentiment=%t", options.Sentiment))
	}
	// end

	URI := version.GetAsyncAPI(apiURI, baseName)
	if len(params) > 1 {
		URI = version.GetAsyncAPI(apiURI, baseName, params)
	}
	klog.V(6).Infof("URI: %s\n", URI)

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
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err = c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, err := io.ReadAll(res.Body)
			if err != nil {
				klog.V(4).Infof("io.ReadAll failed. Err: %e\n", err)
				klog.V(6).Infof("rest.doCommonFile LEAVE\n")
				return err
			}
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(4).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(4).Infof("RawResponse\n")
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(4).Infof("io.Writer\n")
			klog.V(6).Infof("rest.doCommonFile LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(4).Infof("json.NewDecoder\n")
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

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (c *Client) DoURL(ctx context.Context, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	// url
	u, err := url.Parse(options.URL)
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
		return c.doAudioURL(ctx, options, resBody)
	case common.AudioTypeMpeg:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioURL(ctx, options, resBody)
	case common.AudioTypeWav:
		klog.V(3).Infof("IsAudio = TRUE\n")
		return c.doAudioURL(ctx, options, resBody)
	}

	// assume video
	klog.V(3).Infof("Default IsVideo = TRUE\n")
	return c.doVideoURL(ctx, options, resBody)
}

func (c *Client) doAudioURL(ctx context.Context, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	return c.doCommonURL(ctx, version.ProcessAudioURLURI, options, resBody)
}

func (c *Client) doVideoURL(ctx context.Context, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	return c.doCommonURL(ctx, version.ProcessVideoURLURI, options, resBody)
}

func (c *Client) doCommonURL(ctx context.Context, apiURI string, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	klog.V(6).Infof("rest.DoURL ENTER\n")

	// checks
	validURL := IsUrl(options.URL)
	if !validURL {
		klog.V(1).Infof("Invalid URL: %s\n", options.URL)
		klog.V(6).Infof("rest.doCommonURL LEAVE\n")
		return ErrInvalidInput
	}

	baseName := filepath.Base(strings.TrimSpace(options.URL))
	klog.V(4).Infof("url: %s\n", options.URL)
	klog.V(4).Infof("baseName: %s\n", baseName)

	if len(options.Name) == 0 {
		options.Name = baseName
	}

	URI := version.GetAsyncAPI(apiURI)
	klog.V(6).Infof("URI: %s\n", URI)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(options)
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
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err = c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, err := io.ReadAll(res.Body)
			if err != nil {
				klog.V(4).Infof("io.ReadAll failed. Err: %e\n", err)
				klog.V(6).Infof("rest.doCommonURL LEAVE\n")
				return err
			}
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(4).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(4).Infof("RawResponse\n")
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(4).Infof("io.Writer\n")
			klog.V(6).Infof("rest.doCommonURL LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(4).Infof("json.NewDecoder\n")
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
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	err := c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			detail, err := io.ReadAll(res.Body)
			if err != nil {
				klog.V(4).Infof("io.ReadAll failed. Err: %e\n", err)
				klog.V(6).Infof("rest.DoFile LEAVE\n")
				return err
			}
			klog.V(6).Infof("rest.Do LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &interfaces.StatusError{res}
		}

		if resBody == nil {
			klog.V(4).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.Do LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *interfaces.RawResponse:
			klog.V(4).Infof("RawResponse\n")
			klog.V(6).Infof("rest.Do LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(4).Infof("io.Writer\n")
			klog.V(6).Infof("rest.Do LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(4).Infof("json.NewDecoder\n")
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
