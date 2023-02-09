// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package rest

import (
	"time"
)

// AccessToken represents a Symbl platform bearer access token with expiry information.
type AccessToken struct {
	AccessToken string
	ExpiresOn   time.Time
}
