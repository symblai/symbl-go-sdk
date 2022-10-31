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

type EntityRequest struct {
	Type     string   `json:"type" validate:"required"`
	SubType  string   `json:"subType" validate:"required"`
	Category string   `json:"category" validate:"required"`
	Values   []string `json:"values" validate:"required"`
}

// ModifyTrackerRequest to modify a tracker
/*
type ModifyTrackerRequest struct {
	TrackerId string `validate:"required"`
	Op        string `json:"op" validate:"required"`
	Path      string `json:"path" validate:"required"`
	Value     string `json:"value" validate:"required"`
}
*/

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
