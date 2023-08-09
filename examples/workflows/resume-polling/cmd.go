// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	async "github.com/symblai/symbl-go-sdk/pkg/api/async/v1"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

func main() {
	conversationId := "4687270580060160"
	jobId := "e1568a02-533b-4ecb-b5db-21b18446dbc4"

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

	// wait
	completed, err := asyncClient.WaitForJobComplete(ctx, interfaces.WaitForJobStatusOpts{
		JobId:              jobId,
		TotalWaitInSeconds: 3600,
		WaitInSeconds:      30,
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
	transcriptResponse, err := asyncClient.GetTranscript(ctx, conversationId, getTranscript)
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
