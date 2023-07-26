// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package handler

import (
	"context"
	"fmt"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/nebula/v1/interfaces"
	sdkinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1/interfaces"
)

func NewHandler(options HandlerOptions) *Handler {
	handler := Handler{
		cache:        NewMessageCache(),
		nebulaClient: options.NebulaClient,
	}
	return &handler
}

func (h *Handler) InitializedConversation(im *sdkinterfaces.InitializationMessage) error {
	h.conversationID = im.Message.Data.ConversationID
	fmt.Printf("conversationID: %s\n", h.conversationID)
	return nil
}

func (h *Handler) RecognitionResultMessage(rr *sdkinterfaces.RecognitionResult) error {
	// No implementation required. Return Succeess!
	return nil
}

func (h *Handler) MessageResponseMessage(mr *sdkinterfaces.MessageResponse) error {
	for _, msg := range mr.Messages {
		fmt.Printf("\n\nMessage [%s]: %s\n\n", msg.From.Name, msg.Payload.Content)
		h.cache.Push(msg.ID, msg.Payload.Content, msg.From.ID, msg.From.Name)
	}
	return nil
}

func (h *Handler) InsightResponseMessage(ir *sdkinterfaces.InsightResponse) error {
	for _, insight := range ir.Insights {
		switch insight.Type {
		case sdkinterfaces.InsightTypeQuestion:
			err := h.HandleQuestion(&insight, ir.SequenceNumber)
			if err != nil {
				fmt.Printf("HandleQuestion failed. Err: %v\n", err)
				return err
			}
		case sdkinterfaces.InsightTypeFollowUp:
			err := h.HandleFollowUp(&insight, ir.SequenceNumber)
			if err != nil {
				fmt.Printf("HandleFollowUp failed. Err: %v\n", err)
				return err
			}
		case sdkinterfaces.InsightTypeActionItem:
			err := h.HandleActionItem(&insight, ir.SequenceNumber)
			if err != nil {
				fmt.Printf("HandleActionItem failed. Err: %v\n", err)
				return err
			}
		default:
			fmt.Printf("\n\n-------------------------------\n")
			fmt.Printf("Unknown InsightResponseMessage: %s\n\n", insight.Type)
			fmt.Printf("-------------------------------\n\n")
			return nil
		}
	}

	return nil
}

func (h *Handler) TopicResponseMessage(tr *sdkinterfaces.TopicResponse) error {
	conversation := h.cache.ReturnConversation()

	for _, curTopic := range tr.Topics {
		prompt := fmt.Sprintf("The topic of \"%s\" came up in this conversation I am having. Concisely summarize how this topic is relevant to this conversation.\n", curTopic.Phrases)

		request := interfaces.AskNebulaRequest{
			Prompt: interfaces.Prompt{
				Instruction: prompt,
				Conversation: interfaces.Conversation{
					Text: conversation,
				},
			},
		}

		nebulaResult, err := h.nebulaClient.AskNebula(context.Background(), request)
		if err != nil {
			fmt.Printf("AskNebula failed. Err: %v\n", err)
			return nil
		}

		fmt.Printf("\n\n-------------------------------\n")
		fmt.Printf("TOPIC:\n%s\n", prompt)
		fmt.Printf("\n\nNebula Response:\n%s\n", nebulaResult.Output.Text)
		fmt.Printf("-------------------------------\n\n")
	}

	return nil
}

func (h *Handler) TrackerResponseMessage(tr *sdkinterfaces.TrackerResponse) error {
	// No implementation required. Return Succeess!
	return nil
}

func (h *Handler) EntityResponseMessage(er *sdkinterfaces.EntityResponse) error {
	// No implementation required. Return Succeess!
	return nil
}

func (h *Handler) TeardownConversation(tm *sdkinterfaces.TeardownMessage) error {
	// No implementation required. Return Succeess!
	return nil
}

func (h *Handler) UserDefinedMessage(data []byte) error {
	// This is only needed on the client side and not on the plugin side.
	// No implementation required. Return Succeess!
	return nil
}

func (h *Handler) UnhandledMessage(byMsg []byte) error {
	fmt.Printf("\n\n-------------------------------\n")
	fmt.Printf("UnhandledMessage:\n%v\n", string(byMsg))
	fmt.Printf("-------------------------------\n\n")
	return ErrUnhandledMessage
}

func (h *Handler) HandleQuestion(insight *sdkinterfaces.Insight, number int) error {
	// No implementation required. Return Succeess!
	return nil
}

func (h *Handler) HandleActionItem(insight *sdkinterfaces.Insight, number int) error {
	// No implementation required. Return Succeess!
	return nil
}

func (h *Handler) HandleFollowUp(insight *sdkinterfaces.Insight, number int) error {
	// No implementation required. Return Succeess!
	return nil
}
