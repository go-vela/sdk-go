// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// SecretService handles retrieving secrets from
// the server methods of the Vela API.
type SecretService service

// Get returns the provided secret.
func (svc *SecretService) Get(ctx context.Context, engine, sType, org, name, secret string) (*api.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, secret)

	// API Secret type we want to return
	v := new(api.Secret)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all secrets.
func (svc *SecretService) GetAll(ctx context.Context, engine, sType, org, name string, opt *ListOptions) (*[]api.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s", engine, sType, org, name)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Secret type we want to return
	v := new([]api.Secret)

	// send request using client
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)

	return v, resp, err
}

// Add constructs a secret with the provided details.
func (svc *SecretService) Add(ctx context.Context, engine, sType, org, name string, s *api.Secret) (*api.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s", engine, sType, org, name)

	// API Secret type we want to return
	v := new(api.Secret)

	// send request using client
	resp, err := svc.client.Call(ctx, "POST", u, s, v)

	return v, resp, err
}

// Update modifies a secret with the provided details.
func (svc *SecretService) Update(ctx context.Context, engine, sType, org, name string, s *api.Secret) (*api.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, s.GetName())

	// API Secret type we want to return
	v := new(api.Secret)

	// send request using client
	resp, err := svc.client.Call(ctx, "PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided secret.
func (svc *SecretService) Remove(ctx context.Context, engine, sType, org, name, secret string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, secret)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call(ctx, "DELETE", u, nil, v)

	return v, resp, err
}
