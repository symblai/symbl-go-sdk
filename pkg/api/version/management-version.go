// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package version

import (
	"fmt"
)

const (
	ManagementAPIVersion string = "v1"

	ManagementTrackerURI             string = "https://api.symbl.ai/%s/manage/trackers"
	ManagementCreateTrackerURI       string = "https://api.symbl.ai/%s/manage/trackers"
	ManagementModifyDeleteTrackerURI string = "https://api.symbl.ai/%s/manage/trackers/%s"
)

func GetManagementAPI(URI string, args ...interface{}) string {
	return fmt.Sprintf(URI, append([]interface{}{ManagementAPIVersion}, args...)...)
}
