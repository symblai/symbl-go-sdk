// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

type InsightCallback interface {
	RecognitionResultMessage(rr *RecognitionResult) error
	MessageResponseMessage(mr *MessageResponse) error
	InsightResponseMessage(ir *InsightResponse) error
	UnhandledMessage(byMsg []byte) error
}
