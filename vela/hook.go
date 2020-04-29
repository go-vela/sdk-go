// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// HookService handles retrieving hooks from
// the server methods of the Vela API.
type HookService service

// Get returns the provided hook.
func (svc *HookService) Get(org, repo string, hook int) (*library.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s/%d", org, repo, hook)

	// library Hook type we want to return
	v := new(library.Hook)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all hooks.
func (svc *HookService) GetAll(org, repo string, opt *ListOptions) (*[]library.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Hook type we want to return
	v := new([]library.Hook)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a hook with the provided details.
func (svc *HookService) Add(org, repo string, h *library.Hook) (*library.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s", org, repo)

	// library Hook type we want to return
	v := new(library.Hook)

	// send request using client
	resp, err := svc.client.Call("POST", u, h, v)

	return v, resp, err
}

// Update modifies a hook with the provided details.
func (svc *HookService) Update(org, repo string, h *library.Hook) (*library.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s/%d", org, repo, h.GetNumber())

	// library Hook type we want to return
	v := new(library.Hook)

	// send request using client
	resp, err := svc.client.Call("PUT", u, h, v)

	return v, resp, err
}

// Remove deletes the provided hook.
func (svc *HookService) Remove(org, repo string, hook int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/hooks/%s/%s/%d", org, repo, hook)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
