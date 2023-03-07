// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package rest

import (
	"time"

	simple "github.com/dvonthenen/symbl-go-sdk/pkg/client/simple"
)

// AccessToken represents a Symbl platform bearer access token with expiry information.
type AccessToken struct {
	AccessToken string
	ExpiresOn   time.Time
}

// Client which extends basic client to support REST
type Client struct {
	*simple.Client

	auth *AccessToken
}
