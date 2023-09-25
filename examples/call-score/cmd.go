// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	async "github.com/symblai/symbl-go-sdk/pkg/api/async/v1"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

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

	// Process Call Score
	conversationJob, err := asyncClient.PostFileWithOptions(ctx, "newPhonecall.mp3", ufRequest)
	if err == nil {
		fmt.Printf("JobID: %s, ConversationID: %s\n\n", conversationJob.JobID, conversationJob.ConversationID)
	} else {
		fmt.Printf("PostFile failed. Err: %v\n", err)
		os.Exit(1)
	}

	// Wait for Processing (Wait for 20 minutes, increase if needed)
	for i := 0; i < 20; i++ {
		result, err := asyncClient.GetCallScoreStatusById(ctx, conversationJob.ConversationID)
		fmt.Printf("Current status (attempt %d): %s", i+1, result.Status)

		if err == nil && result.Status == "completed" {
			break
		}

		if err != nil {
			fmt.Printf("Error fetching status (attempt %d): %v", i+1, err)
		}

		// hardcoded retryDelay
		time.Sleep(time.Minute)
	}

	// Fetch the CallScore
	callScore, err := asyncClient.GetCallScore(ctx, conversationJob.ConversationID)
	if err == nil {
		fmt.Printf("Call Score: %v\n", callScore)
	} else {
		fmt.Printf("Fetch Call Score failed. Err: %v\n", err)
		// os.Exit(1)
	}

	// Fetch Insights List UI URL
	insightsListURL, err := asyncClient.GetInsightsListUiURI(ctx)
	if err == nil {
		fmt.Printf("Insights List URL: %s\n", insightsListURL)
	} else {
		fmt.Printf("Fetch Insights List URL failed. Err: %v\n", err)
		// os.Exit(1)
	}

	// Fetch Insights Details UI URL
	insightsDetailsURL, err := asyncClient.GetInsightsDetailsUiURI(ctx, conversationJob.ConversationID)
	if err == nil {
		fmt.Printf("Insights Details URL: %s\n", insightsDetailsURL)
	} else {
		fmt.Printf("Fetch Insights Details URL failed. Err: %v\n", err)
		// os.Exit(1)
	}

	fmt.Printf("\n\n")
	fmt.Printf("\n\n")

	fmt.Printf("Succeeded")
}
