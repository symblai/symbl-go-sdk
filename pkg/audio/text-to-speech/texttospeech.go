// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
 Implementation for text-to-speech
*/
package texttospeech

import (
	"bytes"
	"context"
	"io"
	"os"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/audio/text-to-speech/interfaces"
)

// New creates a new text-to-speech Client
func New(ctx context.Context, opts SpeechOpts) (*Client, error) {
	klog.V(6).Infof("TTSClient.New ENTER\n")

	if opts.LanguageCode == "" {
		opts.LanguageCode = DefaultLanguageCode
	}
	if opts.VoiceType == 0 {
		opts.VoiceType = SpeechVoiceNeutral
	}

	var googleApplicationCredentials string
	if v := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); v != "" {
		klog.V(4).Info("GOOGLE_APPLICATION_CREDENTIALS found")
		googleApplicationCredentials = v
	} else {
		klog.Error("GOOGLE_APPLICATION_CREDENTIALS not found")
		klog.V(6).Infof("TTSClient.New LEAVE\n")
		return nil, ErrInvalidInput
	}

	googleClient, err := texttospeech.NewClient(ctx)
	if err != nil {
		klog.V(1).Infof("texttospeech.NewClient failed. Err: %v\n", err)
		klog.V(6).Infof("TTSClient LEAVE\n")
		return nil, err
	}

	client := &Client{
		options:                      opts,
		speechClient:                 googleClient,
		googleApplicationCredentials: googleApplicationCredentials,
		stopChan:                     make(chan struct{}),
		muted:                        false,
	}

	klog.V(3).Infof("TTSClient.New Succeeded\n")
	klog.V(6).Infof("TTSClient.New LEAVE\n")

	return client, nil
}

// Start begins the audio playback of the converted text
func (c *Client) Start() error {
	klog.V(6).Infof("TTSClient.Start ENTER\n")
	klog.V(4).Infof("text: %s\n", c.options.Text)

	ctx := context.Background()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: c.options.Text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: c.options.LanguageCode,
			SsmlGender:   c.options.VoiceType,
		},
		// Select the type of audio file you want returned.
		// TODO: hardcoded since we only support MULAW currently
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding:   texttospeechpb.AudioEncoding_MULAW,
			SampleRateHertz: interfaces.DefaultSampleRateHertz,
		},
	}

	resp, err := c.speechClient.SynthesizeSpeech(ctx, &req)
	if err != nil {
		klog.V(1).Infof("speechClient.SynthesizeSpeech Failed. Err: %v\n", err)
		klog.V(6).Infof("TTSClient.TextToSpeech LEAVE\n")
		return err
	}

	// save to a reader
	klog.V(4).Infof("Bytes generated: %d\n", len(resp.AudioContent))
	c.byteBuf = bytes.NewReader(resp.AudioContent)

	klog.V(3).Infof("TTSClient.TextToSpeech Succeeded\n")
	klog.V(6).Infof("TTSClient.TextToSpeech LEAVE\n")
	return nil

}

// Read gets the raw bits of audio playback
func (c *Client) Read() ([]byte, error) {
	klog.V(7).Infof("byteBuf Size: %d\n", c.byteBuf.Len())
	buf := make([]byte, defaultBytesToRead)

	cnt, err := c.byteBuf.Read(buf)
	if err != nil {
		klog.V(1).Infof("TTSClient.Read failed. Err: %v\n", err)
		return []byte{}, err
	}
	klog.V(7).Infof("TTSClient.Read bytes copied: %d\n", cnt)

	return buf, nil
}

// Stream is a helper function to stream audio to a playback device
func (c *Client) Stream(w io.Writer) error {
	for {
		select {
		case <-c.stopChan:
			klog.V(6).Infof("stopChan signal exit\n")
			return nil
		default:
			byData, err := c.Read()
			if err == io.EOF {
				klog.V(4).Infof("c.Read EOF. Succeeded!\n")
				return nil
			}
			if err != nil {
				klog.V(1).Infof("c.Read failed. Err: %v\n", err)
				return err
			}

			c.mute.Lock()
			isMuted := c.muted
			c.mute.Unlock()

			if isMuted {
				klog.V(7).Infof("Mic is MUTED!\n")
				byData = make([]byte, len(byData))
			}

			byteCount, err := w.Write(byData)
			if err != nil {
				klog.V(1).Infof("w.Write failed. Err: %v\n", err)
				return err
			}
			klog.V(7).Infof("io.Writer succeeded. Bytes written: %d\n", byteCount)
		}
	}

	return nil
}

// Mute silences the audio playback
func (c *Client) Mute() {
	c.mute.Lock()
	c.muted = true
	c.mute.Unlock()
}

// Unmute restores the plyaback audio
func (c *Client) Unmute() {
	c.mute.Lock()
	c.muted = false
	c.mute.Unlock()
}

// Stop terminates the audio playback
func (c *Client) Stop() error {
	c.speechClient.Close()

	close(c.stopChan)
	<-c.stopChan

	return nil
}
