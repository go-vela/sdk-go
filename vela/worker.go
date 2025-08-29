// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// WorkerService handles retrieving workers from
// the server methods of the Vela API.
type WorkerService service

// WorkerListOptions specifies the optional parameters to the
// Worker.GetAll method.
type WorkerListOptions struct {
	Active          string `url:"active,omitempty"`
	CheckedInBefore int64  `url:"checked_in_before,omitempty"`
	CheckedInAfter  int64  `url:"checked_in_after,omitempty"`
}

// Get returns the provided worker.
func (svc *WorkerService) Get(hostname string) (*api.Worker, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/workers/%s", hostname)

	// api Worker type we want to return
	v := new(api.Worker)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all workers.
func (svc *WorkerService) GetAll(opt *WorkerListOptions) (*[]api.Worker, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/workers"

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Worker type we want to return
	v := new([]api.Worker)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a worker with the provided details.
func (svc *WorkerService) Add(w *api.Worker) (*api.Token, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/workers"

	// API Token type we want to return
	v := new(api.Token)

	// send request using client
	resp, err := svc.client.Call("POST", u, w, v)

	return v, resp, err
}

// RefreshAuth exchanges a worker token for a new one.
func (svc *WorkerService) RefreshAuth(worker string) (*api.Token, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/workers/%s/refresh", worker)

	// API Token type we want to return
	v := new(api.Token)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

	return v, resp, err
}

// Update modifies a worker with the provided details.
func (svc *WorkerService) Update(worker string, w *api.Worker) (*api.Worker, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/workers/%s", worker)

	// API Worker type we want to return
	v := new(api.Worker)

	// send request using client
	resp, err := svc.client.Call("PUT", u, w, v)

	return v, resp, err
}

// Remove deletes the provided worker.
func (svc *WorkerService) Remove(worker string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/workers/%s", worker)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
