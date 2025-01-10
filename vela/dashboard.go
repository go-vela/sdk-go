// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// DashboardService handles retrieving Dashboards from
// the server methods of the Vela API.
type DashboardService service

// Get returns the provided Dashboard.
func (svc *DashboardService) Get(dashboard string) (*api.DashCard, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/dashboards/%s", dashboard)

	// API Dashboard type we want to return
	v := new(api.DashCard)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAllUser returns a list of all dashboards for the authenticated user.
func (svc *DashboardService) GetAllUser() (*[]api.DashCard, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/user/dashboards"

	// slice API Dashboard type we want to return
	v := new([]api.DashCard)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a Dashboard with the provided details.
func (svc *DashboardService) Add(d *api.Dashboard) (*api.Dashboard, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/dashboards"

	// api dashboard type we want to return
	v := new(api.Dashboard)

	// send request using client
	resp, err := svc.client.Call("POST", u, d, v)

	return v, resp, err
}

// Update modifies a dashboard with the provided details.
func (svc *DashboardService) Update(d *api.Dashboard) (*api.Dashboard, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/dashboards/%s", d.GetID())

	// API dashboard type we want to return
	v := new(api.Dashboard)

	// send request using client
	resp, err := svc.client.Call("PUT", u, d, v)

	return v, resp, err
}
