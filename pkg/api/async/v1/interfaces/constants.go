// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Defines everything that makes up the Async API interface
*/
package interfaces

const (
	// transcript options
	TranscriptContentTypeMarkdown string = "text/markdown"
	TranscriptContentTypeSrt      string = "text/srt"

	// speaker update
	SpeakerEventTypeStart   string = "started_speaking"
	SpeakerEventTypeStopped string = "stopped_speaking"
)
