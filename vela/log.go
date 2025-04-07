// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// LogService handles retrieving logs for builds
// from the server methods of the Vela API.
type LogService service

// GetService returns the provided service log.
func (svc *LogService) GetService(org, repo string, build, service int) (*api.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// API Log type we want to return
	v := new(api.Log)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// AddService constructs a service log with the provided details.
func (svc *LogService) AddService(org, repo string, build, service int, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// send request using client
	resp, err := svc.client.Call("POST", u, l, nil)

	return resp, err
}

// UpdateService modifies a service log with the provided details.
func (svc *LogService) UpdateService(org, repo string, build, service int, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// send request using client
	resp, err := svc.client.Call("PUT", u, l, nil)

	return resp, err
}

// RemoveService deletes the provided service log.
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
func (svc *LogService) GetStep(org, repo string, build, step int) (*api.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// API Log type we want to return
	v := new(api.Log)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// AddStep constructs a step log with the provided details.
func (svc *LogService) AddStep(org, repo string, build, step int, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// send request using client
	resp, err := svc.client.Call("POST", u, l, nil)

	return resp, err
}

// UpdateStep modifies a step log with the provided details.
func (svc *LogService) UpdateStep(org, repo string, build int64, step int32, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// send request using client
	resp, err := svc.client.Call("PUT", u, l, nil)

	return resp, err
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
