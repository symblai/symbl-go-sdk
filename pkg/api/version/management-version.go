// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package version

import (
	"fmt"
)

const (
	ManagementAPIVersion string = "v1"

	// trackers
	ManagementTrackerURI     string = "https://api.symbl.ai/%s/manage/trackers"
	ManagementtrackerByIdURI string = "https://api.symbl.ai/%s/manage/trackers/%s"

	// entity
	ManagementEntitiesURI          string = "https://api.symbl.ai/%s/manage/entities"
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
