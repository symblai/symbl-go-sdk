// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

// streaming
import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	nebula "github.com/dvonthenen/symbl-go-sdk/pkg/api/nebula/v1"
	microphone "github.com/dvonthenen/symbl-go-sdk/pkg/audio/microphone"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"

	handler "github.com/dvonthenen/symbl-go-sdk/examples/nebula/realtime/handler"
)

func main() {
	// init the library
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelStandard, // LogLevelStandard / LogLevelTrace
	})

	// context
	ctx := context.Background()

	// create the chatgpt client
	fmt.Printf("Connecting to Nebula...\n")
	restClient, err := symbl.NewNebulaRestClient(ctx)
	if err == nil {
		fmt.Println("Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	nebulaClient := nebula.New(restClient)
	fmt.Printf("Connection Succeeded\n")

	// init library
	microphone.Initialize()

	// init the handler
	msgHandler := handler.NewHandler(handler.HandlerOptions{
		NebulaClient: nebulaClient,
	})

	// create a new client
	symblConfig := symbl.GetDefaultConfig()
	symblConfig.Speaker.Name = "John Doe"
	symblConfig.Speaker.UserID = "john.doe@mymail.com"
	symblConfig.Config.DetectEntities = true
	symblConfig.Config.Sentiment = true

	options := symbl.StreamingOptions{
		SymblConfig: symblConfig,
		Callback:    msgHandler,
	}

	client, err := symbl.NewStreamClient(ctx, options)
	if err == nil {
		fmt.Println("Login Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("ConversationID: %s\n", client.GetConversationId())

	err = client.Start()
	if err == nil {
		fmt.Printf("Streaming Session Started!\n")
	} else {
		fmt.Printf("client.Start failed. Err: %v\n", err)
		os.Exit(1)
	}

	// delay...
	time.Sleep(time.Second * 1)

	// mic stuf
	mic, err := microphone.New(microphone.AudioConfig{
		InputChannels: 1,
		SamplingRate:  16000,
	})
	if err != nil {
		fmt.Printf("Initialize failed. Err: %v\n", err)
		os.Exit(1)
	}

	// start the mic
	err = mic.Start()
	if err != nil {
		fmt.Printf("mic.Start failed. Err: %v\n", err)
		os.Exit(1)
	}

	go func() {
		// this is a blocking call
		mic.Stream(client)
	}()

	fmt.Print("Press ENTER to exit!\n\n")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	// close stream
	err = mic.Stop()
	if err != nil {
		fmt.Printf("mic.Stop failed. Err: %v\n", err)
		os.Exit(1)
	}

	// teardown library
	microphone.Teardown()

	// close client
	client.Stop()

	fmt.Printf("Succeeded!\n\n")
}
