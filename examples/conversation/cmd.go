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
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	/*
		------------------------------------
		Conversations
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

	conversations, err := asyncClient.GetConversations(ctx)
	if err != nil {
		fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
		os.Exit(1)
	}

	if len(conversations.Conversations) > 0 {
		conversationID := ""
		for _, conversation := range conversations.Conversations {
			conversationID = conversation.ID
			break
		}

		conversation, err := asyncClient.GetConversation(ctx, conversationID)
		if err != nil {
			fmt.Printf("GetConversation failed. Err: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n\n")
		spew.Dump(conversation)

		fmt.Printf("Succeeded")
	} else {
		fmt.Printf("No conversations have been processed")
	}
}
