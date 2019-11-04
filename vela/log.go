// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// LogService handles retriving logs for builds
// from the server methods of the Vela API.
type LogService service

// GetService returns the provided service log.
func (s *LogService) GetService(org, repo string, buildNum, serviceNum int) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, buildNum, serviceNum)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// AddService constructs a service log with the provided details.
func (s *LogService) AddService(org, repo string, buildNum, serviceNum int, target *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, buildNum, serviceNum)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := s.client.Call("POST", u, target, v)
	return v, resp, err
}

// UpdateService modifies a service log with the provided details.
func (s *LogService) UpdateService(org, repo string, buildNum, serviceNum int, target *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, buildNum, serviceNum)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// RemoveService deletes the provided service log.
func (s *LogService) RemoveService(org, repo string, buildNum, serviceNum int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, buildNum, serviceNum)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// GetStep returns the provided step log.
func (s *LogService) GetStep(org, repo string, buildNum, stepNum int) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, buildNum, stepNum)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// AddStep constructs a step log with the provided details.
func (s *LogService) AddStep(org, repo string, buildNum, stepNum int, target *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, buildNum, stepNum)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := s.client.Call("POST", u, target, v)
	return v, resp, err
}

// UpdateStep modifies a step log with the provided details.
func (s *LogService) UpdateStep(org, repo string, buildNum, stepNum int, target *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, buildNum, stepNum)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := s.client.Call("PUT", u, target, v)
	return v, resp, err
}

// RemoveStep deletes the provided step log.
func (s *LogService) RemoveStep(org, repo string, buildNum, stepNum int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, buildNum, stepNum)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}
