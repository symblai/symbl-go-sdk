// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
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

	nebula "github.com/dvonthenen/symbl-go-sdk/pkg/api/nebula/v1"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/nebula/v1/interfaces"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	var accessToken string
	flag.StringVar(&accessToken, "token", "", "Symbl.ai Nebula Token")
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

	client, err := symbl.NewNebulaClientWithToken(ctx, accessToken)
	if err == nil {
		fmt.Println("Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	nebulaClient := nebula.New(client)

	request := interfaces.AskNebulaRequest{
		Instruction: "TODO",
		Conversation: interfaces.Conversation{
			Text: "TODO",
		},
	}

	nebulaResult, err := nebulaClient.AskNebula(ctx, request)
	if err != nil {
		fmt.Printf("AskNebula failed. Err: %v\n", err)
		os.Exit(1)
	}

	// print it
	byData, err := json.Marshal(nebulaResult)
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
