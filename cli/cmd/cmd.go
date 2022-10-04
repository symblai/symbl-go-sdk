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
	"golang.org/x/exp/constraints"
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
		------------------------------------
		streaming
		------------------------------------
	*/
	// websocket stuff
	ctx := context.Background()

	client, err := symbl.NewStreamClient(ctx)
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

	// give this a try...
	/*
		start here
		Symbl impl
	*/
	// in := make([]byte, 8196)
	// stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(in), in) // old value: 44100
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
	// 	err := stream.Read()
	// 	if err != nil {
	// 		klog.Errorf("stream.Read failed. Err: %v\n", err)
	// 		os.Exit(1)
	// 	}

	// 	bufLength := len(in)
	// 	targetBuffer := make([]int16, bufLength)
	// 	for index := bufLength - 1; index >= 0; index -= 1 {
	// 		targetBuffer[index] = 32767 * min(1, in[index])
	// 	}

	// 	err = client.WriteBinary(targetBuffer)
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
	/*
		end here
		Symbl impl
	*/

	// THIS WORKS!!!!
	/*
		start here
		https://github.com/Raraku/go-transcription/blob/ac2336d4722f0cb356d298e44ff71da4bfad818d/main.go#L23
	*/
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
	/*
		end here
		https://github.com/Raraku/go-transcription/blob/ac2336d4722f0cb356d298e44ff71da4bfad818d/main.go#L23
	*/

	// this does not work...
	/*
		start here
		Variation 2: https://github.com/Raraku/go-transcription/blob/main/main.go
	*/
	// buf := make([]float32, 1024)
	// stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(buf), func(in []float32) {
	// 	for i := range buf {
	// 		buf[i] = in[i]
	// 	}
	// })
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

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	// done := make(chan struct{})

	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case <-ticker.C:
	// 		// doesnt work with either Little or Big Endian
	// 		err := client.WriteBinary(float32ToLittleEndianByte(buf))
	// 		if err != nil {
	// 			klog.Errorf("client.WriteBinary failed. Err: %v\n", err)
	// 			os.Exit(1)
	// 		}

	// 	case <-interrupt:
	// 		log.Println("interrupt")

	// 		select {
	// 		case <-done:
	// 		case <-time.After(time.Second):
	// 		}
	// 		return
	// 	}
	// }
	/*
		https://github.com/Raraku/go-transcription/blob/main/main.go
		end here
	*/

	// TESTING!!!
	/*
		start here
		Variation 3: https://github.com/Raraku/go-transcription/blob/main/main.go
	*/
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	// done := make(chan struct{})

	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case <-ticker.C:
	// 		// doesnt work with either Little or Big Endian
	// 		err := client.WriteBinary(float32ToBigEndianByte(buf))
	// 		if err != nil {
	// 			klog.Errorf("client.WriteBinary failed. Err: %v\n", err)
	// 			os.Exit(1)
	// 		}

	// 	case <-interrupt:
	// 		log.Println("interrupt")

	// 		select {
	// 		case <-done:
	// 		case <-time.After(time.Second):
	// 		}
	// 		return
	// 	}
	// }
	/*
		https://github.com/Raraku/go-transcription/blob/main/main.go
		end here
	*/

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

func int16ToBigEndianByte(f []int16) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func float32ToBigEndianByte(f []float32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func float32ToLittleEndianByte(f []float32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
