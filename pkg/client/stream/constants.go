// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package stream

import (
	"errors"
)

const (
	connectionRetryInfinite  int64 = 0
	defaultConnectRetry      int64 = 3
	defaultDelayBetweenRetry int64 = 2
)

var (
	// ErrInvalidConnection connection is not valid
	ErrInvalidConnection = errors.New("connection is not valid")
)
