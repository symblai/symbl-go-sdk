// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package handler

import (
	"container/list"
	"fmt"
)

func NewMessageCache() *MessageCache {
	cache := MessageCache{
		rotatingWindowOfMsg: list.New(),
		mapIdToMsg:          make(map[string]*Message),
	}
	return &cache
}

func (mc *MessageCache) Push(msgId, text, name, email string) error {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	// if len(mc.mapIdToMsg) >= DefaultNumOfMsgToCache {
	// 	e := mc.rotatingWindowOfMsg.Front()
	// 	if e != nil {
	// 		itemMessage := Message(e.Value.(Message))
	// 		delete(mc.mapIdToMsg, itemMessage.ID)
	// 		mc.rotatingWindowOfMsg.Remove(e)
	// 	}
	// }

	message := Message{
		ID:   msgId,
		Text: text,
		Author: Author{
			Name:  name,
			Email: email,
		},
	}
	mc.mapIdToMsg[msgId] = &message
	mc.rotatingWindowOfMsg.PushBack(message)

	return nil
}

func (mc *MessageCache) Find(msgId string) (*Message, error) {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	message := mc.mapIdToMsg[msgId]

	if message == nil {
		return nil, ErrItemNotFound
	}

	return message, nil
}

func (mc *MessageCache) ReturnConversation() string {
	conversation := ""

	for e := mc.rotatingWindowOfMsg.Front(); e != nil; e = e.Next() {
		msg := Message(e.Value.(Message))
		// fmt.Sprintf("msg: %v\n", msg)
		conversation += fmt.Sprintf("%s: %s\n", msg.Author.Name, msg.Text)
	}

	return conversation
}
