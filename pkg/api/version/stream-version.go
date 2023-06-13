// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
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
	StreamAPIVersion string = "v1"

	// StreamPath string = "%s/streaming/%s?access_token=%s"
	// StreamPath string = "%s/realtime/insights/%s?access_token=%s" // this is bad for library ? -> %3F
	StreamPath string = "%s/realtime/insights/%s"
)

func GetStreamingAPI(URI string, args ...interface{}) string {
	return fmt.Sprintf(URI, append([]interface{}{StreamAPIVersion}, args...)...)
}
