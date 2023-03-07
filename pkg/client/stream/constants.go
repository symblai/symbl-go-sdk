// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package stream

import (
	"errors"
	"time"
)

const (
	pingPeriod = 30 * time.Second
)

var (
	// ErrConnectionNoEstablished connection is not established
	ErrConnectionNoEstablished = errors.New("connection is not established")
)
