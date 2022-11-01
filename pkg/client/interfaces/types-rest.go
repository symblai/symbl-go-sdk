// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import (
	"fmt"

	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
)

/*
	Symbl REST API
*/
type HeadersContext struct{}

type StatusError struct {
	*rest.StatusError
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Resp.Request.Method, e.Resp.Request.URL, e.Resp.Status)
}

// Credentials is the input needed to login to the Symbl.ai platform
type Credentials struct {
	Type      string `json:"type"`
	AppId     string `json:"appId" validate:"required"`
	AppSecret string `json:"appSecret" validate:"required"`
}

// AuthResp represents a Symbl platform bearer access token with expiry information.
type AuthResp struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}
