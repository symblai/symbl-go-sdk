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
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	/*
		Bookmark manipulation
	*/
	conversationId := "6558697145237504"

	ctx := context.Background()

	restClient, err := symbl.NewRestClient(ctx)
	if err == nil {
		fmt.Println("Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	asyncClient := async.New(restClient)

	// list
	bookmarkResult, err := asyncClient.GetBookmarks(ctx, conversationId)
	if err != nil {
		fmt.Printf("GetBookmarks failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(bookmarkResult)
	fmt.Printf("\n")

	// create
	createBookmark := interfaces.BookmarkRequest{
		Label:       "MyLabel",
		Description: "MyDescription",
		User: interfaces.User{
			Name:   "David",
			UserID: "MyUserId",
			Email:  "david.vonthenen@symbl.ai",
		},
		// You can use this below...
		// BeginTimeOffset: 22,
		// Duration:        33,
		// Or this below...
		MessageRefs: []interfaces.MessageRefRequest{
			interfaces.MessageRefRequest{
				ID: "4510581827043328",
			},
		},
	}
	createResponse, err := asyncClient.CreateBookmark(ctx, conversationId, createBookmark)
	if err != nil {
		fmt.Printf("CreateEntity failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(createResponse)
	fmt.Printf("\n")

	// get bookmark summary
	bookmarkSummary, err := asyncClient.GetSummaryOfBookmark(ctx, conversationId, createResponse.ID)
	if err != nil {
		fmt.Printf("CreateEntity failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n")
	spew.Dump(bookmarkSummary)
	fmt.Printf("\n")

	// list again
	bookmarkResult, err = asyncClient.GetBookmarks(ctx, conversationId)
	if err != nil {
		fmt.Printf("GetBookmarks failed. Err: %v\n", err)
		os.Exit(1)
	}

	// delete entities
	for _, bookmark := range bookmarkResult.Bookmarks {
		err = asyncClient.DeleteBookmark(ctx, conversationId, bookmark.ID)
		if err != nil {
			fmt.Printf("DeleteEntity failed. Err: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Succeeded")
}
