// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gordonklaus/portaudio"
	klog "k8s.io/klog/v2"

	// async "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "6")
	flag.Parse()

	/*
		------------------------------------
		async
		------------------------------------
	*/
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

	/*
		// TODO implement an unhandled message that can be passed along to the user
		klog.Errorf("Invalid Message Type: %s\n", smt.Message.Type)
		// b, err := json.MarshalIndent(string(byMsg), "", "    ")
		// if err != nil {
		// 	klog.V(6).Infof("SymblMessageRouter MarshalIndent failed. Err: %v\n", err)
		// 	klog.V(6).Infof("SymblMessageRouter LEAVE\n")
		// 	return err
		// }
		// klog.V(4).Infof("\n\n\n")
		// klog.V(4).Infof("New Object Type:\n")
		// klog.V(4).Infof("%s", string(b))
		// klog.V(4).Infof("\n\n\n")
		fmt.Printf("\n\n\n")
		fmt.Printf("New Object Type:\n")
		fmt.Printf("%s", string(byMsg))
		fmt.Printf("\n\n\n")
	*/

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

	// fmt.Printf("\n\n\n")
	// fmt.Printf("-----------------------------------------")
	// fmt.Print("Press 'Enter' to exit...")
	// fmt.Printf("-----------------------------------------")
	// fmt.Printf("\n\n\n")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	// THIS WORKS!!!!
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

	klog.Info("Succeeded")
}

func int16ToLittleEndianByte(f []int16) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
