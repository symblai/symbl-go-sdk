// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import (
	"encoding/json"

	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1/interfaces"
)

type SymblMessageRouter struct {
	ConversationID string
	callback       interfaces.InsightCallback
}

func NewWithDefault() *SymblMessageRouter {
	return New(NewDefaultMessageRouter())
}

func New(callback interfaces.InsightCallback) *SymblMessageRouter {
	return &SymblMessageRouter{
		callback: callback,
	}
}

func (smr *SymblMessageRouter) GetConversationID() string {
	return smr.ConversationID
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
	// internal messages
	case MessageTypeInitListening:
		klog.V(2).Infof("Symbl Platform Initialized Listening\n")
	case MessageTypeInitConversation:
		return smr.InitializedConversation(byMsg)
	case MessageTypeInitRecognition:
		klog.V(2).Infof("Symbl Platform Initialized Recognition\n")
	case MessageTypeSessionModified:
		klog.V(2).Infof("Symbl Platform Session Modified\n")
	case MessageTypeTeardownConversation:
		klog.V(2).Infof("Symbl Platform Conversation Complete\n")
	case MessageTypeTeardownRecognition:
		klog.V(2).Infof("Symbl Platform Recognition Stopped\n")
	case MessageTypeError:
		return smr.HandleError(byMsg)
		// pass insights to the user
	case interfaces.MessageTypeRecognitionResult:
		return smr.RecognitionResultMessage(byMsg)
	case interfaces.MessageTypeMessageResponse:
		return smr.MessageResponseMessage(byMsg)
	case interfaces.MessageTypeInsightResponse:
		return smr.InsightResponseMessage(byMsg)
		// default handler
	default:
		klog.V(2).Infof("Unhandled Message Type: %s\n", smt.Message.Type)
		return smr.UnhandledMessage(byMsg)
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

	// TODO fix this....
	klog.V(6).Infof("\n\n%s\n\n", string(b))
	klog.V(6).Infof("HandleError LEAVE\n")
	return nil
}

func (smr *SymblMessageRouter) RecognitionResultMessage(byMsg []byte) error {
	klog.V(6).Info("RecognitionResultMessage ENTER\n")

	var rr interfaces.RecognitionResult
	err := json.Unmarshal(byMsg, &rr)
	if err != nil {
		klog.V(6).Infof("RecognitionResultMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("RecognitionResultMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.RecognitionResultMessage(&rr)
		if err != nil {
			klog.V(2).Infof("callback.RecognitionResultMessage failed. Err: %v\n", err)
		} else {
			klog.V(2).Infof("callback.RecognitionResultMessage succeeded\n")
		}
		klog.V(6).Infof("RecognitionResultMessage LEAVE\n")
		return err
	}

	klog.V(2).Infof("User callback is undefined\n")
	klog.V(6).Infof("RecognitionResultMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) MessageResponseMessage(byMsg []byte) error {
	klog.V(6).Info("MessageResponseMessage ENTER\n")

	var mr interfaces.MessageResponse
	err := json.Unmarshal(byMsg, &mr)
	if err != nil {
		klog.V(6).Infof("MessageResponseMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("MessageResponseMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.MessageResponseMessage(&mr)
		if err != nil {
			klog.V(2).Infof("callback.MessageResponseMessage failed. Err: %v\n", err)
		} else {
			klog.V(2).Infof("callback.MessageResponseMessage succeeded\n")
		}
		klog.V(6).Infof("MessageResponseMessage LEAVE\n")
		return err
	}

	klog.V(2).Infof("User callback is undefined\n")
	klog.V(6).Infof("MessageResponseMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) InsightResponseMessage(byMsg []byte) error {
	klog.V(6).Info("InsightResponseMessage ENTER\n")

	var ir interfaces.InsightResponse
	err := json.Unmarshal(byMsg, &ir)
	if err != nil {
		klog.V(6).Infof("InsightResponseMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("InsightResponseMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.InsightResponseMessage(&ir)
		if err != nil {
			klog.V(2).Infof("callback.InsightResponseMessage failed. Err: %v\n", err)
		} else {
			klog.V(2).Infof("callback.InsightResponseMessage succeeded\n")
		}
		klog.V(6).Infof("InsightResponseMessage LEAVE\n")
		return err
	}

	klog.V(2).Infof("User callback is undefined\n")
	klog.V(6).Infof("InsightResponseMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) UnhandledMessage(byMsg []byte) error {
	klog.V(6).Info("UnhandledMessage ENTER\n")

	if smr.callback != nil {
		err := smr.callback.UnhandledMessage(byMsg)
		if err != nil {
			klog.V(2).Infof("callback.UnhandledMessage failed. Err: %v\n", err)
		} else {
			klog.V(2).Infof("callback.UnhandledMessage succeeded\n")
		}
		klog.V(6).Infof("UnhandledMessage LEAVE\n")
		return err
	}

	klog.V(2).Infof("User callback is undefined\n")
	klog.V(6).Infof("UnhandledMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}
