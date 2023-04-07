// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

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
