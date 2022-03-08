// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"
)

// SCMService handles syncing repos from
// the server methods of the Vela API.
type SCMService service

// Sync synchronizes a repo between the database and the SCM.
func (svc *SCMService) Sync(org, repo string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/v1/scm/repos/%s/%s/sync", org, repo)
	v := new(string)
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Sync synchronizes all org repos between the database and the SCM.
func (svc *SCMService) SyncAll(org string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/v1/scm/orgs/%s/sync", org)
	v := new(string)
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}
