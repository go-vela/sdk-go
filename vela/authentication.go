// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

const (
	// AuthenticationToken defines the
	// authentication type for OAuth tokens.
	AuthenticationToken = 1
)

// AuthenticationService contains
// authentication related functions.
type AuthenticationService struct {
	client   *Client
	secret   *string
	authType int
}

// SetTokenAuth sets the authentication type as OAuth Token.
func (s *AuthenticationService) SetTokenAuth(token string) {
	s.secret = String(token)
	s.authType = AuthenticationToken
}

// HasAuth checks if the authentication type is set.
func (s *AuthenticationService) HasAuth() bool {
	return s.authType > 0
}

// HasTokenAuth checks if the authentication type is OAuth Token.
func (s *AuthenticationService) HasTokenAuth() bool {
	return s.authType == AuthenticationToken
}
