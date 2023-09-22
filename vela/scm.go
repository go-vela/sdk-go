// SPDX-License-Identifier: Apache-2.0

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
