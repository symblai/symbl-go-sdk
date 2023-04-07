// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package interfaces

const (
	// message/insight types
	MessageTypeRecognitionResult string = "recognition_result"
	MessageTypeMessageResponse   string = "message_response"
	MessageTypeInsightResponse   string = "insight_response"
	MessageTypeTopicResponse     string = "topic_response"
	MessageTypeTrackerResponse   string = "tracker_response"
	MessageTypeEntityResponse    string = "entity_response"

	// user-defined messages
	MessageTypeUserDefined string = "user_defined"
)

const (
	InsightTypeQuestion   string = "question"
	InsightTypeFollowUp   string = "follow_up"
	InsightTypeActionItem string = "action_item"
)
