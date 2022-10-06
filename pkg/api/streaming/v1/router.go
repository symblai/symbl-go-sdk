// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import (
	"encoding/json"
	"fmt"

	klog "k8s.io/klog/v2"
)

type SymblMessageRouter struct {
	ConversationID string
}

func New() *SymblMessageRouter {
	return &SymblMessageRouter{}
}

func (smr *SymblMessageRouter) Message(byMsg []byte) error {
	klog.V(6).Infof("SymblMessageRouter::Message ENTER\n")

	// TODO delete
	// klog.V(6).Infof("\n\n\n")
	// klog.V(6).Infof("IMPORTANT: Never print in production\n")
	// klog.V(6).Infof("SymblMessageRouter::Message byMsg: %v\n", byMsg)
	// klog.V(6).Infof("SymblMessageRouter::Message byMsg: %s\n", string(byMsg))
	// klog.V(6).Infof("\n\n\n")

	// what is the high level message here?
	var mt MessageType
	err := json.Unmarshal(byMsg, &mt)
	if err != nil {
		klog.V(6).Infof("SymblMessageRouter json.Unmarshal(MessageType) failed. Err: %v\n", err)
		klog.V(6).Infof("SymblMessageRouter LEAVE\n")
		return err
	}

	switch mt.Type {
	case MessageTypeError:
		return smr.HandleError(byMsg)
	case MessageTypeMessage:
		klog.V(6).Infof("Valid message\n")
	default:
		klog.Errorf("Invalid Type: %s\n", mt.Type)
		return ErrInvalidMessageType
	}

	// we know it's a valid message, what type of Symbl message is this?
	var smt SybmlMessageType
	err = json.Unmarshal(byMsg, &smt)
	if err != nil {
		klog.V(6).Infof("SymblMessageRouter json.Unmarshal(SybmlMessageType) failed. Err: %v\n", err)
		klog.V(6).Infof("SymblMessageRouter LEAVE\n")
		return err
	}

	switch smt.Message.Type {
	case MessageTypeInitListening:
		klog.V(2).Infof("Symbl Platform Initialized Listening\n")
	case MessageTypeInitConversation:
		return smr.InitializedConversation(byMsg)
	case MessageTypeInitRecognition:
		klog.V(2).Infof("Symbl Platform Initialized Recognition\n")
	case MessageTypeError:
		return smr.HandleError(byMsg)
	default:
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
	}

	klog.V(6).Infof("SymblMessageRouter Succeeded\n")
	klog.V(6).Infof("SymblMessageRouter LEAVE\n")
	return nil
}

func (smr *SymblMessageRouter) InitializedConversation(byMsg []byte) error {
	klog.V(6).Info("InitializedConversation ENTER\n")

	var symblInit SymblInitializationMessage
	err := json.Unmarshal(byMsg, &symblInit)
	if err != nil {
		klog.V(6).Infof("InitializedConversation json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("InitializedConversation LEAVE\n")
		return err
	}

	smr.ConversationID = symblInit.Message.Data.ConversationID

	klog.V(2).Infof("Setting Symbl ConversationID %s\n", smr.ConversationID)
	klog.V(6).Infof("InitializedConversation LEAVE\n")
	return nil
}

func (smr *SymblMessageRouter) HandleError(byMsg []byte) error {
	klog.V(6).Info("HandleError ENTER\n")

	var symbError SymblError
	err := json.Unmarshal(byMsg, &symbError)
	if err != nil {
		klog.V(6).Infof("HandleError json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("HandleError LEAVE\n")
		return err
	}

	b, err := json.MarshalIndent(symbError, "", "    ")
	if err != nil {
		klog.V(6).Infof("HandleError MarshalIndent failed. Err: %v\n", err)
		klog.V(6).Infof("HandleError LEAVE\n")
		return err
	}

	klog.V(6).Infof("\n\n%s\n\n", string(b))
	klog.V(6).Infof("HandleError LEAVE\n")
	return nil
}
