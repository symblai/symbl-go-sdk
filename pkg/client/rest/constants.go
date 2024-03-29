// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"errors"
)

var (
	// ErrInvalidInput required input was not found
	ErrInvalidInput = errors.New("required input was not found")

	// ErrInvalidURIExtension couldn't find a period to indicate a file extension
	ErrInvalidURIExtension = errors.New("couldn't find a period to indicate a file extension")
)
