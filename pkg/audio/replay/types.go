// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package replay

import (
	"os"
	"sync"

	wav "github.com/youpy/go-wav"
)

type ReplayOpts struct {
	FullFilename string
}

type Client struct {
	options ReplayOpts

	// wav
	file    *os.File
	decoder *wav.Reader

	// operational stuff
	stopChan chan struct{}
	mute     sync.Mutex
	muted    bool
}
