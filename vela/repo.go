// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// RepoService handles retriving repos from
// the server methods of the Vela API.
type RepoService service

// Get returns the provided repo.
func (s *RepoService) Get(org, repo string) (*library.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s", org, repo)

	// library Repo type we want to return
	v := new(library.Repo)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetAll returns a list of all repos.
func (s *RepoService) GetAll() (*[]library.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos")

	// slice library Repo type we want to return
	v := new([]library.Repo)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Add constructs a repo with the provided details.
func (s *RepoService) Add(target *library.Repo) (*library.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos")

	// library Repo type we want to return
	v := new(library.Repo)

	// send request using client
	resp, err := s.client.Call("POST", u, target, v)
	return v, resp, err
}

// Update modifies a repo with the provided details.
func (s *RepoService) Update(org, repo string, target *library.Repo) (*library.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s", org, repo)

	// library Repo type we want to return
	v := new(library.Repo)

	// send request using client
	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// Remove deletes the provided repo.
func (s *RepoService) Remove(org, repo string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s", org, repo)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// Repair modifies a damaged repo webhook.
func (s *RepoService) Repair(org, repo string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/repair", org, repo)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("PATCH", u, nil, v)
	return v, resp, err
}

// Chown modifies the org of a repo.
func (s *RepoService) Chown(org, repo string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/chown", org, repo)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("PATCH", u, nil, v)
	return v, resp, err
}
