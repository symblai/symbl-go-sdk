// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	async "github.com/symblai/symbl-go-sdk/pkg/api/async/v1"
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

const (
	maxRetries     = 20                                          // Maximum number of retries to check for call score status
	retryInterval  = time.Minute                                 // Time to wait before next retry
	conversationID = "5740965687197696"                          // A conversation ID
	newMediaURL    = "https://publicly-accessible-audio-url.mp3" // New media URL for updating insights details
)

func main() {
	// Initialize the Symbl client
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	ctx := context.Background()

	// Create a new REST client
	restClient, err := symbl.NewRestClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create REST client. Error: %v\n", err)
	}
	fmt.Println("REST client created successfully.")

	asyncClient := async.New(restClient)

	// Check the status of CallScore and wait until it's completed
	waitForCallScoreCompletion(ctx, asyncClient)

	// Subsequent operations and their respective log messages
	performAsyncClientOperations(ctx, asyncClient)
}

// waitForCallScoreCompletion waits for the CallScoreStatus to be completed, retrying for defined times and interval.
func waitForCallScoreCompletion(ctx context.Context, asyncClient *async.Client) {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		result, err := asyncClient.GetCallScoreStatusById(ctx, conversationID)
		if err != nil {
			log.Printf("Error fetching status (attempt %d): %v\n", attempt, err)
		} else {
			fmt.Printf("Current status (attempt %d): %s\n", attempt, result.Status)
			if result.Status == "completed" {
				fmt.Println("CallScoreStatus is completed!")
				return
			}
		}
		time.Sleep(retryInterval)
	}
	log.Println("CallScoreStatus did not complete within the maximum retry limit.")
}

// performAsyncClientOperations performs various operations using the asyncClient and logs their outcomes.
func performAsyncClientOperations(ctx context.Context, asyncClient *async.Client) {
	if callScore, err := asyncClient.GetCallScore(ctx, conversationID); err != nil {
		log.Printf("Fetch Call Score failed. Error: %v\n", err)
	} else {
		fmt.Printf("Call Score: %v\n", callScore)
	}

	if insightsListURL, err := asyncClient.GetInsightsListUiURI(ctx); err != nil {
		log.Printf("Fetch Insights List URL failed. Error: %v\n", err)
	} else {
		fmt.Printf("Insights List URL: %s\n", insightsListURL)
	}

	if insightsDetailsURL, err := asyncClient.GetInsightsDetailsUiURI(ctx, conversationID); err != nil {
		log.Printf("Fetch Insights Details URL failed. Error: %v\n", err)
	} else {
		fmt.Printf("Insights Details URL: %s\n", insightsDetailsURL)
	}

	if err := asyncClient.UpdateMediaUrlForInsightsDetailsUI(ctx, conversationID, newMediaURL); err != nil {
		log.Printf("Update Media URL for Insights Details UI failed. Error: %v\n", err)
	} else {
		fmt.Println("Media URL for Insights Details UI updated successfully.")
	}

	fmt.Println("Operations Completed Successfully.")
}
