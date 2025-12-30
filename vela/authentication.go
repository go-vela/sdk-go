// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/constants"
)

type AuthenticationType int

const (
	// AuthenticationToken defines the
	// authentication type for auth tokens.
	AuthenticationToken AuthenticationType = iota + 1
	PersonalAccessToken
	AccessAndRefreshToken
	BuildToken
)

// AuthenticationService contains
// authentication related functions.
type AuthenticationService struct {
	client              *Client
	token               *string
	personalAccessToken *string
	accessToken         *string
	refreshToken        *string
	scmToken            *string
	authType            AuthenticationType
	scmTokenExp         *int64
	buildRepo           *string
	buildNumber         *int64
}

// SetTokenAuth sets the authentication type as a plain token.
func (svc *AuthenticationService) SetTokenAuth(token string) {
	svc.token = String(token)
	svc.authType = AuthenticationToken
}

// SetBuildTokenAuth sets the authentication type and the two tokens used.
func (svc *AuthenticationService) SetBuildTokenAuth(buildTkn, scmTkn string, scmTokenExp int64, buildRepo string, buildNumber int64) {
	svc.token = String(buildTkn)
	svc.scmToken = String(scmTkn)
	svc.buildRepo = String(buildRepo)
	svc.buildNumber = Int64(buildNumber)

	// set expiration if provided - only for installation tokens
	if scmTokenExp > 0 {
		svc.scmTokenExp = &scmTokenExp
	}

	svc.authType = BuildToken
}

// SetPersonalAccessTokenAuth sets the authentication type as personal access token.
func (svc *AuthenticationService) SetPersonalAccessTokenAuth(token string) {
	svc.personalAccessToken = String(token)
	svc.authType = PersonalAccessToken
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

// HasTokenAuth checks if the authentication type is a plain token.
func (svc *AuthenticationService) HasTokenAuth() bool {
	return svc.authType == AuthenticationToken
}

// HasBuildTokenAuth checks if the authentication type is a build and scm token.
func (svc *AuthenticationService) HasBuildTokenAuth() bool {
	return svc.authType == BuildToken
}

// HasPersonalAccessTokenAuth checks if the authentication type is a personal access token.
func (svc *AuthenticationService) HasPersonalAccessTokenAuth() bool {
	return svc.authType == PersonalAccessToken
}

// HasAccessAndRefreshAuth checks if the authentication type is oauth token pair.
func (svc *AuthenticationService) HasAccessAndRefreshAuth() bool {
	return svc.authType == AccessAndRefreshToken
}

// IsTokenAuthExpired returns whether or not the authentication token has expired.
func (svc *AuthenticationService) IsTokenAuthExpired() (bool, error) {
	// verify that the auth type is valid for this type of validation
	if !svc.HasTokenAuth() {
		return true, errors.New("client auth type is not set to auth token")
	}

	// verify a token exists in the client
	if svc.token == nil {
		return true, errors.New("no token in client")
	}

	// check auth token expiration
	return IsTokenExpired(*svc.token), nil
}

// IsSCMTokenExpired checks if the SCM token has expired.
func (svc *AuthenticationService) IsSCMTokenExpired() bool {
	// 5 minute buffer
	if svc.scmTokenExp != nil && time.Now().Unix() >= (*svc.scmTokenExp-300) {
		return true
	}

	return false
}

// SCMExpiration returns the SCM token expiration time.
func (svc *AuthenticationService) SCMExpiration() int64 {
	if svc.scmTokenExp != nil {
		return *svc.scmTokenExp
	}

	return 0
}

// SCMToken returns the SCM token.
func (svc *AuthenticationService) SCMToken() string {
	if svc.scmToken != nil {
		return *svc.scmToken
	}

	return ""
}

// RefreshAccessToken uses the supplied refresh token to attempt and refresh
// the access token.
func (svc *AuthenticationService) RefreshAccessToken(ctx context.Context, refreshToken string) (*Response, error) {
	// set the API endpoint path we send the request to
	u := "/token-refresh"

	v := new(api.Token)

	// building a custom request -
	// we can't use svc.client.NewRequest because
	// that's what can send us here
	url, err := svc.client.buildURLForRequest(u)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

// AuthenticateWithToken attempts to authenticate with the provided token, typically
// a personal access token created in the source provider, eg. GitHub. It will
// return a short-lived Vela Access Token, if successful.
func (svc *AuthenticationService) AuthenticateWithToken(ctx context.Context, token string) (string, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/authenticate/token"

	// check for token
	if len(token) == 0 {
		return "", nil, fmt.Errorf("token must not be empty")
	}

	// will hold access token
	v := new(api.Token)

	// building a custom request -
	// we can't use svc.client.NewRequest because
	// that's what can send us here
	url, err := svc.client.buildURLForRequest(u)
	if err != nil {
		return "", nil, err
	}

	// create a new request that we can attach a header to
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return "", nil, err
	}

	// add the token as a header
	req.Header.Add("Token", token)

	// send the request
	resp, err := svc.client.Do(req, v)

	return v.GetToken(), resp, err
}

// ExchangeTokens handles the last part of the OAuth flow. It uses the supplied
// code and state values to attempt to exchange them for Vela Access and
// Refresh tokens.
func (svc *AuthenticationService) ExchangeTokens(ctx context.Context, opt *OAuthExchangeOptions) (string, string, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/authenticate"

	// will hold access token
	v := new(api.Token)

	// check required arguments
	if len(opt.Code) == 0 || len(opt.State) == 0 {
		return "", "", nil, fmt.Errorf("code and state must be provided")
	}

	// add required arguments
	u, err := addOptions(u, opt)
	if err != nil {
		return "", "", nil, err
	}

	// attempt to exchange code + state for tokens
	resp, err := svc.client.Call(ctx, "GET", u, nil, v)
	if err != nil {
		return "", "", resp, err
	}

	// the refresh token will be in a cookie in the response
	rt := extractRefreshToken(resp.Cookies())

	// get the access token
	at := v.GetToken()

	// set the received tokens
	svc.accessToken = &at
	svc.refreshToken = &rt

	return at, rt, resp, err
}

// extractRefreshToken is a helper function to extract
// the refresh token from the supplied cookie slice.
func extractRefreshToken(cookies []*http.Cookie) string {
	c := ""

	// loop over the cookies to find the refresh cookie
	for _, cookie := range cookies {
		if cookie.Name == constants.RefreshTokenName {
			c = cookie.Value
		}
	}

	return c
}

// ValidateToken makes a request to validate tokens with the Vela server.
func (svc *AuthenticationService) ValidateToken(ctx context.Context) (*Response, error) {
	// set the API endpoint path we send the request to
	u := "/validate-token"

	// attempt to validate a server token
	resp, err := svc.client.Call(ctx, "GET", u, nil, nil)

	return resp, err
}

// ValidateOAuthToken makes a request to validate user oauth tokens with the Vela server.
func (svc *AuthenticationService) ValidateOAuthToken(ctx context.Context) (*Response, error) {
	// set the API endpoint path we send the request to
	u := "/validate-oauth"

	// attempt to validate an oauth token
	resp, err := svc.client.Call(ctx, "GET", u, nil, nil)

	return resp, err
}

// RefreshInstallToken refreshes the SCM install token for a build.
func (svc *AuthenticationService) RefreshInstallToken(ctx context.Context, org, repo string, build int64) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/install_token", org, repo, build)

	// will hold access token
	v := new(api.Token)

	// building a custom request -
	// we can't use svc.client.NewRequest because
	// that's what can send us here
	url, err := svc.client.buildURLForRequest(u)
	if err != nil {
		return nil, err
	}

	// create a new request that we can attach a header to
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *svc.token))
	req.Header.Add("Token", *svc.scmToken)

	resp, err := svc.client.Do(req, v)

	// set the received access token
	svc.scmToken = v.Token
	svc.scmTokenExp = v.Expiration

	return resp, err
}
