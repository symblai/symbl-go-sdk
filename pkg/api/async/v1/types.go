// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package async

/*
	Output structs for API calls
*/
// JobStatus captures the API for getting status
type JobStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// JobConversation represents processing an Async API request
type JobConversation struct {
	JobID          string `json:"jobId"`
	ConversationID string `json:"conversationId"`
}
