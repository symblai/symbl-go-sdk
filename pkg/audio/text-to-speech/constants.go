// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package texttospeech

import (
	"errors"

	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

const (
	defaultBytesToRead int = 2048

	SpeechVoiceNeutral = texttospeechpb.SsmlVoiceGender_NEUTRAL
	SpeechVoiceFemale  = texttospeechpb.SsmlVoiceGender_FEMALE
	SpeechVoiceMale    = texttospeechpb.SsmlVoiceGender_MALE

	DefaultLanguageCode string = "en-US"
)

var (
	// ErrInvalidInput required input was not found
	ErrInvalidInput = errors.New("required input was not found")
)
