// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
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

// PipelineOptions represents the optional parameters
// to the PipelineService.
type PipelineOptions struct {
	// Output of the pipeline being returned.
	//
	// Can be: json or yaml
	//
	// Default: yaml
	Output string `url:"output,omitempty"`

	// Reference of the pipeline from the repo.
	//
	// Typically would be a commit SHA but can also be a branch or tag.
	//
	// Default: master
	Ref string `url:"ref,omitempty"`

	// Enables expanding templates when validating a pipeline.
	//
	// Can be: true or false
	//
	// Default: true
	Template bool `url:"template,omitempty"`
}

// Get returns the provided pipeline.
//
// nolint: lll // ignore long line length due to variable names
func (svc *PipelineService) Get(org, repo string, opt *PipelineOptions) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v, nil)

	return v, resp, err
}

// Compile returns the provided fully compiled pipeline.
//
// nolint: lll // ignore long line length due to variable names
func (svc *PipelineService) Compile(org, repo string, opt *PipelineOptions) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/compile", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v, nil)

	return v, resp, err
}

// Expand returns the provided pipeline fully compiled.
//
// nolint: lll // ignore long line length due to variable names
func (svc *PipelineService) Expand(org, repo string, opt *PipelineOptions) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/expand", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v, nil)

	return v, resp, err
}

// Templates returns the provided templates for a pipeline.
//
// nolint: lll // ignore long line length due to variable names
func (svc *PipelineService) Templates(org, repo string, opt *PipelineOptions) (map[string]*yaml.Template, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/templates", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// yaml Templates type we want to return
	v := make(map[string]*yaml.Template)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v, nil)

	return v, resp, err
}

// Validate returns the validation status of the provided pipeline.
//
// nolint: lll // ignore long line length due to variable names
func (svc *PipelineService) Validate(org, repo string, opt *PipelineOptions) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/validate", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v, nil)

	return v, resp, err
}

func (svc *PipelineService) ValidateRaw(b64Pipeline string, opt *PipelineOptions) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/pipeline/raw"

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("POST", u, b64Pipeline, v, nil)

	return v, resp, err
}
