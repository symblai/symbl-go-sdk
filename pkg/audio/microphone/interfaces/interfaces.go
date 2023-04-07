// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

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
