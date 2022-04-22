// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
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

	// Enables expanding templates when validating a pipeline.
	//
	// Can be: true or false
	//
	// Default: true
	Template bool `url:"template,omitempty"`
}

// Get returns the provided pipeline.
func (svc *PipelineService) Get(org, repo, ref string, opt *PipelineOptions) (*library.Pipeline, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/%s", org, repo, ref)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// library Pipeline type we want to return
	v := new(library.Pipeline)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all pipelines.
func (svc *PipelineService) GetAll(org, repo string, opt *PipelineOptions) (*[]library.Pipeline, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Pipeline type we want to return
	v := new([]library.Pipeline)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a pipeline with the provided details.
func (svc *PipelineService) Add(org, repo string, h *library.Pipeline) (*library.Pipeline, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s", org, repo)

	// library Pipeline type we want to return
	v := new(library.Pipeline)

	// send request using client
	resp, err := svc.client.Call("POST", u, h, v)

	return v, resp, err
}

// Update modifies a pipeline with the provided details.
func (svc *PipelineService) Update(org, repo string, p *library.Pipeline) (*library.Pipeline, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/%s", org, repo, p.GetCommit())

	// library Pipeline type we want to return
	v := new(library.Pipeline)

	// send request using client
	resp, err := svc.client.Call("PUT", u, p, v)

	return v, resp, err
}

// Remove deletes the provided pipeline.
func (svc *PipelineService) Remove(org, repo string, pipeline string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/%s", org, repo, pipeline)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}

// Compile returns the provided fully compiled pipeline.
func (svc *PipelineService) Compile(org, repo, ref string, opt *PipelineOptions) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/%s/compile", org, repo, ref)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

	return v, resp, err
}

// Expand returns the provided pipeline fully compiled.
func (svc *PipelineService) Expand(org, repo, ref string, opt *PipelineOptions) (*yaml.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/%s/expand", org, repo, ref)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// yaml Build type we want to return
	v := new(yaml.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

	return v, resp, err
}

// Templates returns the provided templates for a pipeline.
func (svc *PipelineService) Templates(org, repo, ref string, opt *PipelineOptions) (map[string]*yaml.Template, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/%s/templates", org, repo, ref)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// yaml Templates type we want to return
	v := make(map[string]*yaml.Template)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Validate returns the validation status of the provided pipeline.
func (svc *PipelineService) Validate(org, repo, ref string, opt *PipelineOptions) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/pipelines/%s/%s/%s/validate", org, repo, ref)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

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
	resp, err := svc.client.Call("POST", u, b64Pipeline, v)

	return v, resp, err
}
