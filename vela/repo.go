// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// RepoService handles retrieving repos from
// the server methods of the Vela API.
type RepoService service

// Get returns the provided repo.
func (svc *RepoService) Get(ctx context.Context, org, repo string) (*api.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s", org, repo)

	// API Repo type we want to return
	v := new(api.Repo)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all repos.
func (svc *RepoService) GetAll(ctx context.Context, opt *ListOptions) (*[]api.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/repos"

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Repo type we want to return
	v := new([]api.Repo)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// Add constructs a repo with the provided details.
func (svc *RepoService) Add(ctx context.Context, r *api.Repo) (*api.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/repos"

	// API Repo type we want to return
	v := new(api.Repo)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, r, v)

	return v, resp, err
}

// Update modifies a repo with the provided details.
func (svc *RepoService) Update(ctx context.Context, org, repo string, r *api.Repo) (*api.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s", org, repo)

	// API Repo type we want to return
	v := new(api.Repo)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, r, v)

	return v, resp, err
}

// Remove deletes the provided repo.
func (svc *RepoService) Remove(ctx context.Context, org, repo string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s", org, repo)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "DELETE", u, nil, v)

	return v, resp, err
}

// Repair modifies a damaged repo webhook.
func (svc *RepoService) Repair(ctx context.Context, org, repo string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/repair", org, repo)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "PATCH", u, nil, v)

	return v, resp, err
}

// Chown modifies the org of a repo.
func (svc *RepoService) Chown(ctx context.Context, org, repo string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/chown", org, repo)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "PATCH", u, nil, v)

	return v, resp, err
}
