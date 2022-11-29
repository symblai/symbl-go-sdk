// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

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
