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

	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
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

// WithHeader returns a new Context populated with the provided headers map
func (c *Client) WithHeader(
	ctx context.Context,
	headers http.Header) context.Context {

	return context.WithValue(ctx, HeadersContext{}, headers)
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

// 	baseName := filepath.Base(filePath)
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

// 	if headers, ok := ctx.Value(HeadersContext{}).(http.Header); ok {
// 		for k, v := range headers {
// 			for _, v := range v {
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
// 			return &StatusError{res}
// 		}

// 		if resBody == nil {
// 			klog.V(4).Infof("resBody == nil\n")
// 			klog.V(6).Infof("rest.DoFile LEAVE\n")
// 			return nil
// 		}

// 		switch b := resBody.(type) {
// 		case *RawResponse:
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

func (c *Client) DoFile(ctx context.Context, filePath string, options interfaces.AsyncOptions, resBody interface{}) error {
	klog.V(6).Infof("rest.DoFile ENTER\n")

	// checks
	fileInfo, err := os.Stat(filePath)
	if err != nil || errors.Is(err, os.ErrNotExist) {
		klog.V(1).Infof("File %s does not exist. Err : %v\n", filePath, err)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return err
	}

	if fileInfo.IsDir() && fileInfo.Size() > 0 {
		klog.V(1).Infof("%s is a directory not a file\n", filePath)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return ErrInvalidInput
	}

	baseName := filepath.Base(filePath)
	klog.V(4).Infof("filePath: %s\n", filePath)
	klog.V(4).Infof("baseName: %s\n", baseName)

	file, err := os.Open(filePath)
	if err != nil {
		klog.V(1).Infof("os.Open failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
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

	URI := version.GetAsyncAPI(version.ProcessAudioURI, baseName)
	if len(params) > 1 {
		URI = version.GetAsyncAPI(version.ProcessAudioURI, baseName, params)
	}
	klog.V(6).Infof("URI: %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "POST", URI, file)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return err
	}

	if headers, ok := ctx.Value(HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				req.Header.Add(k, v)
			}
		}
	}

	req.Header.Set("Accept", "application/json")
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	if headers, ok := ctx.Value(HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				req.Header.Add(k, v)
			}
		}
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
				klog.V(6).Infof("rest.DoFile LEAVE\n")
				return err
			}
			klog.V(6).Infof("rest.DoFile LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &StatusError{res}
		}

		if resBody == nil {
			klog.V(4).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.DoFile LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *RawResponse:
			klog.V(4).Infof("RawResponse\n")
			klog.V(6).Infof("rest.DoFile LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(4).Infof("io.Writer\n")
			klog.V(6).Infof("rest.DoFile LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(4).Infof("json.NewDecoder\n")
			d := json.NewDecoder(res.Body)
			klog.V(6).Infof("rest.DoFile LEAVE\n")
			return d.Decode(resBody)
		}
	})

	if err != nil {
		klog.V(1).Infof("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return err
	}

	klog.V(3).Infof("rest.DoFile Succeeded\n")
	klog.V(6).Infof("rest.DoFile LEAVE\n")
	return nil
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (c *Client) DoURL(ctx context.Context, options interfaces.AsyncOptions, resBody interface{}) error {
	klog.V(6).Infof("rest.DoURL ENTER\n")

	// checks
	validURL := IsUrl(options.URL)
	if !validURL {
		klog.V(1).Infof("Invalid URL: %s\n", options.URL)
		klog.V(6).Infof("rest.DoURL LEAVE\n")
		return ErrInvalidInput
	}

	baseName := filepath.Base(options.URL)
	klog.V(4).Infof("url: %s\n", options.URL)
	klog.V(4).Infof("baseName: %s\n", baseName)

	if len(options.Name) == 0 {
		options.Name = baseName
	}

	URI := version.GetAsyncAPI(version.ProcessURLURI)
	klog.V(6).Infof("URI: %s\n", URI)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(options)
	if err != nil {
		klog.V(1).Infof("json.NewEncoder().Encode() failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoURL LEAVE\n")
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, &buf)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoURL LEAVE\n")
		return err
	}

	if headers, ok := ctx.Value(HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				req.Header.Add(k, v)
			}
		}
	}

	switch req.Method {
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	if headers, ok := ctx.Value(HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				req.Header.Add(k, v)
			}
		}
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
				klog.V(6).Infof("rest.DoURL LEAVE\n")
				return err
			}
			klog.V(6).Infof("rest.DoURL LEAVE\n")
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &StatusError{res}
		}

		if resBody == nil {
			klog.V(4).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.DoURL LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *RawResponse:
			klog.V(4).Infof("RawResponse\n")
			klog.V(6).Infof("rest.DoURL LEAVE\n")
			return res.Write(b)
		case io.Writer:
			klog.V(4).Infof("io.Writer\n")
			klog.V(6).Infof("rest.DoURL LEAVE\n")
			_, err := io.Copy(b, res.Body)
			return err
		default:
			klog.V(4).Infof("json.NewDecoder\n")
			d := json.NewDecoder(res.Body)
			klog.V(6).Infof("rest.DoURL LEAVE\n")
			return d.Decode(resBody)
		}
	})

	if err != nil {
		klog.V(1).Infof("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoURL LEAVE\n")
		return err
	}

	klog.V(3).Infof("rest.DoURL Succeeded\n")
	klog.V(6).Infof("rest.DoURL LEAVE\n")
	return nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, resBody interface{}) error {
	klog.V(6).Infof("rest.Do ENTER\n")

	if headers, ok := ctx.Value(HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				req.Header.Add(k, v)
			}
		}
	}

	switch req.Method {
	case http.MethodPost, http.MethodPatch, http.MethodPut:
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
			return &StatusError{res}
		}

		if resBody == nil {
			klog.V(4).Infof("resBody == nil\n")
			klog.V(6).Infof("rest.Do LEAVE\n")
			return nil
		}

		switch b := resBody.(type) {
		case *RawResponse:
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
