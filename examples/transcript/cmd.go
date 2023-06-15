// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	async "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	/*
		Entity manipulation
	*/
	ctx := context.Background()

	restClient, err := symbl.NewRestClient(ctx)
	if err == nil {
		fmt.Println("Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	asyncClient := async.New(restClient)

	// post url
	jobConvo, err := asyncClient.PostFile(ctx, "newPhonecall.mp3")
	if err == nil {
		fmt.Printf("JobID: %s, ConversationID: %s\n\n", jobConvo.JobID, jobConvo.ConversationID)
	} else {
		fmt.Printf("PostFile failed. Err: %v\n", err)
		os.Exit(1)
	}

	// wait
	completed, err := asyncClient.WaitForJobComplete(ctx, interfaces.WaitForJobStatusOpts{
		JobId: jobConvo.JobID,
		// TotalWaitInSeconds: 600,
		// WaitInSeconds:      5,
	})
	if err != nil {
		fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
		os.Exit(1)
	}
	if !completed {
		fmt.Printf("WaitForJobComplete failed to complete. Use larger timeout\n")
		os.Exit(1)
	}

	// create transcription
	getTranscript := interfaces.TranscriptRequest{
		ContentType: interfaces.TranscriptContentTypeSrt,
	}
	transcriptResponse, err := asyncClient.GetTranscript(ctx, jobConvo.ConversationID, getTranscript)
	if err != nil {
		fmt.Printf("CreateEntity failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(transcriptResponse)
	fmt.Printf("\n")

	// save to a file
	f, err := os.Create("transcript.srt")
	if err != nil {
		fmt.Printf("os.Create failed. Err: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err2 := f.WriteString(transcriptResponse.Transcript.Payload)
	if err2 != nil {
		fmt.Printf("f.WriteString failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Succeeded")
}
