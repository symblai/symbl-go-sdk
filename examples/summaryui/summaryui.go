// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

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
		------------------------------------
		Summary UI
		------------------------------------
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

	jobConvo, err := asyncClient.PostURL(ctx, "https://symbltestdata.s3.us-east-2.amazonaws.com/newPhonecall.mp3")
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

	summaryResult, err := asyncClient.GetSummaryUI(ctx, jobConvo.ConversationID, "https://symbltestdata.s3.us-east-2.amazonaws.com/newPhonecall.mp3")
	if err != nil {
		fmt.Printf("GetSummaryUI failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n\n")
	spew.Dump(summaryResult)

	fmt.Printf("Succeeded")
}
