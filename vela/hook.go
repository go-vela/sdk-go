// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// HookService handles retrieving hooks from
// the server methods of the Vela API.
type HookService service

// Get returns the provided hook.
func (svc *HookService) Get(org, repo string, hook int64) (*api.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s/%d", org, repo, hook)

	// API Hook type we want to return
	v := new(api.Hook)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all hooks.
func (svc *HookService) GetAll(org, repo string, opt *ListOptions) (*[]api.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Hook type we want to return
	v := new([]api.Hook)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a hook with the provided details.
func (svc *HookService) Add(org, repo string, h *api.Hook) (*api.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s", org, repo)

	// API Hook type we want to return
	v := new(api.Hook)

	// send request using client
	resp, err := svc.client.Call("POST", u, h, v)

	return v, resp, err
}

// Update modifies a hook with the provided details.
func (svc *HookService) Update(org, repo string, h *api.Hook) (*api.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s/%d", org, repo, h.GetNumber())

	// API Hook type we want to return
	v := new(api.Hook)

	// send request using client
	resp, err := svc.client.Call("PUT", u, h, v)

	return v, resp, err
}

// Remove deletes the provided hook.
func (svc *HookService) Remove(org, repo string, hook int64) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s/%d", org, repo, hook)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
