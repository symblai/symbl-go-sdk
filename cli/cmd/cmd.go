// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package main

// streaming
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	klog "k8s.io/klog/v2"

	microphone "github.com/dvonthenen/symbl-go-sdk/pkg/audio/microphone"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.Init(symbl.SybmlInit{
		LogLevel: symbl.LogLevelStandard,
	})

	ctx := context.Background()

	// create a new client
	client, err := symbl.NewStreamClientWithDefaults(ctx)
	if err == nil {
		fmt.Println("Login Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	// delay...
	time.Sleep(time.Second * 3)

	// mic stuf
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	mic, err := microphone.Initialize(microphone.AudioConfig{
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

	// this is a blocking call
	mic.Stream(client)

	// close stream
	err = mic.Stop()
	if err != nil {
		klog.V(1).Infof("mic.Stop failed. Err: %v\n", err)
		os.Exit(1)
	}
	microphone.Teardown()

	// close client
	client.Stop()

	fmt.Printf("Succeeded!\n\n")
}
