// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package handler

import (
	"container/list"
	"sync"

	nebula "github.com/dvonthenen/symbl-go-sdk/pkg/api/nebula/v1"
)

/*
	MessageCache
*/
type Author struct {
	Name  string
	Email string
}
type Message struct {
	ID     string
	Text   string
	Author Author
}

type MessageCache struct {
	rotatingWindowOfMsg *list.List
	mapIdToMsg          map[string]*Message
	mu                  sync.Mutex
}

/*
	Handler for messages
*/
type HandlerOptions struct {
	NebulaClient *nebula.Client
}

type Handler struct {
	// properties
	conversationID string

	// housekeeping
	cache        *MessageCache
	nebulaClient *nebula.Client
}
