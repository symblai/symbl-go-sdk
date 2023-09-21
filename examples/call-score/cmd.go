// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	async "github.com/symblai/symbl-go-sdk/pkg/api/async/v1"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	/*
		------------------------------------
		async (file)
		------------------------------------
	*/
	ctx := context.Background()

	restClient, err := symbl.NewRestClient(ctx)
	if err == nil {
		fmt.Println("Succeeded!\n\n")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	asyncClient := async.New(restClient)

	ufRequest := interfaces.AsyncURLFileRequest{
		DetectEntities:           true,
		EnableSpeakerDiarization: true,
		DiarizationSpeakerCount:  2,
		ParentRefs:               true,
		Sentiment:                true,
		Mode:                     "default",
		Features: interfaces.Features{
			FeatureList: []string{"insights", "callScore"},
		},
		ConversationType: "sales",
		Metadata: interfaces.Metadata{
			SalesStage:   "general",
			ProspectName: "John Doe",
		},
	}
	jobConvo, err := asyncClient.PostFileWithOptions(ctx, "newPhonecall.mp3", ufRequest)
	if err == nil {
		fmt.Printf("JobID: %s, ConversationID: %s\n\n", jobConvo.JobID, jobConvo.ConversationID)
	} else {
		fmt.Printf("PostFile failed. Err: %v\n", err)
		os.Exit(1)
	}

	completed, err := asyncClient.WaitForJobComplete(ctx, interfaces.WaitForJobStatusOpts{JobId: jobConvo.JobID})
	if err != nil {
		fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
		os.Exit(1)
	}
	if !completed {
		fmt.Printf("WaitForJobComplete failed to complete. Use larger timeout\n")
		os.Exit(1)
	}

	fmt.Printf("\n\n")
	fmt.Printf("\n\n")

	fmt.Printf("Succeeded")
}
