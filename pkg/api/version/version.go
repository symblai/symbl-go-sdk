// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package version

import (
	"fmt"
)

const (
	AsyncAPIVersion string = "v1"
)

const (
	JobStatusURI    string = "https://api.symbl.ai/%s/job/%s"
	ProcessAudioURI string = "https://api.symbl.ai/%s/process/audio?name=%s"
	TopicsURI       string = "https://api.symbl.ai/%s/conversations/%s/topics?parentRefs=true&sentiment=true"
)

func GetAsyncAPI(URI string, args ...interface{}) string {
	return fmt.Sprintf(URI, append([]interface{}{AsyncAPIVersion}, args...)...)
}
