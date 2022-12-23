// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

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
