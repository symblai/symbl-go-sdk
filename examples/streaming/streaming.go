// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gordonklaus/portaudio"
	klog "k8s.io/klog/v2"

	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	symbl.InitLogging(6)

	/*
		------------------------------------
		streaming
		------------------------------------
	*/
	// websocket stuff
	ctx := context.Background()

	client, err := symbl.NewStreamClientWithDefaults(ctx)
	if err == nil {
		fmt.Println("Succeeded!")
	} else {
		fmt.Printf("New failed. Err: %v\n", err)
		os.Exit(1)
	}

	// delay...
	time.Sleep(time.Second * 3)

	// mic stuf
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	portaudio.Initialize()
	defer portaudio.Terminate()

	in := make([]int16, 1024)
	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(in), in)
	if err != nil {
		fmt.Println("OpenDefaultStream failed. Err: %v\n", err)
		os.Exit(1)
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		fmt.Printf("Mic failed to start. Err: %v\n", err)
		os.Exit(1)
	}

	for {
		err = stream.Read()
		if err != nil {
			klog.Errorf("stream.Read failed. Err: %v\n", err)
			os.Exit(1)
		}

		// doesnt work with example code
		err = client.WriteBinary(int16ToLittleEndianByte(in))
		if err != nil {
			klog.Errorf("client.WriteBinary failed. Err: %v\n", err)
			os.Exit(1)
		}

		select {
		case <-sig:
			return
		default:
		}
	}

	client.Stop()

	err = stream.Stop()
	if err != nil {
		klog.Errorf("stream.Stop failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Succeeded")
}

func int16ToLittleEndianByte(f []int16) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
