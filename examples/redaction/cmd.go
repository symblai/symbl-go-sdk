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

	async "github.com/symblai/symbl-go-sdk/pkg/api/async/v1"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
	cfginterfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
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

	// jobConvo, err := asyncClient.PostFile(ctx, "phoneNumber.mp3")
	ufRequest := interfaces.AsyncURLFileRequest{
		DetectEntities:           true,
		EnableSpeakerDiarization: true,
		DiarizationSpeakerCount:  2,
		ParentRefs:               true,
		Sentiment:                true,
	}
	jobConvo, err := asyncClient.PostFileWithOptions(ctx, "phoneNumber.mp3", ufRequest)
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

	// custom headers to enable options
	params := make(map[string][]string, 0)
	// params["exclude"] = []string{"[\"PERSON_NAME\"]"}
	// params["exclude"] = []string{"[\"PHONE_NUMBER\"]"}
	// params["exclude"] = []string{"[\"PERSON_NAME\"]", "[\"PHONE_NUMBER\"]"}
	params["exclude"] = []string{"\"PERSON_NAME\"", "\"PHONE_NUMBER\""}
	params["redact"] = []string{"true"}
	ctx = cfginterfaces.WithCustomParameters(ctx, params)

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
