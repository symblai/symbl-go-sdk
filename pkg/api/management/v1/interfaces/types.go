// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import "time"

/*
	Shared definitions
*/
type Tracker struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description,omitempty"`
	Categories  []string  `json:"categories" validate:"required"`
	Languages   []string  `json:"languages" validate:"required"`
	Vocabulary  []string  `json:"vocabulary" validate:"required"`
	CreatedOn   time.Time `json:"createdOn,omitempty"`
	UpdatedOn   time.Time `json:"updatedOn,omitempty"`
}

type Entity struct {
	ID       string   `json:"id,omitempty"`
	Type     string   `json:"type" validate:"required"`
	SubType  string   `json:"subType" validate:"required"`
	Category string   `json:"category" validate:"required"`
	Values   []string `json:"values" validate:"required"`
}

type Group struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Criteria    string `json:"criteria,omitempty" validate:"required"`
}

/*
	Input parameters for Management API calls
*/
// TrackerRequest
type TrackerRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description,omitempty"`
	Categories  []string `json:"categories" validate:"required"`
	Languages   []string `json:"languages" validate:"required"`
	Vocabulary  []string `json:"vocabulary" validate:"required"`
}

// EntityRequest minus the ID
type EntityRequest struct {
	Type     string   `json:"type" validate:"required"`
	SubType  string   `json:"subType" validate:"required"`
	Category string   `json:"category" validate:"required"`
	Values   []string `json:"values" validate:"required"`
}

// CreateEntityRequest the request
type CreateEntityRequest struct {
	EntityArray []EntityRequest
}

// TrackerTupleRequest to modify a tracker
type TrackerTupleRequest struct {
	Op    string `json:"op" validate:"required"`
	Path  string `json:"path" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateTrackerRequest container for TrackerTupleRequest requests
type UpdateTrackerRequest struct {
	TrackerArray []TrackerTupleRequest
}

/*
	Output structs for Management API calls
*/
// TrackersResponse lists Trackers
type TrackersResponse struct {
	Trackers []Tracker `json:"trackers"`
}

// TrackerResponse result for an individual tracker
type TrackerResponse struct {
	Tracker Tracker `json:"tracker"`
}

// EntitiesResponse list of Entities
type EntitiesResponse struct {
	Entities []Entity `json:"entities"`
}

// EntitiesResponse list of Entities
type EntityResponse struct {
	Entity Entity `json:"entity"`
}

// ConversationGroupResponse when create
type ConversationGroupResponse struct {
	Group Group `json:"group"`
}

// ConversationGroupsResponse
type ConversationGroupsResponse struct {
	Groups []Group `json:"groups"`
}
