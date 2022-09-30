// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"

	klog "k8s.io/klog/v2"

	// async "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "6")
	flag.Parse()

	// // async
	// ctx := context.Background()

	// restClient, err := symbl.NewRestClient(ctx)
	// if err == nil {
	// 	fmt.Println("Succeeded!")
	// } else {
	// 	fmt.Printf("New failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// asyncClient := async.New(restClient)

	// jobConvo, err := asyncClient.PostFile(ctx, "newPhonecall.mp3")
	// if err == nil {
	// 	fmt.Printf("JobID: %s, ConversationID: %s\n", jobConvo.JobID, jobConvo.ConversationID)
	// } else {
	// 	fmt.Printf("PostFile failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// completed, err := asyncClient.WaitForJobComplete(ctx, async.WaitForJobStatusOpts{JobId: jobConvo.JobID})
	// if err != nil {
	// 	fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }
	// if !completed {
	// 	fmt.Printf("WaitForJobComplete failed to complete. Use larger timeout\n")
	// 	os.Exit(1)
	// }

	// _, err = asyncClient.GetTopics(ctx, jobConvo.ConversationID)
	// if err != nil {
	// 	fmt.Printf("Topics failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// klog.Info("Succeeded")

	// streaming
	ctx := context.Background()

	client, err := symbl.NewStreamClient(ctx, nil)
	if err == nil {
		fmt.Println("Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n\n\n")
	fmt.Printf("-----------------------------------------")
	fmt.Print("Press 'Enter' to exit...")
	fmt.Printf("-----------------------------------------")
	fmt.Printf("\n\n\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	client.Stop()

	klog.Info("Succeeded")
}
