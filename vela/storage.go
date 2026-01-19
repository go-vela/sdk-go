// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"

	api "github.com/go-vela/server/api/types"
)

// StorageService handles retrieving storage s3 info from
// the server methods of the Vela API.
type StorageService service

// GetInfo fetches queue info, primarily used during worker onboarding.
func (svc *StorageService) GetInfo(ctx context.Context) (*api.StorageInfo, *Response, error) {
	// set the API endpoint path we send the request to
	url := "/api/v1/storage/info"

	// API StorageInfo type we want to return
	t := new(api.StorageInfo)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", url, nil, t)

	return t, resp, err
}
