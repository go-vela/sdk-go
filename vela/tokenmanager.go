// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

// TokenService handles retrieving workers from
// the server methods of the Vela API VADER: TODO: <fix me>.
type TokenService service

// Refresh will take a token and give back a token VADER: TODO: <fix me>.
func (svc *TokenService) Refresh(token string) (string, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/system-refresh"

	// VADER: TODO: <fix me>
	type T struct {
		Token string `json:"JWTToken"`
	}

	// token type we want to return
	t := new(T)
	h := map[string]string{
		"Authorization": token,
	}

	// send request using client
	resp, err := svc.client.CallWithHeaders("POST", u, nil, t, h)

	return t.Token, resp, err
}
