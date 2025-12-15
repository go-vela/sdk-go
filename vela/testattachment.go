// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// TestAttachmentService handles retrieving a test attachment from
// the server methods of the Vela API.
type TestAttachmentService service

// Add constructs a test attachment with the provided details.
func (svc *TestAttachmentService) Add(ctx context.Context, org, repo string, build int64) (*api.TestAttachment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/reports/testattachment", org, repo, build)

	// API StorageInfo type we want to return
	ta := new(api.TestAttachment)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, nil, ta)

	return ta, resp, err
}

// Update modifies a step with the provided details.
func (svc *TestAttachmentService) Update(ctx context.Context, org, repo string, build int64, ta *api.TestAttachment) (*api.TestAttachment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/reports/testattachment", org, repo, build)

	// API Step type we want to return
	_ta := new(api.TestAttachment)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, ta, _ta)

	return _ta, resp, err
}
