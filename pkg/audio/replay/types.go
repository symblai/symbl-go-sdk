// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

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
