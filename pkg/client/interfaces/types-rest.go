// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

/*
	Symbl REST API
*/
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
