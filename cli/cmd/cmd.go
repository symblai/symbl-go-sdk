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

	"github.com/davecgh/go-spew/spew"
	klog "k8s.io/klog/v2"

	management "github.com/dvonthenen/symbl-go-sdk/pkg/api/management/v1"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/management/v1/interfaces"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

// async
/*
import (
	"bytes"
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	klog "k8s.io/klog/v2"

	async "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1"
	"github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)
*/

// streaming
/*
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

	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)
*/

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "6")
	flag.Parse()

	/*
		Management manipulation
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

	// list again, again
	trackersResult, err = mgmtClient.GetTrackers(ctx)
	if err != nil {
		fmt.Printf("GetTrackers failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(trackersResult)
	fmt.Printf("\n")

	for _, tracker := range trackersResult.Trackers {
		err = mgmtClient.DeleteTracker(ctx, tracker.ID)
		if err != nil {
			fmt.Printf("DeleteTracker failed. Err: %v\n", err)
			os.Exit(1)
		}
	}

	// list again, again, again
	trackersResult, err = mgmtClient.GetTrackers(ctx)
	if err != nil {
		fmt.Printf("GetTrackers failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
	spew.Dump(trackersResult)
	fmt.Printf("\n")

	klog.Info("Succeeded")

	/*
		------------------------------------
		async (file)
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

	// completed, err := asyncClient.WaitForJobComplete(ctx, interfaces.WaitForJobStatusOpts{JobId: jobConvo.JobID})
	// if err != nil {
	// 	fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }
	// if !completed {
	// 	fmt.Printf("WaitForJobComplete failed to complete. Use larger timeout\n")
	// 	os.Exit(1)
	// }

	// topicsResult, err := asyncClient.GetTopics(ctx, jobConvo.ConversationID)
	// if err != nil {
	// 	fmt.Printf("Topics failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// spew.Dump(topicsResult)

	// klog.Info("Succeeded")

	/*
		------------------------------------
		async (url)
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

	// jobConvo, err := asyncClient.PostURL(ctx, "https://symbltestdata.s3.us-east-2.amazonaws.com/newPhonecall.mp3")
	// if err == nil {
	// 	fmt.Printf("JobID: %s, ConversationID: %s\n", jobConvo.JobID, jobConvo.ConversationID)
	// } else {
	// 	fmt.Printf("PostFile failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// completed, err := asyncClient.WaitForJobComplete(ctx, interfaces.WaitForJobStatusOpts{JobId: jobConvo.JobID})
	// if err != nil {
	// 	fmt.Printf("WaitForJobComplete failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }
	// if !completed {
	// 	fmt.Printf("WaitForJobComplete failed to complete. Use larger timeout\n")
	// 	os.Exit(1)
	// }

	// topicsResult, err := asyncClient.GetTopics(ctx, jobConvo.ConversationID)
	// if err != nil {
	// 	fmt.Printf("Topics failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// spew.Dump(topicsResult)

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
	// // websocket stuff
	// ctx := context.Background()

	// client, err := symbl.NewStreamClientWithDefaults(ctx)
	// if err == nil {
	// 	fmt.Println("Succeeded!")
	// } else {
	// 	fmt.Printf("New failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// // delay...
	// time.Sleep(time.Second * 3)

	// // mic stuf
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, os.Interrupt, os.Kill)

	// portaudio.Initialize()
	// defer portaudio.Terminate()

	// in := make([]int16, 1024)
	// stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(in), in)
	// if err != nil {
	// 	fmt.Println("OpenDefaultStream failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer stream.Close()

	// err = stream.Start()
	// if err != nil {
	// 	fmt.Printf("Mic failed to start. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// for {
	// 	err = stream.Read()
	// 	if err != nil {
	// 		klog.Errorf("stream.Read failed. Err: %v\n", err)
	// 		os.Exit(1)
	// 	}

	// 	// doesnt work with example code
	// 	err = client.WriteBinary(int16ToLittleEndianByte(in))
	// 	if err != nil {
	// 		klog.Errorf("client.WriteBinary failed. Err: %v\n", err)
	// 		os.Exit(1)
	// 	}

	// 	select {
	// 	case <-sig:
	// 		return
	// 	default:
	// 	}
	// }

	// client.Stop()

	// err = stream.Stop()
	// if err != nil {
	// 	klog.Errorf("stream.Stop failed. Err: %v\n", err)
	// 	os.Exit(1)
	// }

	// klog.Info("Succeeded")
}

func int16ToLittleEndianByte(f []int16) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
