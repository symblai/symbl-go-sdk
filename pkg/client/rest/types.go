// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"time"

	simple "github.com/dvonthenen/symbl-go-sdk/pkg/client/simple"
)

// AccessToken represents a Symbl platform bearer access token with expiry information.
type AccessToken struct {
	AccessToken string
	NebulaToken string
	ExpiresOn   time.Time
}

// Client which extends basic client to support REST
type Client struct {
	*simple.Client

	auth *AccessToken
}
