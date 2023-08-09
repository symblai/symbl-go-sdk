// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package symbl

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
	rest "github.com/symblai/symbl-go-sdk/pkg/client/rest"
)

const (
	defaultAuthType    string = "application"
	defaultAuthTimeout int64  = 5

	defaultAttemptsToReauth   int   = 3
	defaultDelayBetweenReauth int64 = 2
)

// NewRestClient creates a new client on the Symbl.ai platform.
// The client authenticates with the server with APP_ID/APP_SECRET as defined in environment variables.
func NewRestClient(ctx context.Context) (*RestClient, error) {
	var appId string
	if v := os.Getenv("APP_ID"); v != "" {
		klog.V(4).Info("APP_ID found")
		appId = v
	} else {
		klog.Error("APP_ID not found")
		return nil, ErrInvalidInput
	}
	var appSecret string
	if v := os.Getenv("APP_SECRET"); v != "" {
		klog.V(4).Info("APP_SECRET found")
		appSecret = v
	} else {
		klog.Errorln("APP_SECRET not found")
		return nil, ErrInvalidInput
	}
	var symblEndpoint string
	if v := os.Getenv("SYMBL_ENDPOINT"); v != "" {
		klog.V(4).Info("SYMBL_ENDPOINT found")
		symblEndpoint = v
	} else {
		klog.V(3).Infof("SYMBL_ENDPOINT not found. Use default.")
	}

	c := interfaces.Credentials{
		AuthURI:   symblEndpoint,
		AppId:     appId,
		AppSecret: appSecret,
	}
	return NewRestClientWithCreds(ctx, c)
}

// NewRestClientWithCreds creates a new client on the Symbl.ai platform.
// The client authenticates with the server using APP_ID/APP_SECRET provided in Credentials struct
func NewRestClientWithCreds(ctx context.Context, creds interfaces.Credentials) (*RestClient, error) {
	klog.V(6).Infof("NewRestClientWithCreds ENTER\n")

	if len(creds.AuthURI) > 0 {
		klog.V(3).Infof("[OVERRIDE] AuthURI: %s\n", creds.AuthURI)
	} else {
		creds.AuthURI = defaultAuthURI
	}

	// checks
	if ctx == nil {
		klog.V(3).Infof("Empty Context... Creating new one!\n")
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(creds)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("NewRestClientWithCreds validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("NewRestClientWithCreds LEAVE\n")
		return nil, err
	}

	if len(creds.Type) == 0 {
		creds.Type = defaultAuthType
	}

	// let's auth
	jsonStr, err := json.Marshal(creds)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("NewRestClientWithCreds LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", creds.AuthURI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("NewRestClientWithCreds LEAVE\n")
		return nil, err
	}

	// restore application options to HTTP header
	if headers, ok := ctx.Value(interfaces.HeadersContext{}).(http.Header); ok {
		for k, vs := range headers {
			for _, v := range vs {
				klog.V(4).Infof("NewRestClientWithCreds() RESTORE Header: %s = %s\n", k, v)
				req.Header.Add(k, v)
			}
		}
	}

	// do it!
	var resp interfaces.AuthResp

	restClient := rest.New()
	err = restClient.Do(ctx, req, &resp)
	if err != nil {
		klog.V(1).Infof("restClient.Do failed. Err: %v\n", err)
		return nil, err
	}

	if resp.AccessToken == "" {
		klog.V(1).Infof("Symbl auth token is empty\n")
		klog.V(6).Infof("NewRestClientWithCreds LEAVE\n")
		return nil, ErrAuthFailure
	}

	restClient.SetAuthorization(&rest.AccessToken{
		AccessToken: resp.AccessToken,
		ExpiresOn:   time.Now().Add(time.Second * time.Duration(resp.ExpiresIn)),
	})

	c := &RestClient{
		Client: restClient,
		creds:  &creds,
		auth:   &resp,
	}

	klog.V(3).Infof("NewRestClientWithCreds Succeeded\n")
	klog.V(6).Infof("NewRestClientWithCreds LEAVE\n")
	return c, nil
}

// NewRestClientWithToken creates a new client on the Symbl.ai platform.
// The client authenticates reusing an already valid Symbl Platform auth token
func NewRestClientWithToken(ctx context.Context, accessToken string) (*RestClient, error) {
	klog.V(6).Infof("NewRestClientWithToken ENTER\n")

	creds := interfaces.Credentials{
		Type: defaultAuthType,
	}
	resp := interfaces.AuthResp{
		AccessToken: accessToken,
	}

	if len(creds.AuthURI) > 0 {
		klog.V(3).Infof("[OVERRIDE] AuthURI: %s\n", creds.AuthURI)
	} else {
		creds.AuthURI = defaultAuthURI
	}

	// checks
	if ctx == nil {
		klog.V(3).Infof("Empty Context... Creating new one!\n")
		ctx = context.Background()
	}

	// validate input
	if resp.AccessToken == "" {
		klog.V(1).Infof("Symbl auth token is empty\n")
		klog.V(6).Infof("NewRestClientWithToken LEAVE\n")
		return nil, ErrInvalidInput
	}

	restClient := rest.New()
	restClient.SetAuthorization(&rest.AccessToken{
		AccessToken: resp.AccessToken,
		ExpiresOn:   time.Now().Add(time.Hour * 24),
	})

	c := &RestClient{
		Client: restClient,
		creds:  &creds,
		auth:   &resp,
	}

	klog.V(3).Infof("NewRestClientWithToken Succeeded\n")
	klog.V(6).Infof("NewRestClientWithToken LEAVE\n")
	return c, nil
}

// DoTextWithOptions wrapper function for REST Client. Please see pkg/client/rest
func (c *RestClient) DoTextWithOptions(ctx context.Context, options asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	return c.Client.DoText(ctx, options, resBody)
}

// DoAppendTextWithOptions wrapper function for REST Client. Please see pkg/client/rest
func (c *RestClient) DoAppendTextWithOptions(ctx context.Context, conversationId string, options asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	return c.Client.DoAppendText(ctx, conversationId, options, resBody)
}

// DoFileWithOptions wrapper function for REST Client. Please see pkg/client/rest
func (c *RestClient) DoFileWithOptions(ctx context.Context, filePath string, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	return c.Client.DoFile(ctx, filePath, ufRequest, resBody)
}

// DoURLWithOptions wrapper function for REST Client. Please see pkg/client/rest
func (c *RestClient) DoURLWithOptions(ctx context.Context, ufRequest asyncinterfaces.AsyncURLFileRequest, resBody interface{}) error {
	return c.Client.DoURL(ctx, ufRequest, resBody)
}

// DoFile wrapper function for REST Client. Please see pkg/client/rest
func (c *RestClient) DoFile(ctx context.Context, filePath string, resBody interface{}) error {
	ufRequest := asyncinterfaces.AsyncURLFileRequest{}
	return c.DoFileWithOptions(ctx, filePath, ufRequest, resBody)
}

// DoURL wrapper function for REST Client. Please see pkg/client/rest
func (c *RestClient) DoURL(ctx context.Context, url string, resBody interface{}) error {
	ufRequest := asyncinterfaces.AsyncURLFileRequest{
		URL: url,
	}

	return c.DoURLWithOptions(ctx, ufRequest, resBody)
}
