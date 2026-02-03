// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// StorageService handles retrieving storage s3 info from
// the server methods of the Vela API.
type StorageService service

func (s *StorageService) GetSTSCreds(ctx context.Context, org, repo string, build int64) (*api.STSCreds, *Response, error) {
	u := fmt.Sprintf("/api/v1/storage/sts/%s/%s/%d", org, repo, build)
	out := new(api.STSCreds)
	resp, err := s.client.Call(ctx, "GET", u, nil, out)
	return out, resp, err
}
