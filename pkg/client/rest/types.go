// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package rest

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

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
