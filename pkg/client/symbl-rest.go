// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

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

	asyncinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
)

const (
	defaultAuthType    string = "application"
	defaultAuthTimeout int64  = 5

	defaultAttemptsToReauth   int   = 3
	defaultDelayBetweenReauth int64 = 2
)

// NewRestClient creates a new client on the Symbl.ai platform. The client authenticates with the
// server with APP_ID/APP_SECRET.
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

// NewRestClientWithCreds creates a new client on the Symbl.ai platform. The client authenticates with the
// server with APP_ID/APP_SECRET.
func NewRestClientWithCreds(ctx context.Context, creds interfaces.Credentials) (*RestClient, error) {
	klog.V(6).Infof("NewWithCreds ENTER\n")

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
			klog.V(1).Infof("NewWithCreds validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("NewWithCreds LEAVE\n")
		return nil, err
	}

	if len(creds.Type) == 0 {
		creds.Type = defaultAuthType
	}

	// let's auth
	jsonStr, err := json.Marshal(creds)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("NewWithCreds LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", creds.AuthURI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("NewWithCreds LEAVE\n")
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
		klog.V(6).Infof("NewWithCreds LEAVE\n")
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

	klog.V(3).Infof("NewWithCreds Succeeded\n")
	klog.V(6).Infof("NewWithCreds LEAVE\n")
	return c, nil
}

func (c *RestClient) DoTextWithOptions(ctx context.Context, options asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	return c.Client.DoText(ctx, options, resBody)
}

func (c *RestClient) DoAppendTextWithOptions(ctx context.Context, conversationId string, options asyncinterfaces.AsyncTextRequest, resBody interface{}) error {
	return c.Client.DoAppendText(ctx, conversationId, options, resBody)
}

func (c *RestClient) DoFileWithOptions(ctx context.Context, filePath string, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	return c.Client.DoFile(ctx, filePath, options, resBody)
}

func (c *RestClient) DoURLWithOptions(ctx context.Context, options asyncinterfaces.AsyncOptions, resBody interface{}) error {
	return c.Client.DoURL(ctx, options, resBody)
}

func (c *RestClient) DoFile(ctx context.Context, filePath string, resBody interface{}) error {
	options := asyncinterfaces.AsyncOptions{}
	return c.DoFileWithOptions(ctx, filePath, options, resBody)
}

func (c *RestClient) DoURL(ctx context.Context, url string, resBody interface{}) error {
	options := asyncinterfaces.AsyncOptions{
		URL: url,
	}

	return c.DoURLWithOptions(ctx, options, resBody)
}

func (c *RestClient) Do(ctx context.Context, req *http.Request, resBody interface{}) error {
	klog.V(6).Infof("symbl.Do ENTER\n")

	var err error
	for i := 1; i <= defaultAttemptsToReauth; i++ {
		// delay on subsequent calls
		if i > 1 {
			klog.V(4).Info("Sleep for retry...\n")
			time.Sleep(time.Second * time.Duration(defaultDelayBetweenReauth))
		}

		// run request
		err = c.Client.Do(ctx, req, resBody)

		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode == http.StatusUnauthorized {

				klog.V(3).Info("Received http.StatusUnauthorized\n")
				newClient, reauthErr := NewRestClientWithCreds(ctx, *c.creds)
				if reauthErr != nil {
					klog.V(1).Infof("unable to re-authorize to symbl platform\n")
					klog.V(6).Infof("symbl.Do LEAVE\n")
					return reauthErr
				}

				klog.V(4).Info("Re-authorized with the symbl.ai platform\n")
				c.Client = newClient.Client
				c.auth = newClient.auth
			}
		} else {
			return err
		}
	}

	klog.V(1).Infof("Failed with (%s) %s\n", req.Method, req.URL)
	klog.V(6).Infof("symbl.Do LEAVE\n")
	return err
}
