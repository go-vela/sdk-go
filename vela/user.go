// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// UserService handles retrieving users from
// the server methods of the Vela API.
type UserService service

// Get returns the provided user by name.
func (svc *UserService) Get(ctx context.Context, name string) (*api.User, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/users/%s", name)

	// api user type we want to return
	v := new(api.User)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// GetCurrent returns the current user.
func (svc *UserService) GetCurrent(ctx context.Context) (*api.User, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/user"

	// api user type we want to return
	v := new(api.User)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// Update modifies a user with the provided details.
func (svc *UserService) Update(ctx context.Context, name string, user *api.User) (*api.User, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/users/%s", name)

	// api User type we want to return
	v := new(api.User)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, user, v)

	return v, resp, err
}

// Update modifies the current user with the provided details.
func (svc *UserService) UpdateCurrent(ctx context.Context, user *api.User) (*api.User, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/user"

	// api User type we want to return
	v := new(api.User)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, user, v)

	return v, resp, err
}
