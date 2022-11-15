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
	entitiesResult, err := mgmtClient.GetEntites(ctx)
	if err != nil {
		fmt.Printf("GetEntites failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(entitiesResult)
	fmt.Printf("\n")

	// create
	createEntity := interfaces.CreateEntityRequest{
		EntityArray: []interfaces.EntityRequest{
			interfaces.EntityRequest{
				Type:     "Vehicle",
				SubType:  "Honda",
				Category: "Custom",
				Values:   []string{"hrv", "crv"},
			},
		},
	}
	createResponse, err := mgmtClient.CreateEntity(ctx, createEntity)
	if err != nil {
		fmt.Printf("CreateEntity failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(createResponse)
	fmt.Printf("\n")

	// list again
	entitiesResult, err = mgmtClient.GetEntites(ctx)
	if err != nil {
		fmt.Printf("GetEntites failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(entitiesResult)
	fmt.Printf("\n")

	// update
	for _, entity := range entitiesResult.Entities {

		updateEntity := interfaces.Entity{
			Type:     entity.Type,
			SubType:  entity.SubType,
			Category: entity.Category,
			Values:   append(entity.Values, "crx"),
		}
		updateResponse, err := mgmtClient.UpdateEntity(ctx, entity.ID, updateEntity)
		if err != nil {
			fmt.Printf("UpdateEntity failed. Err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\n")
		spew.Dump(updateResponse)
		fmt.Printf("\n")
	}

	// list again, again
	entitiesResult, err = mgmtClient.GetEntites(ctx)
	if err != nil {
		fmt.Printf("GetEntites failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(entitiesResult)
	fmt.Printf("\n")

	// delete entities
	for _, entity := range entitiesResult.Entities {
		err = mgmtClient.DeleteEntity(ctx, entity.ID)
		if err != nil {
			fmt.Printf("DeleteEntity failed. Err: %v\n", err)
			os.Exit(1)
		}
	}

	// list again, again, again
	entitiesResult, err = mgmtClient.GetEntites(ctx)
	if err != nil {
		fmt.Printf("GetEntites failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(entitiesResult)
	fmt.Printf("\n")

	fmt.Printf("Succeeded")
}
