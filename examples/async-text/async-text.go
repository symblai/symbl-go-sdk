// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	prettyjson "github.com/hokaccha/go-prettyjson"

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
		async (text)
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

	// new convo
	messages := []string{
		"hello",
		"how are you doing today?",
		"the weather is sunny and hot in Long Beach",
	}

	jobConvo, err := asyncClient.PostText(ctx, messages)
	if err == nil {
		fmt.Printf("JobID: %s, ConversationID: %s\n\n", jobConvo.JobID, jobConvo.ConversationID)
	} else {
		fmt.Printf("PostText failed. Err: %v\n", err)
		os.Exit(1)
	}

	conversationId := jobConvo.ConversationID

	completed, err := asyncClient.WaitForJobComplete(ctx, interfaces.WaitForJobStatusOpts{JobId: jobConvo.JobID})
	if err != nil {
		fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
		os.Exit(1)
	}
	if !completed {
		fmt.Printf("WaitForJobComplete failed to complete. Use larger timeout\n")
		os.Exit(1)
	}

	topicsResult, err := asyncClient.GetTopics(ctx, jobConvo.ConversationID)
	if err != nil {
		fmt.Printf("Topics failed. Err: %v\n", err)
		os.Exit(1)
	}

	// print it
	byData, err := json.Marshal(topicsResult)
	if err != nil {
		fmt.Printf("RecognitionResult json.Marshal failed. Err: %v\n", err)
		os.Exit(1)
	}

	prettyJson, err := prettyjson.Format(byData)
	if err != nil {
		fmt.Printf("prettyjson.Marshal failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n\n")
	fmt.Printf("%s\n", prettyJson)
	fmt.Printf("\n\n")

	// append to convo
	messages = []string{
		"the weather is cold and rainy in Seattle today",
		"I think next week the weather is going to improve a lot",
	}

	jobConvo, err = asyncClient.PostAppendText(ctx, conversationId, messages)
	if err == nil {
		fmt.Printf("JobID: %s, ConversationID: %s\n\n", jobConvo.JobID, jobConvo.ConversationID)
	} else {
		fmt.Printf("PostAppendText failed. Err: %v\n", err)
		os.Exit(1)
	}

	completed, err = asyncClient.WaitForJobComplete(ctx, interfaces.WaitForJobStatusOpts{
		JobId:         jobConvo.JobID,
		WaitInSeconds: 120,
	})
	if err != nil {
		fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
		os.Exit(1)
	}
	if !completed {
		fmt.Printf("WaitForJobComplete failed to complete. Use larger timeout\n")
		os.Exit(1)
	}

	topicsResult, err = asyncClient.GetTopics(ctx, jobConvo.ConversationID)
	if err != nil {
		fmt.Printf("Topics failed. Err: %v\n", err)
		os.Exit(1)
	}

	// print it
	byData, err = json.Marshal(topicsResult)
	if err != nil {
		fmt.Printf("RecognitionResult json.Marshal failed. Err: %v\n", err)
		os.Exit(1)
	}

	prettyJson, err = prettyjson.Format(byData)
	if err != nil {
		fmt.Printf("prettyjson.Marshal failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n\n")
	fmt.Printf("%s\n", prettyJson)
	fmt.Printf("\n\n")

	fmt.Printf("Succeeded")
}
