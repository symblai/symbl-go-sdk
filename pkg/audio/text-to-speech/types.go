// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package texttospeech

import (
	"bytes"
	"sync"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type SpeechOpts struct {
	VoiceType    texttospeechpb.SsmlVoiceGender
	LanguageCode string
	Text         string
}

type Client struct {
	options SpeechOpts

	// google stuff
	speechClient                 *texttospeech.Client
	googleApplicationCredentials string

	// operational stuff
	stopChan chan struct{}
	mute     sync.Mutex
	muted    bool

	// raw buffer
	byteBuf *bytes.Reader
}
