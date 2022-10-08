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

	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
)

const (
	defaultAuthType    string = "application"
	defaultAuthTimeout int64  = 5

	defaultAttemptsToReauth   int   = 3
	defaultDelayBetweenReauth int64 = 2
)

type RestClient struct {
	*rest.Client

	creds *Credentials
	auth  *authResp
}

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

	c := Credentials{
		AppId:     appId,
		AppSecret: appSecret,
	}
	return NewRestClientWithCreds(ctx, c)
}

// NewRestClientWithCreds creates a new client on the Symbl.ai platform. The client authenticates with the
// server with APP_ID/APP_SECRET.
func NewRestClientWithCreds(ctx context.Context, creds Credentials) (*RestClient, error) {
	klog.V(6).Infof("NewWithCreds ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(creds)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.Errorf("NewWithCreds validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("NewWithCreds LEAVE\n")
		return nil, err
	}

	if creds.Type == "" {
		creds.Type = defaultAuthType
	}

	// let's auth
	jsonStr, err := json.Marshal(creds)
	if err != nil {
		klog.Errorf("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("NewWithCreds LEAVE\n")
		return nil, err
	}

	// klog.V(6).Infof("------------------------\n")
	// klog.V(6).Infof("IMPORTANT: Never print in production\n")
	// klog.V(6).Infof("creds:\n%v\n", creds)
	// klog.V(6).Infof("------------------------\n")

	req, err := http.NewRequestWithContext(ctx, "POST", AuthURI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("NewWithCreds LEAVE\n")
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
		klog.Errorf("Symbl auth token is empty\n")
		klog.V(6).Infof("NewWithCreds LEAVE\n")
		return nil, ErrAuthFailure
	}

	// klog.V(6).Infof("------------------------\n")
	// klog.V(6).Infof("IMPORTANT: Never print in production\n")
	// klog.V(6).Infof("resp:\n%v\n", resp)
	// klog.V(6).Infof("------------------------\n")

	restClient.SetAuthorization(&rest.AccessToken{
		AccessToken: resp.AccessToken,
		ExpiresOn:   time.Now().Add(time.Second * time.Duration(resp.ExpiresIn)),
	})

	c := &RestClient{
		Client: restClient,
		creds:  &creds,
		auth:   &resp,
	}

	klog.V(4).Infof("NewWithCreds Succeeded\n")
	klog.V(6).Infof("NewWithCreds LEAVE\n")
	return c, nil
}

func (c *RestClient) DoFile(ctx context.Context, filePath string, resBody interface{}) error {
	return c.Client.DoFile(ctx, filePath, resBody)
}

func (c *RestClient) DoURL(ctx context.Context, url string, resBody interface{}) error {
	return c.Client.DoURL(ctx, url, resBody)
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

		if e, ok := err.(*rest.StatusError); ok {
			if e.Resp.StatusCode == http.StatusUnauthorized {

				klog.V(3).Info("Received http.StatusUnauthorized\n")
				newClient, reauthErr := NewRestClientWithCreds(ctx, *c.creds)
				if reauthErr != nil {
					klog.Errorf("unable to re-authorize to symbl platform\n")
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

	klog.Errorf("Failed with (%s) %s\n", req.Method, req.URL)
	klog.V(6).Infof("symbl.Do LEAVE\n")
	return err
}
