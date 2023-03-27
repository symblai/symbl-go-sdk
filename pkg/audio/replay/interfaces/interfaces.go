// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import "io"

type Replay interface {
	Start() error
	Read() ([]byte, error)
	Stream(w io.Writer) error
	Mute()
	Unmute()
	Stop() error
}
