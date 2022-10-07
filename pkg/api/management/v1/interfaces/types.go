// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package interfaces

import "time"

/*
	Definitions
*/
type Tracker struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Categories  []string  `json:"categories" validate:"required"`
	Languages   []string  `json:"languages" validate:"required"`
	Vocabulary  []string  `json:"vocabulary" validate:"required"`
	CreatedOn   time.Time `json:"createdOn"`
	UpdatedOn   time.Time `json:"updatedOn"`
}

/*
	Input parameters for Management API calls
*/
// TrackerRequest
type TrackerRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Categories  []string `json:"categories" validate:"required"`
	Languages   []string `json:"languages" validate:"required"`
	Vocabulary  []string `json:"vocabulary" validate:"required"`
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
