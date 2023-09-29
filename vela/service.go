// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore dupl linter false positive
package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// SvcService handles retrieving services for builds
// from the server methods of the Vela API.
type SvcService service

// Get returns the provided service.
func (svc *SvcService) Get(org, repo string, build, service int) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, build, service)

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all services.
func (svc *SvcService) GetAll(org, repo string, build int, opt *ListOptions) (*[]library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services", org, repo, build)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Service type we want to return
	v := new([]library.Service)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a service with the provided details.
func (svc *SvcService) Add(org, repo string, build int, s *library.Service) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services", org, repo, build)

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("POST", u, s, v)

	return v, resp, err
}

// Update modifies a service with the provided details.
func (svc *SvcService) Update(org, repo string, build int, s *library.Service) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, build, s.GetNumber())

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided service.
func (svc *SvcService) Remove(org, repo string, build, service int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d", org, repo, build, service)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
