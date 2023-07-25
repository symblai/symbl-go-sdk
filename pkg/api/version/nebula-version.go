// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	This package handles the versioning in the API both async and streaming
*/
package version

import (
	"fmt"
)

const (
	NebulaAsyncAPIVersion string = "v1"

	// processing audio
	AskNebulaURI string = "https://api-nebula.symbl.ai/%s/model/generate"
)

func GetNebulaAsyncAPI(URI string, args ...interface{}) string {
	return fmt.Sprintf(URI, append([]interface{}{NebulaAsyncAPIVersion}, args...)...)
}
