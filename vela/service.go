// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// SvcService handles retriving services for builds
// from the server methods of the Vela API.
type SvcService service

// Get returns the provided service.
func (svc *SvcService) Get(org, repo string, buildNum, target int) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, buildNum, target)

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetAll returns a list of all services.
func (svc *SvcService) GetAll(org, repo string, buildNum int) (*[]library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services", org, repo, buildNum)

	// slice library Service type we want to return
	v := new([]library.Service)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Add constructs a service with the provided details.
func (svc *SvcService) Add(org, repo string, buildNum int, target *library.Service) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services", org, repo, buildNum)

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("POST", u, target, v)
	return v, resp, err
}

// Update modifies a service with the provided details.
func (svc *SvcService) Update(org, repo string, buildNum int, target *library.Service) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, buildNum, target.GetNumber())

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Remove deletes the provided service.
func (svc *SvcService) Remove(org, repo string, buildNum, target int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, buildNum, target)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
