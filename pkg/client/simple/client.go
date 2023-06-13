// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package simple

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

const (
	ClientName = "symbl.ai"
)

// defaultUserAgent is the default user agent string
var defaultUserAgent = fmt.Sprintf(
	"%s (%s)",
	ClientName,
	strings.Join([]string{runtime.Version(), runtime.GOOS, runtime.GOARCH}, ";"),
)

// New allocated a Simple HTTP client
func New() *Client {
	// TODO: add verification later, pick up from ENV or FILE
	/* #nosec G402 */
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := Client{
		Client: http.Client{
			Transport: tr,
		},
		d:         newDebug(),
		UserAgent: defaultUserAgent,
	}
	return &c
}

// Do performs a simple HTTP-style call
func (c *Client) Do(ctx context.Context, req *http.Request, f func(*http.Response) error) error {
	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// Create debugging context for this round trip
	d := c.d.newRoundTrip()
	if d.enabled() {
		defer d.done()
	}

	req.Header.Set("User-Agent", c.UserAgent)

	ext := ""
	if d.enabled() {
		ext = d.debugRequest(req)
	}

	tstart := time.Now()
	res, err := c.Client.Do(req.WithContext(ctx))
	tstop := time.Now()

	if d.enabled() {
		name := fmt.Sprintf("%s %s", req.Method, req.URL)
		d.logf("%6dms (%s)", tstop.Sub(tstart)/time.Millisecond, name)
	}

	if err != nil {
		return err
	}

	if d.enabled() {
		d.debugResponse(res, ext)
	}

	defer res.Body.Close()
	return f(res)
}
