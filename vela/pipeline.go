// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/yaml"
)

// PipelineService handles retrieving pipelines from
// the server methods of the Vela API.
type PipelineService service

// Get returns the provided pipeline.
func (svc *PipelineService) Get(org, repo string) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s", org, repo)

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Compile returns the provided fully compiled pipeline.
func (svc *PipelineService) Compile(org, repo string) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/compile", org, repo)

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

	return v, resp, err
}

// Expand returns the provided pipeline fully compiled.
func (svc *PipelineService) Expand(org, repo string) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/expand", org, repo)

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

	return v, resp, err
}

// Templates returns the provided templates for a pipeline.
func (svc *PipelineService) Templates(org, repo string) (map[string]*yaml.Template, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/templates", org, repo)

	// yaml Templates type we want to return
	v := make(map[string]*yaml.Template)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Validate returns the validation status of the provided pipeline.
func (svc *PipelineService) Validate(org, repo string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/validate", org, repo)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

	return v, resp, err
}
