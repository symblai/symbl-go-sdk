// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	simple "github.com/dvonthenen/symbl-go-sdk/pkg/client/simple"
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

func (c *Client) Do(ctx context.Context, req *http.Request, resBody interface{}) error {
	switch req.Method {
	case http.MethodPost, http.MethodPatch, http.MethodPut:
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	// TODO: verify this is correct
	if c.auth != nil && c.auth.AccessToken != "" {
		req.Header.Set("Authorization", c.auth.AccessToken)
	}

	if headers, ok := ctx.Value(HeadersContext{}).(http.Header); ok {
		for k, v := range headers {
			for _, v := range v {
				req.Header.Add(k, v)
			}
		}
	}

	return c.Client.Do(ctx, req, func(res *http.Response) error {
		switch res.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusNoContent:
		case http.StatusBadRequest:
			// TODO: structured error types
			detail, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			return fmt.Errorf("%s: %s", res.Status, bytes.TrimSpace(detail))
		default:
			return &StatusError{res}
		}

		if resBody == nil {
			return nil
		}

		switch b := resBody.(type) {
		case *RawResponse:
			return res.Write(b)
		case io.Writer:
			_, err := io.Copy(b, res.Body)
			return err
		default:
			d := json.NewDecoder(res.Body)
			return d.Decode(resBody)
		}
	})
}
