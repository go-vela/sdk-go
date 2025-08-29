// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// ScheduleService handles retrieving schedules from the server methods of the Vela API.
type ScheduleService service

// Get returns the provided schedule from the repo.
func (svc *ScheduleService) Get(ctx context.Context, org, repo, schedule string) (*api.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s/%s", org, repo, schedule)

	// API Schedule type we want to return
	v := new(api.Schedule)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all schedules from the repo.
func (svc *ScheduleService) GetAll(ctx context.Context, org, repo string, opt *ListOptions) (*[]api.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Schedule type we want to return
	v := new([]api.Schedule)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// Add constructs a schedule with the provided details.
func (svc *ScheduleService) Add(ctx context.Context, org, repo string, s *api.Schedule) (*api.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s", org, repo)

	// API Schedule type we want to return
	v := new(api.Schedule)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, s, v)

	return v, resp, err
}

// Update modifies a schedule with the provided details.
func (svc *ScheduleService) Update(ctx context.Context, org, repo string, s *api.Schedule) (*api.Schedule, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s/%s", org, repo, s.GetName())

	// API Schedule type we want to return
	v := new(api.Schedule)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided schedule.
func (svc *ScheduleService) Remove(ctx context.Context, org, repo, schedule string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/schedules/%s/%s/%s", org, repo, schedule)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "DELETE", u, nil, v)

	return v, resp, err
}
