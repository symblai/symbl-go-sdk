// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
Package provides Go library for extracting Asynchoronous and Real-Time conversation insights
on the Symbl.ai platform.

GitHub repo: https://github.com/dvonthenen/symbl-go-sdk
Go SDK Examples: https://github.com/dvonthenen/symbl-go-sdk/tree/main/examples

Symbl Platform API reference: https://docs.symbl.ai/reference

The two main entry points are:
1. cli/cmd - which contains the unimplemented Symbl CLI
2. pkg/client - which contains the SDK objects, functions, etc
*/
package sdk

import (
	_ "github.com/dvonthenen/symbl-go-sdk/cli/cmd"
	_ "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)
