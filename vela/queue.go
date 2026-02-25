// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"

	api "github.com/go-vela/server/api/types"
)

// QueueService handles retrieving queue info from
// the server methods of the Vela API.
type QueueService service

// GetInfo fetches queue info, primarily used during worker onboarding.
func (qvc *QueueService) GetInfo(ctx context.Context) (*api.QueueInfo, *Response, error) {
	// set the API endpoint path we send the request to
	url := "/api/v1/queue/info"

	// API QueueInfo type we want to return
	t := new(api.QueueInfo)

	// send request using client
	resp, err := qvc.client.Call(ctx, "GET", url, nil, t)

	return t, resp, err
}
