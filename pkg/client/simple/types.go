// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package simple

import (
	"net/http"
)

// Client which extends HTTP client
type Client struct {
	http.Client

	d         *debugContainer
	UserAgent string
}
