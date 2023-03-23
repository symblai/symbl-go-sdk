// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package replay

import (
	"io"
	"os"

	wav "github.com/youpy/go-wav"
	klog "k8s.io/klog/v2"
)

func New(opts ReplayOpts) (*Client, error) {
	klog.V(6).Infof("Replay.New ENTER\n")

	client := &Client{
		options:  opts,
		stopChan: make(chan struct{}),
		muted:    false,
	}

	// create wav decoder instance
	f, err := os.Open(opts.FullFilename)
	if err != nil {
		klog.V(1).Infof("ReplayClient.New os.Open failed. Err: %v\n", err)
		klog.V(6).Infof("Replay.New LEAVE\n")
		return nil, err
	}

	// housekeeping
	client.file = f

	klog.V(3).Infof("Replay.New Succeeded\n")
	klog.V(6).Infof("Replay.New LEAVE\n")

	return client, nil
}

func (c *Client) Start() error {
	reader := wav.NewReader(c.file)
	if reader == nil {
		klog.V(1).Infof("ReplayClient.New wav.NewDecoder is nil\n")
		klog.V(6).Infof("Replay.New LEAVE\n")
		return ErrInvalidInput
	}

	// housekeeping
	c.decoder = reader

	return nil
}

func (c *Client) Read() ([]byte, error) {
	buf := make([]byte, defaultBytesToRead)

	byteCount, err := c.decoder.Read(buf)
	if err != nil {
		klog.V(1).Infof("byteBuf.Read failed. Err: %v\n", err)
		return []byte{}, err
	}
	klog.V(7).Infof("byteBuf.Read bytes copied: %d\n", byteCount)

	return buf, nil
}

func (c *Client) Stream(w io.Writer) error {
	for {
		select {
		case <-c.stopChan:
			klog.V(6).Infof("stopChan signal exit\n")
			return nil
		default:
			byData, err := c.Read()
			if err == io.EOF {
				klog.V(6).Infof("decoder.Read EOF\n")
				return nil
			}
			if err != nil {
				klog.V(1).Infof("decoder.Read failed. Err: %v\n", err)
				return err
			}

			c.mute.Lock()
			isMuted := c.muted
			c.mute.Unlock()

			if isMuted {
				klog.V(7).Infof("Mic is MUTED!\n")
				byData = make([]byte, len(byData))
			}

			byteCount, err := w.Write(byData)
			if err != nil {
				klog.V(1).Infof("w.Write failed. Err: %v\n", err)
				return err
			}
			klog.V(7).Infof("io.Writer succeeded. Bytes written: %d\n", byteCount)
		}
	}

	return nil
}

func (c *Client) Mute() {
	c.mute.Lock()
	c.muted = true
	c.mute.Unlock()
}

func (c *Client) Unmute() {
	c.mute.Lock()
	c.muted = false
	c.mute.Unlock()
}

func (c *Client) Stop() error {
	c.decoder = nil

	if c.file != nil {
		c.file.Close()
	}
	c.file = nil

	close(c.stopChan)
	<-c.stopChan

	return nil
}
