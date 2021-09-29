// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// LogService handles retrieving logs for builds
// from the server methods of the Vela API.
type LogService service

// GetService returns the provided service log.
//
// nolint: lll // ignore long line length due to variable names
func (svc *LogService) GetService(org, repo string, build, service int) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// AddService constructs a service log with the provided details.
//
// nolint: lll // ignore long line length due to variable names
func (svc *LogService) AddService(org, repo string, build, service int, l *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := svc.client.Call("POST", u, l, v)

	return v, resp, err
}

// UpdateService modifies a service log with the provided details.
//
// nolint: lll // ignore long line length due to variable names
func (svc *LogService) UpdateService(org, repo string, build, service int, l *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := svc.client.Call("PUT", u, l, v)

	return v, resp, err
}

// RemoveService deletes the provided service log.
//
// nolint: lll // ignore long line length due to variable names
func (svc *LogService) RemoveService(org, repo string, build, service int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}

// GetStep returns the provided step log.
func (svc *LogService) GetStep(org, repo string, build, step int) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// AddStep constructs a step log with the provided details.
//
// nolint: lll // ignore long line length due to variable names
func (svc *LogService) AddStep(org, repo string, build, step int, l *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := svc.client.Call("POST", u, l, v)

	return v, resp, err
}

// UpdateStep modifies a step log with the provided details.
//
// nolint: lll // ignore long line length due to variable names
func (svc *LogService) UpdateStep(org, repo string, build, step int, l *library.Log) (*library.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// library Log type we want to return
	v := new(library.Log)

	// send request using client
	resp, err := svc.client.Call("PUT", u, l, v)

	return v, resp, err
}

// RemoveStep deletes the provided step log.
func (svc *LogService) RemoveStep(org, repo string, build, step int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
