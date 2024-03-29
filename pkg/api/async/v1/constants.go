// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Async package for processing Async conversations
*/
package async

import (
	"errors"
)

const (
	JobStatusInProgress string = "in_progress"
	JobStatusComplete   string = "completed"
)

var (
	// ErrInvalidInput required input was not found
	ErrInvalidInput = errors.New("required input was not found")

	// ErrJobStatusTimeout the job status check timed out
	ErrJobStatusTimeout = errors.New("the job status check timed out")

	// ErrInvalidURIExtension couldn't find a period to indicate a file extension
	ErrInvalidURIExtension = errors.New("couldn't find a period to indicate a file extension")
)
