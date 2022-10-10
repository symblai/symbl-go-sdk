// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	management "github.com/dvonthenen/symbl-go-sdk/pkg/api/management/v1"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/management/v1/interfaces"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	/*
		Tracker manipulation
	*/
	ctx := context.Background()

	restClient, err := symbl.NewRestClient(ctx)
	if err == nil {
		fmt.Println("Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	mgmtClient := management.New(restClient)

	// list
	trackersResult, err := mgmtClient.GetTrackers(ctx)
	if err != nil {
		fmt.Printf("GetTrackers failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(trackersResult)
	fmt.Printf("\n")

	// create
	createTracker := interfaces.TrackerRequest{
		Name:       "Test1",
		Categories: []string{"cat1"},
		Languages:  []string{interfaces.TrackerLanguageDefault},
		Vocabulary: []string{"hello", "hi"},
	}
	createResponse, err := mgmtClient.CreateTracker(ctx, createTracker)
	if err != nil {
		fmt.Printf("CreateTracker failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(createResponse)
	fmt.Printf("\n")

	// list again
	trackersResult, err = mgmtClient.GetTrackers(ctx)
	if err != nil {
		fmt.Printf("GetTrackers failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(trackersResult)
	fmt.Printf("\n")

	// delete trackers
	for _, tracker := range trackersResult.Trackers {
		err = mgmtClient.DeleteTracker(ctx, tracker.ID)
		if err != nil {
			fmt.Printf("DeleteTracker failed. Err: %v\n", err)
			os.Exit(1)
		}
	}

	// list again, again
	trackersResult, err = mgmtClient.GetTrackers(ctx)
	if err != nil {
		fmt.Printf("GetTrackers failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(trackersResult)
	fmt.Printf("\n")

	fmt.Printf("Succeeded")
}
