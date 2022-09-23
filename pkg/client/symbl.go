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

	common "github.com/dvonthenen/symbl-go-sdk/pkg/client/common"
	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
)

const (
	defaultAuthType    string = "application"
	defaultAuthTimeout int64  = 5

	defaultAttemptsToReauth   int   = 3
	defaultDelayBetweenReauth int64 = 2
)

type Client struct {
	*rest.Client

	creds *Credentials
}

// Credentials is the input needed to login to the Symbl.ai platform
type Credentials struct {
	Type      string `json:"type"`
	AppId     string `json:"appId" validate:"required"`
	AppSecret string `json:"appSecret" validate:"required"`
}

// authResp represents a Symbl platform bearer access token with expiry information.
type authResp struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}

// NewClient creates a new client on the Symbl.ai platform. The client authenticates with the
// server with APP_ID/APP_SECRET.
func New(ctx context.Context) (*Client, error) {
	var appId string
	if v := os.Getenv("APP_ID"); v != "" {
		klog.V(4).Info("APP_ID found")
		appId = v
	} else {
		klog.Errorln("APP_ID not found")
		return nil, common.ErrInvalidInput
	}
	var appSecret string
	if v := os.Getenv("APP_SECRET"); v != "" {
		klog.V(4).Info("APP_SECRET found")
		appSecret = v
	} else {
		klog.Errorln("APP_SECRET not found")
		return nil, common.ErrInvalidInput
	}

	c := Credentials{
		AppId:     appId,
		AppSecret: appSecret,
	}
	return NewWithCreds(ctx, c)
}

// NewClientWithCreds creates a new client on the Symbl.ai platform. The client authenticates with the
// server with APP_ID/APP_SECRET.
func NewWithCreds(ctx context.Context, creds Credentials) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
		// ctx, cancel := context.WithTimeout(ctx, time.Second*defaultAuthTimeout)
		// defer cancel()
	}

	// validate input
	v := validator.New()
	err := v.Struct(creds)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.Errorln(e)
		}
		return nil, err
	}

	if creds.Type == "" {
		creds.Type = defaultAuthType
	}

	// let's auth
	jsonStr, err := json.Marshal(creds)
	if err != nil {
		klog.Errorf("json.Marshal failed. Err: %v\n", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", common.AuthURI, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		return nil, err
	}

	// do it!
	var resp authResp

	restClient := rest.New()
	err = restClient.Do(ctx, req, &resp)
	if err != nil {
		klog.Errorf("restClient.Do failed. Err: %v\n", err)
		return nil, err
	}

	if resp.AccessToken == "" {
		klog.Errorf("Symbl auth token is empty")
		return nil, common.ErrAuthFailure
	}

	restClient.SetAuthorization(&rest.AccessToken{
		AccessToken: resp.AccessToken,
		ExpiresOn:   time.Now().Add(time.Second * time.Duration(resp.ExpiresIn)),
	})

	c := &Client{
		Client: restClient,
		creds:  &creds,
	}

	return c, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, resBody interface{}) error {
	var err error
	for i := 1; i <= defaultAttemptsToReauth; i++ {
		err = c.Client.Do(ctx, req, resBody)

		if i == 1 {
			klog.V(4).Info("Sleep for retry...")
			time.Sleep(time.Second * time.Duration(defaultDelayBetweenReauth))
		}

		if e, ok := err.(*rest.StatusError); ok {
			if e.Resp.StatusCode == http.StatusUnauthorized {

				klog.V(4).Info("Received http.StatusUnauthorized")
				newClient, reauthErr := NewWithCreds(ctx, *c.creds)
				if reauthErr != nil {
					klog.Errorf("unable to re-authorize to symbl platform")
					return reauthErr
				}

				klog.V(2).Info("Re-authorized with the symbl.ai platform")
				c.Client = newClient.Client
			}
		} else {
			return err
		}
	}

	klog.V(2).Infof("Failed with (%s) %s\n", req.Method, req.URL)
	return err
}
