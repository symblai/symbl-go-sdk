// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import "io"

type Microphone interface {
	Start() error
	Read() ([]int16, error)
	Stream(w io.Writer) error
	Mute()
	Unmute()
	Stop() error
}
