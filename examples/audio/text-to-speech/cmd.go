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

	streaming "github.com/symblai/symbl-go-sdk/pkg/api/streaming/v1"
	tts "github.com/symblai/symbl-go-sdk/pkg/audio/text-to-speech"
	symbl "github.com/symblai/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelVerbose, // LogLevelStandard, LogLevelFull, LogLevelTrace, LogLevelVerbose
	})

	ctx := context.Background()

	// create a new client
	cfg := symbl.GetDefaultConfig()
	cfg.Config.SpeechRecognition.Encoding = "MULAW"
	cfg.Config.SpeechRecognition.SampleRateHertz = 8000
	cfg.Speaker.Name = "John Doe"
	cfg.Speaker.UserID = "john.doe@mymail.com"

	options := symbl.StreamingOptions{
		SymblConfig: cfg,
		Callback:    streaming.NewDefaultMessageRouter(),
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
	time.Sleep(time.Second * 5)

	// replay stuff
	play, err := tts.New(ctx, tts.SpeechOpts{
		Text: "Testing, 1, 2, 3.",
	})
	if err != nil {
		fmt.Printf("replay.New failed. Err: %v\n", err)
		os.Exit(1)
	}

	// start replay
	err = play.Start()
	if err != nil {
		fmt.Printf("replay.Start failed. Err: %v\n", err)
		os.Exit(1)
	}

	go func() {
		// this is a blocking call
		play.Stream(client)
	}()

	fmt.Print("Press ENTER to exit!\n\n")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	// close stream
	err = play.Stop()
	if err != nil {
		fmt.Printf("replay.Stop failed. Err: %v\n", err)
		os.Exit(1)
	}

	// close client
	client.Stop()

	fmt.Printf("Succeeded!\n\n")
}
