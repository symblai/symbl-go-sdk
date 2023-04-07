// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

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
