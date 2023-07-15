// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Nebula package for processing Nebula Async conversations
*/
package nebula

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	nebulainterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/nebula/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	client "github.com/dvonthenen/symbl-go-sdk/pkg/client"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

// Context switch for processing Async functionality
type Client struct {
	*client.NebulaClient
}

// New changes the context of the REST client to an Async client
func New(client *client.NebulaClient) *Client {
	return &Client{client}
}

// AskNebula obtains conversation insights from nebula
func (c *Client) AskNebula(ctx context.Context, request nebulainterfaces.AskNebulaRequest) (*nebulainterfaces.AskNebulaResponse, error) {
	klog.V(6).Infof("nebula.AskNebula ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("AskNebula validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("nebula.AskNebula LEAVE\n")
		return nil, err
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetNebulaAsyncAPI(version.AskNebulaURI),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("nebula.AskNebula LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("nebula.AskNebula LEAVE\n")
		return nil, err
	}

	// check the status
	var result nebulainterfaces.AskNebulaResponse

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("nebula.AskNebula LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("nebula.AskNebula LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET AskNebula Succeeded\n")
	klog.V(6).Infof("nebula.AskNebula LEAVE\n")
	return &result, nil
}
