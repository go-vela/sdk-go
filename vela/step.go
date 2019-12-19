// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// StepService handles retrieving steps for builds
// from the server methods of the Vela API.
type StepService service

// Get returns the provided step.
func (svc *StepService) Get(org, repo string, build, step int) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, build, step)

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all steps.
func (svc *StepService) GetAll(org, repo string, build int) (*[]library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps", org, repo, build)

	// slice library Step type we want to return
	v := new([]library.Step)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a step with the provided details.
func (svc *StepService) Add(org, repo string, build int, s *library.Step) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps", org, repo, build)

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := svc.client.Call("POST", u, s, v)

	return v, resp, err
}

// Update modifies a step with the provided details.
func (svc *StepService) Update(org, repo string, build int, s *library.Step) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, build, s.GetNumber())

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided step.
func (svc *StepService) Remove(org, repo string, build, step int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d", org, repo, build, step)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
