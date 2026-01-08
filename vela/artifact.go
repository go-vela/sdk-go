// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// ArtifactService handles retrieving artifacts from
// the server methods of the Vela API.
type ArtifactService service

// Add constructs an artifact with the provided details.
func (svc *ArtifactService) Add(ctx context.Context, org, repo string, build int64) (*api.Artifact, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/artifacts", org, repo, build)

	// API Artifact type we want to return
	a := new(api.Artifact)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, nil, a)

	return a, resp, err
}

// Update modifies an artifact with the provided details.
func (svc *ArtifactService) Update(ctx context.Context, org, repo string, build int64, a *api.Artifact) (*api.Artifact, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/artifacts", org, repo, build)

	// API Artifact type we want to return
	_a := new(api.Artifact)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, a, _a)

	return _a, resp, err
}
