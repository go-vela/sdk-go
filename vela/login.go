// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"
)

// AuthorizationService handles user login actions
// against the server methods of the Vela API.
type AuthorizationService service

// GetLoginURL returns the login url with the give login options.
func (svc *AuthorizationService) GetLoginURL(opt *LoginOptions) (string, error) {
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

	// check that we have a client
	if svc.client == nil {
		return "", fmt.Errorf("client not found")
	}

	// build the url
	loginURL, err := svc.client.buildURLForRequest(l)
	if err != nil {
		return "", err
	}

	// return the url
	return loginURL, nil
}
