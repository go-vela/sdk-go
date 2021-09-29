// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// WorkerService handles retrieving workers from
// the server methods of the Vela API.
type WorkerService service

// Get returns the provided worker.
func (svc *WorkerService) Get(hostname string) (*library.Worker, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/workers/%s", hostname)

	// library Worker type we want to return
	v := new(library.Worker)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all workers.
func (svc *WorkerService) GetAll() (*[]library.Worker, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/workers"

	// slice library Worker type we want to return
	v := new([]library.Worker)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a worker with the provided details.
func (svc *WorkerService) Add(w *library.Worker) (*library.Worker, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/workers"

	// library Worker type we want to return
	v := new(library.Worker)

	// send request using client
	resp, err := svc.client.Call("POST", u, w, v)

	return v, resp, err
}

// Update modifies a worker with the provided details.
//
// nolint: lll // ignore long line length due to variable names
func (svc *WorkerService) Update(worker string, w *library.Worker) (*library.Worker, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/workers/%s", worker)

	// library Worker type we want to return
	v := new(library.Worker)

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
