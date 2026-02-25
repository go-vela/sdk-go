// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// LogService handles retrieving logs for builds
// from the server methods of the Vela API.
type LogService service

// GetService returns the provided service log.
func (svc *LogService) GetService(ctx context.Context, org, repo string, build int64, service int32) (*api.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// API Log type we want to return
	v := new(api.Log)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// AddService constructs a service log with the provided details.
func (svc *LogService) AddService(ctx context.Context, org, repo string, build, service int, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, l, nil)

	return resp, err
}

// UpdateService modifies a service log with the provided details.
func (svc *LogService) UpdateService(ctx context.Context, org, repo string, build int64, service int32, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, l, nil)

	return resp, err
}

// RemoveService deletes the provided service log.
func (svc *LogService) RemoveService(ctx context.Context, org, repo string, build, service int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/services/%d/logs", org, repo, build, service)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "DELETE", u, nil, v)

	return v, resp, err
}

// GetStep returns the provided step log.
func (svc *LogService) GetStep(ctx context.Context, org, repo string, build int64, step int32) (*api.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// API Log type we want to return
	v := new(api.Log)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// AddStep constructs a step log with the provided details.
func (svc *LogService) AddStep(ctx context.Context, org, repo string, build, step int, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, l, nil)

	return resp, err
}

// UpdateStep modifies a step log with the provided details.
func (svc *LogService) UpdateStep(ctx context.Context, org, repo string, build int64, step int32, l *api.Log) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, l, nil)

	return resp, err
}

// RemoveStep deletes the provided step log.
func (svc *LogService) RemoveStep(ctx context.Context, org, repo string, build, step int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/steps/%d/logs", org, repo, build, step)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "DELETE", u, nil, v)

	return v, resp, err
}
