// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package management

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	mgmtinterfaces "github.com/symblai/symbl-go-sdk/pkg/api/management/v1/interfaces"
	version "github.com/symblai/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
)

func (m *Management) GetTrackers(ctx context.Context) (*mgmtinterfaces.TrackersResponse, error) {
	klog.V(6).Infof("mgmt.GetTrackers ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// request
	URI := version.GetManagementAPI(version.ManagementTrackerURI)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.GetTrackers LEAVE\n")
		return nil, err
	}

	// check the status
	var result mgmtinterfaces.TrackersResponse

	err = m.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetTrackers LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetTrackers LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET Management Trackers succeeded\n")
	klog.V(6).Infof("mgmt.GetTrackers LEAVE\n")
	return &result, nil
}

func (m *Management) CreateTracker(ctx context.Context, request mgmtinterfaces.TrackerRequest) (*mgmtinterfaces.TrackerResponse, error) {
	klog.V(6).Infof("mgmt.CreateTracker ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("CreateTracker validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("mgmt.CreateTracker LEAVE\n")
		return nil, err
	}

	// request
	URI := version.GetManagementAPI(version.ManagementTrackerURI)
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.CreateTracker LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.CreateTracker LEAVE\n")
		return nil, err
	}

	// check the status
	var result mgmtinterfaces.TrackerResponse

	err = m.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.CreateTracker LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.CreateTracker LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET Create Trackers succeeded\n")
	klog.V(6).Infof("mgmt.CreateTracker LEAVE\n")
	return &result, nil
}

func (m *Management) UpdateTracker(ctx context.Context, trackerId string, request mgmtinterfaces.UpdateTrackerRequest) (*mgmtinterfaces.TrackerResponse, error) {
	klog.V(6).Infof("mgmt.UpdateTracker ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("UpdateTracker validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("mgmt.UpdateTracker LEAVE\n")
		return nil, err
	}

	// request
	URI := version.GetManagementAPI(version.ManagementTrackerByIdURI, trackerId)
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request.TrackerArray)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.UpdateTracker LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.UpdateTracker LEAVE\n")
		return nil, err
	}

	// check the status
	var result mgmtinterfaces.TrackerResponse

	err = m.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.UpdateTracker LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.UpdateTracker LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("PATCH UpdateTracker succeeded\n")
	klog.V(6).Infof("mgmt.UpdateTracker LEAVE\n")
	return &result, nil
}

func (m *Management) DeleteTracker(ctx context.Context, trackerId string) error {
	klog.V(6).Infof("mgmt.DeleteTracker ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	if trackerId == "" {
		klog.V(1).Infof("trackerId is empty\n")
		klog.V(6).Infof("mgmt.DeleteTracker LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.ManagementTrackerByIdURI, trackerId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "DELETE", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("mgmt.DeleteTracker LEAVE\n")
		return err
	}

	// check the status
	err = m.Client.Do(ctx, req, nil)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.DeleteTracker LEAVE\n")
				return err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.DeleteTracker LEAVE\n")
		return err
	}

	klog.V(3).Infof("GET Delete Trackers succeeded\n")
	klog.V(6).Infof("mgmt.DeleteTracker LEAVE\n")
	return nil
}
