// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package management

import (
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

type Management struct {
	*symbl.RestClient
}

func New(client *symbl.RestClient) *Management {
	return &Management{client}
}
