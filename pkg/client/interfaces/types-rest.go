// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package interfaces

/*
	Symbl REST API
*/
// Credentials is the input needed to login to the Symbl.ai platform
type Credentials struct {
	AuthURI   string
	Type      string `json:"type"`
	AppId     string `json:"appId" validate:"required"`
	AppSecret string `json:"appSecret" validate:"required"`
}

// AuthResp represents a Symbl platform bearer access token with expiry information.
type AuthResp struct {
	AccessToken string `json:"accessToken"`
	NebulaToken string `json:"newbulaToken"`
	ExpiresIn   int    `json:"expiresIn"`
}
