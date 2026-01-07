// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore dupl linter false positive
package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// SvcService handles retrieving services for builds
// from the server methods of the Vela API.
type SvcService service

// Get returns the provided service.
func (svc *SvcService) Get(ctx context.Context, org, repo string, build int64, service int32) (*api.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, build, service)

	// API Service type we want to return
	v := new(api.Service)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all services.
func (svc *SvcService) GetAll(ctx context.Context, org, repo string, build int64, opt *ListOptions) (*[]api.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services", org, repo, build)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Service type we want to return
	v := new([]api.Service)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// Add constructs a service with the provided details.
func (svc *SvcService) Add(ctx context.Context, org, repo string, build int, s *api.Service) (*api.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services", org, repo, build)

	// API Service type we want to return
	v := new(api.Service)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, s, v)

	return v, resp, err
}

// Update modifies a service with the provided details.
func (svc *SvcService) Update(ctx context.Context, org, repo string, build int64, s *api.Service) (*api.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, build, s.GetNumber())

	// API Service type we want to return
	v := new(api.Service)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided service.
func (svc *SvcService) Remove(ctx context.Context, org, repo string, build, service int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, build, service)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "DELETE", u, nil, v)

	return v, resp, err
}
