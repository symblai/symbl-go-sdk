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
	ManagementAPIVersion string = "v1"

	// trackers
	ManagementTrackerURI     string = "https://api.symbl.ai/%s/manage/trackers"
	ManagementTrackerByIdURI string = "https://api.symbl.ai/%s/manage/trackers/%s"

	// entity
	ManagementEntitiesURI          string = "https://api.symbl.ai/%s/manage/entities"
	ManagementEntitiesBulkURI      string = "https://api.symbl.ai/%s/manage/entities/bulk"
	ManagementEntitiesByIdURI      string = "https://api.symbl.ai/%s/manage/entities/%s"
	ManagementEntitiesBySubTypeURI string = "https://api.symbl.ai/%s/manage/entities?subType=%s"

	// conversation groups
	ManagementConversationGroupURI     string = "https://api.symbl.ai/%s/manage/group"
	ManagementConversationGroupsURI    string = "https://api.symbl.ai/%s/manage/groups"
	ManagementConversationGroupByIdURI string = "https://api.symbl.ai/%s/manage/group/%s"
)

func GetManagementAPI(URI string, args ...interface{}) string {
	return fmt.Sprintf(URI, append([]interface{}{ManagementAPIVersion}, args...)...)
}
