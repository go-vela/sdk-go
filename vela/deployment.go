// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// DeploymentService handles retrieving deployments from
// the server methods of the Vela API.
type DeploymentService service

// Get returns the provided deployment.
func (svc *DeploymentService) Get(ctx context.Context, org, repo string, deployment int64) (*api.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/deployments/%s/%s/%d", org, repo, deployment)

	// API Deployment type we want to return
	v := new(api.Deployment)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all deployments.
func (svc *DeploymentService) GetAll(ctx context.Context, org, repo string, opt *ListOptions) (*[]api.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/deployments/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Deployment type we want to return
	v := new([]api.Deployment)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// Add constructs a deployment with the provided details.
func (svc *DeploymentService) Add(ctx context.Context, org, repo string, d *api.Deployment) (*api.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/deployments/%s/%s", org, repo)

	// API Deployment type we want to return
	v := new(api.Deployment)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, d, v)

	return v, resp, err
}
