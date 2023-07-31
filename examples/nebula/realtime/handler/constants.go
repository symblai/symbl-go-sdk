// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package handler

import "errors"

// const (
// 	// default number of message to cache at any given time
// 	DefaultNumOfMsgToCache int = 50
// )

var (
	// ErrUnhandledMessage runhandled message from nebula example handler
	ErrUnhandledMessage = errors.New("unhandled message from nebula example handler")

	// ErrItemNotFound item not found
	ErrItemNotFound = errors.New("item not found")
)
