// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package microphone

import (
	"os"
	"sync"

	"github.com/gordonklaus/portaudio"
)

// AudioConfig init config for library
type AudioConfig struct {
	InputChannels int
	SamplingRate  float32
}

// Microphone...
type Microphone struct {
	stream *portaudio.Stream

	intBuf []int16
	sig    chan os.Signal

	mute  sync.Mutex
	muted bool
}
