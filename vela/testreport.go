// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// TestReportService handles retrieving a test report from
// the server methods of the Vela API.
type TestReportService service

// Add constructs a test report with the provided details.
func (svc *TestReportService) Add(org, repo string, build int64) (*api.TestReport, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/reports/testreport", org, repo, build)

	// API StorageInfo type we want to return
	tr := new(api.TestReport)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, tr)

	return tr, resp, err
}

// Update modifies a step with the provided details.
func (svc *TestReportService) Update(org, repo string, build int64, tr *api.TestReport) (*api.TestReport, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/reports/testreport", org, repo, build)

	// API Step type we want to return
	_tr := new(api.TestReport)

	// send request using client
	resp, err := svc.client.Call("PUT", u, nil, _tr)

	return tr, resp, err
}
