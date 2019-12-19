// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// BuildService handles retrieving builds from
// the server methods of the Vela API.
type BuildService service

// Get returns the provided build.
func (svc *BuildService) Get(org, repo string, build int) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, build)

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetAll returns a list of all builds.
func (svc *BuildService) GetAll(org, repo string) (*[]library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds", org, repo)

	// slice library Build type we want to return
	v := new([]library.Build)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetLogs returns the provided build logs.
func (svc *BuildService) GetLogs(org, repo string, build int) (*[]library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/logs", org, repo, build)

	// slice database Log type we want to return
	v := new([]library.Log)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Add constructs a build with the provided details.
func (svc *BuildService) Add(org, repo string, b *library.Build) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds", org, repo)

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, b, v)
	return v, resp, err
}

// Update modifies a build with the provided details.
func (svc *BuildService) Update(org, repo string, b *library.Build) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, b.GetNumber())

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := svc.client.Call("PUT", u, b, v)
	return v, resp, err
}

// Remove deletes the provided build.
func (svc *BuildService) Remove(org, repo string, build int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, build)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// Restart takes the build provided and restarts it
func (svc *BuildService) Restart(org, repo string, build int) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, build)

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)
	return v, resp, err
}
