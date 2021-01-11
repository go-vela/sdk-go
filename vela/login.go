// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import "github.com/go-vela/types/library"

// AuthorizationService handles user login actions
// against the server methods of the Vela API.
type AuthorizationService service

// Login constructs a build with the provided details.
func (svc *AuthorizationService) Login(l *library.Login) (*library.Login, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/login"

	// library Login type we want to return
	v := new(library.Login)

	// send request using client
	resp, err := svc.client.Call("POST", u, l, v)

	return v, resp, err
}

// GetLoginURL returns the login url with the give login options.
func (svc *AuthorizationService) GetLoginURL(opt *LoginOpts) (string, error) {
	var err error

	// set the API endpoint path we send the request to
	l := "/login"

	// add the login options to the request
	if opt != nil && len(opt.Type) > 0 {
		l, err = addOptions(l, opt)
		if err != nil {
			return "", err
		}
	}

	// build the url
	loginURL, err := svc.client.buildURLForRequest(l)
	if err != nil {
		return "", err
	}

	// return the url
	return loginURL, nil
}
