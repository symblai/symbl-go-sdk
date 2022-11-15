// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package management

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/management/v1/interfaces"
	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"
)

func (m *Management) GetConversationGroups(ctx context.Context) (*interfaces.ConversationGroupsResponse, error) {
	klog.V(6).Infof("mgmt.GetConversationGroups ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := version.GetManagementAPI(version.ManagementConversationGroupsURI)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.GetConversationGroups LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.ConversationGroupsResponse

	err = m.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.GetConversationGroups LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET ConversationGroups succeeded\n")
	klog.V(6).Infof("mgmt.GetConversationGroups LEAVE\n")
	return &result, nil
}

func (m *Management) GetConversationGroupById(ctx context.Context, conversationGroupId string) (*interfaces.ConversationGroupResponse, error) {
	klog.V(6).Infof("mgmt.GetConversationGroupById ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := version.GetManagementAPI(version.ManagementConversationGroupByIdURI, conversationGroupId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.GetConversationGroupById LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.ConversationGroupResponse

	err = m.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.GetConversationGroupById LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET ConversationGroupById succeeded\n")
	klog.V(6).Infof("mgmt.GetConversationGroupById LEAVE\n")
	return &result, nil
}

func (m *Management) CreateConversationGroup(ctx context.Context, request interfaces.Group) (*interfaces.ConversationGroupResponse, error) {
	klog.V(6).Infof("mgmt.CreateConversationGroup ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("CreateConversationGroup validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("mgmt.CreateConversationGroup LEAVE\n")
		return nil, err
	}

	// request
	URI := version.GetManagementAPI(version.ManagementConversationGroupURI)
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.CreateConversationGroup LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.CreateConversationGroup LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.ConversationGroupResponse

	err = m.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.CreateConversationGroup LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("POST CreateConversationGroup succeeded\n")
	klog.V(6).Infof("mgmt.CreateConversationGroup LEAVE\n")
	return &result, nil
}

func (m *Management) UpdateConversationGroup(ctx context.Context, request interfaces.Group) (*interfaces.ConversationGroupResponse, error) {
	klog.V(6).Infof("mgmt.UpdateConversationGroup ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	if request.ID == "" {
		klog.V(1).Infof("group.ID is empty\n")
		klog.V(6).Infof("async.UpdateConversationGroup LEAVE\n")
		return nil, ErrInvalidInput
	}

	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("UpdateConversationGroup validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("mgmt.UpdateConversationGroup LEAVE\n")
		return nil, err
	}

	// request
	URI := version.GetManagementAPI(version.ManagementConversationGroupByIdURI, request.ID)
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.CreateConversationGroup LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.UpdateConversationGroup LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.ConversationGroupResponse

	err = m.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.UpdateConversationGroup LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("PUT UpdateConversationGroup succeeded\n")
	klog.V(6).Infof("mgmt.UpdateConversationGroup LEAVE\n")
	return &result, nil
}

func (m *Management) DeleteConversationGroup(ctx context.Context, conversationGroupId string) error {
	klog.V(6).Infof("mgmt.DeleteConversationGroup ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	if conversationGroupId == "" {
		klog.V(1).Infof("entityId is empty\n")
		klog.V(6).Infof("mgmt.DeleteConversationGroup LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.ManagementConversationGroupByIdURI, conversationGroupId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "DELETE", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.DeleteConversationGroup LEAVE\n")
		return err
	}

	// check the status
	err = m.Client.Do(ctx, req, nil)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.DeleteConversationGroup LEAVE\n")
			return e
		}
	}

	klog.V(3).Infof("DELETE ConversationGroup succeeded\n")
	klog.V(6).Infof("mgmt.DeleteConversationGroup LEAVE\n")
	return nil
}
