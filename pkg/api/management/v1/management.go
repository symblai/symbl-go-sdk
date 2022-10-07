// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package management

import (
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

type Management struct {
	*symbl.RestClient
}

func New(client *symbl.RestClient) *Management {
	return &Management{client}
}
