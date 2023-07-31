// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	prettyjson "github.com/hokaccha/go-prettyjson"

	async "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	var accessToken string
	flag.StringVar(&accessToken, "token", "", "Symbl.ai Access Token")
	flag.Parse()

	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	/*
		------------------------------------
		async (url)
		------------------------------------
	*/
	ctx := context.Background()

	restClient, err := symbl.NewRestClientWithToken(ctx, accessToken)
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

	messagesResult, err := asyncClient.GetMessages(ctx, jobConvo.ConversationID)
	if err != nil {
		fmt.Printf("Messages failed. Err: %v\n", err)
		os.Exit(1)
	}

	// print it
	byData, err := json.Marshal(messagesResult)
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

	fmt.Printf("Succeeded")
}
