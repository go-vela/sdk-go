// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// SecretService handles retrieving secrets from
// the server methods of the Vela API.
type SecretService service

// Get returns the provided secret.
func (svc *SecretService) Get(engine, sType, org, name, secret string) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, secret)

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all secrets.
func (svc *SecretService) GetAll(engine, sType, org, name string, opt *ListOptions) (*[]library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s", engine, sType, org, name)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Secret type we want to return
	v := new([]library.Secret)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a secret with the provided details.
func (svc *SecretService) Add(engine, sType, org, name string, s *library.Secret) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s", engine, sType, org, name)

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := svc.client.Call("POST", u, s, v)

	return v, resp, err
}

// Update modifies a secret with the provided details.
func (svc *SecretService) Update(engine, sType, org, name string, s *library.Secret) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, s.GetName())

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Remove deletes the provided secret.
func (svc *SecretService) Remove(engine, sType, org, name, secret string) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/secrets/%s/%s/%s/%s/%s", engine, sType, org, name, secret)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}
