// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore dupl linter false positive
package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// StepService handles retrieving steps for builds
// from the server methods of the Vela API.
type StepService service

// Get returns the provided step.
func (svc *StepService) Get(ctx context.Context, org, repo string, build int64, step int32) (*api.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, build, step)

	// API Step type we want to return
	v := new(api.Step)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all steps.
func (svc *StepService) GetAll(ctx context.Context, org, repo string, build int64, opt *ListOptions) (*[]api.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps", org, repo, build)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Step type we want to return
	v := new([]api.Step)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// Add constructs a step with the provided details.
func (svc *StepService) Add(ctx context.Context, org, repo string, build int, s *api.Step) (*api.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps", org, repo, build)

	// API Step type we want to return
	v := new(api.Step)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, s, v)

	return v, resp, err
}

// Update modifies a step with the provided details.
func (svc *StepService) Update(ctx context.Context, org, repo string, build int64, s *api.Step) (*api.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, build, s.GetNumber())

	// API Step type we want to return
	v := new(api.Step)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided step.
func (svc *StepService) Remove(ctx context.Context, org, repo string, build, step int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, build, step)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "DELETE", u, nil, v)

	return v, resp, err
}
