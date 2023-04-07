// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package symbl

import (
	"errors"
)

const (
	defaultAuthURI string = "https://api.symbl.ai/oauth2/token:generate"
)

var (
	// ErrInvalidInput required input was not found
	ErrInvalidInput = errors.New("required input was not found")

	// ErrAuthFailure failed to authenticate to the symbl platform
	ErrAuthFailure = errors.New("failed to authenticate to the symbl platform")

	// ErrReauthFailure failed to re-authenticate to the symbl platform
	ErrReauthFailure = errors.New("failed to re-authenticate to the symbl platform")

	// ErrWebSocketInitializationFailed websocket initialization failed
	ErrWebSocketInitializationFailed = errors.New("websocket initialization failed")
)
