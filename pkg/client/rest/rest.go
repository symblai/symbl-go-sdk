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
	"os"
	"path/filepath"
	"time"

	klog "k8s.io/klog/v2"

	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	simple "github.com/dvonthenen/symbl-go-sdk/pkg/client/simple"
)

var (
	// ErrInvalidInput required input was not found
	ErrInvalidInput = errors.New("required input was not found")
)

// Client which extends basic client to support REST
type Client struct {
	*simple.Client

	auth *AccessToken
}

// AccessToken represents a Symbl platform bearer access token with expiry information.
type AccessToken struct {
	AccessToken string
	ExpiresOn   time.Time
}

// RawResponse may be used with the Do method as the resBody argument in order
// to capture the raw response data.
type RawResponse struct {
	bytes.Buffer
}

type HeadersContext struct{}

type StatusError struct {
	Resp *http.Response
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Resp.Request.Method, e.Resp.Request.URL, e.Resp.Status)
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
// 		klog.Errorf("File %s does not exist. Err : %v\n", filePath, err)
// 		klog.V(6).Infof("rest.DoMultiPartFile LEAVE\n")
// 		return err
// 	}

// 	if fileInfo.IsDir() && fileInfo.Size() > 0 {
// 		klog.Errorf("%sis a directory not a file\n", filePath)
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
// 		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
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
// 	// TODO: verify this is correct
// 	if c.auth != nil && c.auth.AccessToken != "" {
// 		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
// 	}

// 	contentType := m.FormDataContentType()
// 	req.Header.Set("Content-Type", contentType)
// 	klog.V(4).Infof("Content-Type; %s\n", contentType)

// 	klog.V(6).Infof("------------------------\n")
// 	klog.V(6).Infof("req:\n%v\n", req)
// 	klog.V(6).Infof("------------------------\n")

// 	err = c.Client.Do(ctx, req, func(res *http.Response) error {
// 		switch res.StatusCode {
// 		case http.StatusOK:
// 		case http.StatusCreated:
// 		case http.StatusNoContent:
// 		case http.StatusBadRequest:
// 			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
// 			// TODO: structured error types
// 			detail, err := io.ReadAll(res.Body)
// 			if err != nil {
// 				klog.V(4).Infof("io.ReadAll failed. Err: %e\n", err)
// 				klog.V(6).Infof("rest.DoFile LEAVE\n")
// 				return err
// 			}
// 			klog.V(6).Infof("rest.DoFile LEAVE\n")
// 			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
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
// 		klog.Errorf("err = c.Client.Do failed. Err: %v\n", err)
// 		return err
// 	}

// 	klog.V(6).Infof("------------------------\n")
// 	klog.V(6).Infof("resBody:\n%v\n", resBody)
// 	klog.V(6).Infof("------------------------\n")

// 	klog.V(4).Infof("rest.DoMultiPartFile Succeeded\n")
// 	klog.V(6).Infof("rest.DoMultiPartFile LEAVE\n")
// 	return nil
// }

func (c *Client) DoFile(ctx context.Context, filePath string, resBody interface{}) error {
	klog.V(6).Infof("rest.DoFile ENTER\n")

	// checks
	fileInfo, err := os.Stat(filePath)
	if err != nil || errors.Is(err, os.ErrNotExist) {
		klog.Errorf("File %s does not exist. Err : %v\n", filePath, err)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return err
	}

	if fileInfo.IsDir() && fileInfo.Size() > 0 {
		klog.Errorf("%s is a directory not a file\n", filePath)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return ErrInvalidInput
	}

	baseName := filepath.Base(filePath)
	klog.V(4).Infof("filePath: %s\n", filePath)
	klog.V(4).Infof("baseName: %s\n", baseName)

	file, err := os.Open(filePath)
	if err != nil {
		klog.Errorf("os.Open failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return err
	}
	defer file.Close()

	URI := version.GetAsyncAPI(version.ProcessAudioURI, baseName)
	klog.V(6).Infof("URI: %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "POST", URI, file)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
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
	// TODO: verify this is correct
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

	klog.V(6).Infof("------------------------\n")
	klog.V(6).Infof("req:\n%v\n", req)
	klog.V(6).Infof("------------------------\n")

	err = c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			// TODO: structured error types
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
		klog.Errorf("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.DoFile LEAVE\n")
		return err
	}

	klog.V(6).Infof("------------------------\n")
	klog.V(6).Infof("resBody:\n%v\n", resBody)
	klog.V(6).Infof("------------------------\n")

	klog.V(2).Infof("rest.DoFile Succeeded\n")
	klog.V(6).Infof("rest.DoFile LEAVE\n")
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
	// TODO: verify this is correct
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.AccessToken)
	}

	klog.V(6).Infof("------------------------\n")
	klog.V(6).Infof("req:\n%v\n", req)
	klog.V(6).Infof("------------------------\n")

	err := c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			klog.V(4).Infof("HTTP Error Code: %d\n", res.StatusCode)
			// TODO: structured error types
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
		klog.Errorf("err = c.Client.Do failed. Err: %v\n", err)
		klog.V(6).Infof("rest.Do LEAVE\n")
		return err
	}

	klog.V(6).Infof("------------------------\n")
	klog.V(6).Infof("resBody:\n%v\n", resBody)
	klog.V(6).Infof("------------------------\n")

	klog.V(2).Infof("rest.Do Succeeded\n")
	klog.V(6).Infof("rest.Do LEAVE\n")
	return nil
}
