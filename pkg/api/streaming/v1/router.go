// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package streaming

import (
	"encoding/json"
	"errors"

	prettyjson "github.com/hokaccha/go-prettyjson"
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

	// what is the high level message here?
	var mt MessageType
	err := json.Unmarshal(byMsg, &mt)
	if err != nil {
		klog.V(1).Infof("SymblMessageRouter json.Unmarshal(MessageType) failed. Err: %v\n", err)
		klog.V(6).Infof("SymblMessageRouter LEAVE\n")
		return err
	}

	switch mt.Type {
	// errors
	case MessageTypeError:
		return smr.HandleError(byMsg)
	// platform messages
	case MessageTypeMessage:
		return smr.handlePlatformMessage(byMsg)
	// insights
	case interfaces.MessageTypeMessageResponse:
		return smr.MessageResponseMessage(byMsg)
	case interfaces.MessageTypeInsightResponse:
		return smr.InsightResponseMessage(byMsg)
	case interfaces.MessageTypeTopicResponse:
		return smr.TopicResponseMessage(byMsg)
	case interfaces.MessageTypeTrackerResponse:
		return smr.TrackerResponseMessage(byMsg)
	case interfaces.MessageTypeEntityResponse:
		return smr.EntityResponseMessage(byMsg)
	// user defined messages
	case interfaces.MessageTypeUserDefined:
		return smr.UserDefinedMessage(byMsg)
	// default
	default:
		return smr.UnhandledMessage(byMsg)
	}

	klog.V(3).Infof("SymblMessageRouter Succeeded\n")
	klog.V(6).Infof("SymblMessageRouter LEAVE\n")
	return nil
}

func (smr *SymblMessageRouter) handlePlatformMessage(byMsg []byte) error {
	klog.V(6).Infof("handlePlatformMessage ENTER\n")

	// we know it's a valid message, what type of Symbl message is this?
	var smt SybmlMessageType
	err := json.Unmarshal(byMsg, &smt)
	if err != nil {
		klog.V(1).Infof("json.Unmarshal(SybmlMessageType) failed. Err: %v\n", err)
		klog.V(6).Infof("handlePlatformMessage LEAVE\n")
		return err
	}

	switch smt.Message.Type {
	// internal messages
	case MessageTypeInitListening:
		smr.printDebugMessages("SymblMessageRouter.PlatformMessage", byMsg)
		klog.V(3).Infof("Symbl Platform Initialized Listening\n")
	case MessageTypeInitConversation:
		return smr.InitializedConversation(byMsg)
	case MessageTypeInitRecognition:
		smr.printDebugMessages("SymblMessageRouter.PlatformMessage", byMsg)
		klog.V(3).Infof("Symbl Platform Initialized Recognition\n")
	case MessageTypeSessionModified:
		smr.printDebugMessages("SymblMessageRouter.PlatformMessage", byMsg)
		klog.V(3).Infof("Symbl Platform Session Modified\n")
	case MessageTypeTeardownConversation:
		return smr.TeardownConversation(byMsg)
	case MessageTypeTeardownRecognition:
		smr.printDebugMessages("SymblMessageRouter.PlatformMessage", byMsg)
		klog.V(3).Infof("Symbl Platform Teardown Recognition\n")
	// transcription
	case interfaces.MessageTypeRecognitionResult:
		return smr.RecognitionResultMessage(byMsg)
	// error
	case MessageTypeError:
		return smr.HandleError(byMsg)
	// default handler
	default:
		klog.V(1).Infof("\n\nInvalid PlatformMessage Type: %s\n", smt.Message.Type)
		klog.V(1).Infof("%s\n\n", string(byMsg))
		return ErrInvalidMessageType
	}

	klog.V(3).Infof("handlePlatformMessage Succeeded\n")
	klog.V(6).Infof("handlePlatformMessage LEAVE\n")
	return nil
}

func (smr *SymblMessageRouter) InitializedConversation(byMsg []byte) error {
	klog.V(6).Info("InitializedConversation ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.InitializedConversation", byMsg)

	var im interfaces.InitializationMessage
	err := json.Unmarshal(byMsg, &im)
	if err != nil {
		klog.V(6).Infof("InitializedConversation json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("InitializedConversation LEAVE\n")
		return err
	}

	// save for the router
	smr.ConversationID = im.Message.Data.ConversationID

	// callback
	if smr.callback != nil {
		err := smr.callback.InitializedConversation(&im)
		if err != nil {
			klog.V(1).Infof("callback.InitializedConversation failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.InitializedConversation succeeded\n")
		}

		klog.V(6).Infof("InitializedConversation LEAVE\n")
		return err
	}

	klog.V(3).Infof("InitializedConversation: ConversationID %s\n", smr.ConversationID)
	klog.V(6).Infof("InitializedConversation LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) HandleError(byMsg []byte) error {
	klog.V(6).Info("HandleError ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.HandleError", byMsg)

	var symbError SymblError
	err := json.Unmarshal(byMsg, &symbError)
	if err != nil {
		klog.V(1).Infof("HandleError json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("HandleError LEAVE\n")
		return err
	}

	b, err := json.MarshalIndent(symbError, "", "    ")
	if err != nil {
		klog.V(1).Infof("HandleError MarshalIndent failed. Err: %v\n", err)
		klog.V(6).Infof("HandleError LEAVE\n")
		return err
	}

	klog.V(1).Infof("\n\nError: %s\n\n", string(b))
	klog.V(6).Infof("HandleError LEAVE\n")
	return errors.New(string(b))
}

func (smr *SymblMessageRouter) RecognitionResultMessage(byMsg []byte) error {
	klog.V(6).Info("RecognitionResultMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.RecognitionResultMessage", byMsg)

	var rr interfaces.RecognitionResult
	err := json.Unmarshal(byMsg, &rr)
	if err != nil {
		klog.V(1).Infof("RecognitionResultMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("RecognitionResultMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.RecognitionResultMessage(&rr)
		if err != nil {
			klog.V(1).Infof("callback.RecognitionResultMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.RecognitionResultMessage succeeded\n")
		}
		klog.V(6).Infof("RecognitionResultMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("RecognitionResultMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) MessageResponseMessage(byMsg []byte) error {
	klog.V(6).Info("MessageResponseMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.MessageResponseMessage", byMsg)

	var mr interfaces.MessageResponse
	err := json.Unmarshal(byMsg, &mr)
	if err != nil {
		klog.V(1).Infof("MessageResponseMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("MessageResponseMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.MessageResponseMessage(&mr)
		if err != nil {
			klog.V(1).Infof("callback.MessageResponseMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.MessageResponseMessage succeeded\n")
		}
		klog.V(6).Infof("MessageResponseMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("MessageResponseMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) InsightResponseMessage(byMsg []byte) error {
	klog.V(6).Info("InsightResponseMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.InsightResponseMessage", byMsg)

	var ir interfaces.InsightResponse
	err := json.Unmarshal(byMsg, &ir)
	if err != nil {
		klog.V(1).Infof("InsightResponseMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("InsightResponseMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.InsightResponseMessage(&ir)
		if err != nil {
			klog.V(1).Infof("callback.InsightResponseMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.InsightResponseMessage succeeded\n")
		}
		klog.V(6).Infof("InsightResponseMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("InsightResponseMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) TopicResponseMessage(byMsg []byte) error {
	klog.V(6).Info("TopicResponseMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.TopicResponseMessage", byMsg)

	var tr interfaces.TopicResponse
	err := json.Unmarshal(byMsg, &tr)
	if err != nil {
		klog.V(1).Infof("TopicResponseMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("TopicResponseMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.TopicResponseMessage(&tr)
		if err != nil {
			klog.V(1).Infof("callback.TopicResponseMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.TopicResponseMessage succeeded\n")
		}
		klog.V(6).Infof("TopicResponseMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("TopicResponseMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) TrackerResponseMessage(byMsg []byte) error {
	klog.V(6).Info("TrackerResponseMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.TrackerResponseMessage", byMsg)

	var tr interfaces.TrackerResponse
	err := json.Unmarshal(byMsg, &tr)
	if err != nil {
		klog.V(1).Infof("TrackerResponseMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("TrackerResponseMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.TrackerResponseMessage(&tr)
		if err != nil {
			klog.V(1).Infof("callback.TrackerResponseMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.TrackerResponseMessage succeeded\n")
		}
		klog.V(6).Infof("TrackerResponseMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("TrackerResponseMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) EntityResponseMessage(byMsg []byte) error {
	klog.V(6).Info("EntityResponseMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.EntityResponseMessage", byMsg)

	var er interfaces.EntityResponse
	err := json.Unmarshal(byMsg, &er)
	if err != nil {
		klog.V(1).Infof("EntityResponseMessage json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("EntityResponseMessage LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.EntityResponseMessage(&er)
		if err != nil {
			klog.V(1).Infof("callback.EntityResponseMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.EntityResponseMessage succeeded\n")
		}
		klog.V(6).Infof("EntityResponseMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("EntityResponseMessage LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) TeardownConversation(byMsg []byte) error {
	klog.V(6).Info("TeardownConversation ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.TeardownConversation", byMsg)

	var tm interfaces.TeardownMessage
	err := json.Unmarshal(byMsg, &tm)
	if err != nil {
		klog.V(6).Infof("TeardownConversation json.Unmarshal failed. Err: %v\n", err)
		klog.V(6).Infof("TeardownConversation LEAVE\n")
		return err
	}

	if smr.callback != nil {
		err := smr.callback.TeardownConversation(&tm)
		if err != nil {
			klog.V(1).Infof("callback.TeardownConversation failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.TeardownConversation succeeded\n")
		}

		klog.V(6).Infof("TeardownConversation LEAVE\n")
		return err
	}

	klog.V(6).Infof("TeardownConversation LEAVE\n")
	return ErrUserCallbackNotDefined
}

func (smr *SymblMessageRouter) UnhandledMessage(byMsg []byte) error {
	klog.V(6).Info("UnhandledMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.UnhandledMessage", byMsg)

	if smr.callback != nil {
		err := smr.callback.UnhandledMessage(byMsg)
		if err != nil {
			klog.V(1).Infof("callback.UnhandledMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.UnhandledMessage succeeded\n")
		}
		klog.V(6).Infof("UnhandledMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("UnhandledMessage LEAVE\n")
	return ErrInvalidMessageType
}

func (smr *SymblMessageRouter) UserDefinedMessage(byMsg []byte) error {
	klog.V(6).Info("UserDefinedMessage ENTER\n")

	// trace debugging
	smr.printDebugMessages("SymblMessageRouter.UserDefinedMessage", byMsg)

	if smr.callback != nil {
		err := smr.callback.UserDefinedMessage(byMsg)
		if err != nil {
			klog.V(1).Infof("callback.UserDefinedMessage failed. Err: %v\n", err)
		} else {
			klog.V(3).Infof("callback.UserDefinedMessage succeeded\n")
		}
		klog.V(6).Infof("UserDefinedMessage LEAVE\n")
		return err
	}

	klog.V(1).Infof("User callback is undefined\n")
	klog.V(6).Infof("UserDefinedMessage LEAVE\n")
	return ErrInvalidMessageType
}

func (smr *SymblMessageRouter) printDebugMessages(function string, byMsg []byte) {
	prettyJson, err := prettyjson.Format(byMsg)
	if err != nil {
		klog.V(1).Infof("prettyjson.Marshal failed. Err: %v\n", err)
	}

	klog.V(6).Infof("\n\n-----------------------------------------------\n")
	klog.V(6).Infof("%s RAW:\n%s\n", function, prettyJson)
	klog.V(6).Infof("-----------------------------------------------\n\n\n")
}
