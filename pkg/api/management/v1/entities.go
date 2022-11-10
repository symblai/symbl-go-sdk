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

func (m *Management) GetEntites(ctx context.Context) (*interfaces.EntitiesResponse, error) {
	klog.V(6).Infof("mgmt.GetEntites ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := version.GetManagementAPI(version.ManagementEntitiesURI)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.GetEntites LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.EntitiesResponse

	err = m.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.GetEntites LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Management Entities succeeded\n")
	klog.V(6).Infof("mgmt.GetEntites LEAVE\n")
	return &result, nil
}

func (m *Management) GetEntitById(ctx context.Context, entityId string) (*interfaces.Entity, error) {
	klog.V(6).Infof("mgmt.GetEntitById ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := version.GetManagementAPI(version.ManagementEntitiesByIdURI, entityId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.GetEntitById LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.Entity

	err = m.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.GetEntitById LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Management Entity succeeded\n")
	klog.V(6).Infof("mgmt.GetEntitById LEAVE\n")
	return &result, nil
}

/*
	TODO: create doesn't return Entity object that's populated
*/
func (m *Management) CreateEntity(ctx context.Context, request interfaces.EntityRequest) (*interfaces.Entity, error) {
	klog.V(6).Infof("mgmt.CreateEntity ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("CreateEntity validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("mgmt.CreateEntity LEAVE\n")
		return nil, err
	}

	// request
	URI := version.GetManagementAPI(version.ManagementEntitiesURI)
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.CreateEntity LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.CreateEntity LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.Entity

	err = m.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.CreateEntity LEAVE\n")
			return nil, err
		}
	}

	klog.V(3).Infof("GET Create Entity succeeded\n")
	klog.V(6).Infof("mgmt.CreateEntity LEAVE\n")
	return &result, nil
}

/*
	TODO: This appears broken... This is the error we get back:

	HTTP Code: 400
	{
		"message":"Entity request body must contain 'type'. Please refer Symbl docs for more information."
	}

	Which contradicts the documentation: https://docs.symbl.ai/docs/update-entities
*/
// func (m *Management) UpdateEntity(ctx context.Context, request interfaces.Entity) (*interfaces.Entity, error) {
// 	klog.V(6).Infof("mgmt.UpdateEntity ENTER\n")

// 	// checks
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	// validate input
// 	v := validator.New()
// 	err := v.Struct(request)
// 	if err != nil {
// 		for _, e := range err.(validator.ValidationErrors) {
// 			klog.V(1).Infof("UpdateEntity validation failed. Err: %v\n", e)
// 		}
// 		klog.V(6).Infof("mgmt.UpdateEntity LEAVE\n")
// 		return nil, err
// 	}

// 	// request
// 	URI := version.GetManagementAPI(version.ManagementEntitiesByIdURI, request.ID)
// 	klog.V(6).Infof("Calling %s\n", URI)

// 	req, err := http.NewRequestWithContext(ctx, "PUT", URI, nil)
// 	if err != nil {
// 		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
// 		klog.V(6).Infof("mgmt.UpdateEntity LEAVE\n")
// 		return nil, err
// 	}

// 	// check the status
// 	var result interfaces.Entity

// 	err = m.Client.Do(ctx, req, &result)

// 	if e, ok := err.(*symbl.StatusError); ok {
// 		if e.Resp.StatusCode != http.StatusOK {
// 			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
// 			klog.V(6).Infof("mgmt.UpdateEntity LEAVE\n")
// 			return nil, err
// 		}
// 	}

// 	klog.V(3).Infof("GET Update Entity succeeded\n"))
// 	klog.V(6).Infof("mgmt.UpdateEntity LEAVE\n")
// 	return &result, nil
// }

func (m *Management) DeleteEntity(ctx context.Context, entityId string) error {
	klog.V(6).Infof("mgmt.DeleteEntity ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	if entityId == "" {
		klog.V(1).Infof("entityId is empty\n")
		klog.V(6).Infof("mgmt.DeleteEntity LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.ManagementEntitiesByIdURI, entityId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "DELETE", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.DeleteEntity LEAVE\n")
		return err
	}

	// check the status
	err = m.Client.Do(ctx, req, nil)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.DeleteEntity LEAVE\n")
			return e
		}
	}

	klog.V(3).Infof("GET Delete Entity succeeded\n")
	klog.V(6).Infof("mgmt.DeleteEntity LEAVE\n")
	return nil
}

func (m *Management) DeleteEntityBySubType(ctx context.Context, subType string) error {
	klog.V(6).Infof("mgmt.DeleteEntityBySubType ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	if subType == "" {
		klog.V(1).Infof("subType is empty\n")
		klog.V(6).Infof("mgmt.DeleteEntityBySubType LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.ManagementEntitiesBySubTypeURI, subType)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "DELETE", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.DeleteEntityBySubType LEAVE\n")
		return err
	}

	// check the status
	err = m.Client.Do(ctx, req, nil)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("mgmt.DeleteEntityBySubType LEAVE\n")
			return err
		}
	}

	klog.V(3).Infof("GET Delete EntityBySubType succeeded\n")
	klog.V(6).Infof("mgmt.DeleteEntityBySubType LEAVE\n")
	return nil
}
