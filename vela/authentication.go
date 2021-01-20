// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"net/http"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
)

type AuthenticationType int

const (
	// AuthenticationToken defines the
	// authentication type for auth tokens.
	AuthenticationToken AuthenticationType = iota + 1
	AccessAndRefreshToken
)

// AuthenticationService contains
// authentication related functions.
type AuthenticationService struct {
	client       *Client
	token        *string
	accessToken  *string
	refreshToken *string
	authType     AuthenticationType
}

// SetTokenAuth sets the authentication type as personal access token.
func (svc *AuthenticationService) SetTokenAuth(token string) {
	svc.token = String(token)
	svc.authType = AuthenticationToken
}

// SetAccessAndRefreshAuth sets the authentication type as oauth token pair.
func (svc *AuthenticationService) SetAccessAndRefreshAuth(access, refresh string) {
	svc.accessToken = String(access)
	svc.refreshToken = String(refresh)
	svc.authType = AccessAndRefreshToken
}

// HasAuth checks if the authentication type is set.
func (svc *AuthenticationService) HasAuth() bool {
	return svc.authType > 0
}

// HasTokenAuth checks if the authentication type is a personal access token.
func (svc *AuthenticationService) HasTokenAuth() bool {
	return svc.authType == AuthenticationToken
}

// HasAccessAndRefreshAuth checks if the authentication type is oauth token pair.
func (svc *AuthenticationService) HasAccessAndRefreshAuth() bool {
	return svc.authType == AccessAndRefreshToken
}

// RefreshAccessToken uses the supplied refresh token to attempt and refresh
// the access token.
func (svc *AuthenticationService) RefreshAccessToken(refreshToken string) (*Response, error) {
	u := "/token-refresh"

	v := new(library.Login)

	// building a custom request -
	// we can't use svc.client.NewRequest because
	// that's what can send us here
	url, err := svc.client.buildURLForRequest(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// set a minimal cookie with the refresh token value
	cookie := &http.Cookie{
		Name:  constants.RefreshTokenName,
		Value: refreshToken,
	}

	req.AddCookie(cookie)

	// send the request
	resp, err := svc.client.Do(req, v)

	// set the received access token
	svc.accessToken = v.Token

	return resp, err
}
