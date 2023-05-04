// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/server/api/types"
)

// ScheduleService handles retrieving schedules from the server methods of the Vela API.
type ScheduleService service

// Get returns the provided schedule from the repo.
func (svc *ScheduleService) Get(org, repo, schedule string) (*types.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s/%s", org, repo, schedule)

	// library Schedule type we want to return
	v := new(types.Schedule)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all schedules from the repo.
func (svc *ScheduleService) GetAll(org, repo string, opt *ListOptions) (*[]types.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Schedule type we want to return
	v := new([]types.Schedule)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a schedule with the provided details.
func (svc *ScheduleService) Add(org, repo string, s *types.Schedule) (*types.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s", org, repo)

	// library Schedule type we want to return
	v := new(types.Schedule)

	// send request using client
	resp, err := svc.client.Call("POST", u, s, v)

	return v, resp, err
}

// Update modifies a schedule with the provided details.
func (svc *ScheduleService) Update(org, repo string, s *types.Schedule) (*types.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s/%s", org, repo, s.GetName())

	// library Schedule type we want to return
	v := new(types.Schedule)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided schedule.
func (svc *ScheduleService) Remove(org, repo, schedule string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s/%s", org, repo, schedule)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
