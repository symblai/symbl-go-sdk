// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Implementation for a replay device. In this case, replays an audio file to stream into a mic
*/
package replay

import (
	"errors"
)

const (
	defaultBytesToRead int = 2048
)

var (
	// ErrInvalidInput required input was not found
	ErrInvalidInput = errors.New("required input was not found")
)
