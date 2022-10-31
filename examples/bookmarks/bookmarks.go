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
	createBookmark := interfaces.BookmarkByMessageRefsRequest{
		Label:       "TODO",
		Description: "TODO",
		User: interfaces.User{
			Name:   "David",
			UserID: "TODO",
			Email:  "david.vonthenen@symbl.ai",
		},
		// BeginTimeOffset: 22,
		// Duration:        33,
		MessageRefs: []interfaces.MessageRefRequest{
			interfaces.MessageRefRequest{
				ID: "4510581827043328",
			},
		},
	}
	createResponse, err := asyncClient.CreateBookmarkByMessageRefs(ctx, conversationId, createBookmark)
	if err != nil {
		fmt.Printf("CreateEntity failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(createResponse)
	fmt.Printf("\n")

	// list again
	bookmarkResult, err = asyncClient.GetBookmarks(ctx, conversationId)
	if err != nil {
		fmt.Printf("GetBookmarks failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(bookmarkResult)
	fmt.Printf("\n")

	// delete entities
	for _, bookmark := range bookmarkResult.Bookmarks {
		err = asyncClient.DeleteBookmark(ctx, conversationId, bookmark.ID)
		if err != nil {
			fmt.Printf("DeleteEntity failed. Err: %v\n", err)
			os.Exit(1)
		}
	}

	// list again, again
	bookmarkResult, err = asyncClient.GetBookmarks(ctx, conversationId)
	if err != nil {
		fmt.Printf("GetBookmarks failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(bookmarkResult)
	fmt.Printf("\n")

	fmt.Printf("Succeeded")
}
