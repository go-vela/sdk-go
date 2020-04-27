// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// DeploymentService handles retrieving deployments from
// the server methods of the Vela API.
type DeploymentService service

// Get returns the provided deployment.
func (svc *DeploymentService) Get(org, repo string, deployment int) (*library.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/deployments/%s/%s/%d", org, repo, deployment)

	// library Deployment type we want to return
	v := new(library.Deployment)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all deployments.
func (svc *DeploymentService) GetAll(org, repo string, opt *ListOptions) (*[]library.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/deployments/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Deployment type we want to return
	v := new([]library.Deployment)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a deployment with the provided details.
func (svc *DeploymentService) Add(org, repo string, d *library.Deployment) (*library.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/deployments/%s/%s", org, repo)

	// library Deployment type we want to return
	v := new(library.Deployment)

	// send request using client
	resp, err := svc.client.Call("POST", u, d, v)

	return v, resp, err
}
