// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	management "github.com/symblai/symbl-go-sdk/pkg/api/management/v1"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/api/management/v1/interfaces"
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelTrace,
	})

	/*
		Entity manipulation
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
	groupResult, err := mgmtClient.GetConversationGroups(ctx)
	if err != nil {
		fmt.Printf("GetConversationGroups failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(groupResult)
	fmt.Printf("\n")

	// create
	createRequest := interfaces.Group{
		Name:        "Test1",
		Description: "MyDescription1",
		Criteria:    "conversation.metadata.agentId==johndoe",
	}

	createResponse, err := mgmtClient.CreateConversationGroup(ctx, createRequest)
	if err != nil {
		fmt.Printf("CreateConversationGroup failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(createResponse)
	fmt.Printf("\n")

	// list again
	groupResult, err = mgmtClient.GetConversationGroups(ctx)
	if err != nil {
		fmt.Printf("GetConversationGroups failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(groupResult)
	fmt.Printf("\n")

	// update
	for _, group := range groupResult.Groups {
		group.Description = "Updated Description"
		updateResponse, err := mgmtClient.UpdateConversationGroup(ctx, group)
		if err != nil {
			fmt.Printf("CreateConversationGroup failed. Err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\n")
		spew.Dump(updateResponse)
		fmt.Printf("\n")
	}

	// list again, again
	groupResult, err = mgmtClient.GetConversationGroups(ctx)
	if err != nil {
		fmt.Printf("GetConversationGroups failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(groupResult)
	fmt.Printf("\n")

	// delete entities
	for _, conversationGroup := range groupResult.Groups {
		err = mgmtClient.DeleteConversationGroup(ctx, conversationGroup.ID)
		if err != nil {
			fmt.Printf("DeleteEntity failed. Err: %v\n", err)
			os.Exit(1)
		}
	}

	// list again, again, again
	groupResult, err = mgmtClient.GetConversationGroups(ctx)
	if err != nil {
		fmt.Printf("GetConversationGroups failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(groupResult)
	fmt.Printf("\n")

	fmt.Printf("Succeeded")
}
