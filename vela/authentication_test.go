// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types/library"
)

func TestVela_Authentication_SetTokenAuth(t *testing.T) {
	// setup types
	c, _ := NewClient("http://localhost:8080", "", nil)

	// run test
	c.Authentication.SetTokenAuth("someToken")

	if !c.Authentication.HasAuth() {
		t.Errorf("SetTokenAuth did not set an authentication type")
	}

	if !c.Authentication.HasTokenAuth() {
		t.Errorf("SetTokenAuth did not set AuthenticationToken type")
	}
}

func TestVela_Authentication_SetAccessAndRefreshAuth(t *testing.T) {
	// setup types
	c, _ := NewClient("http://localhost:8080", "", nil)

	// run test
	c.Authentication.SetAccessAndRefreshAuth("someAccessToken", "someRefreshToken")

	if !c.Authentication.HasAuth() {
		t.Errorf("SetAccessAndRefreshAuth did not set an authentication type")
	}

	if !c.Authentication.HasAccessAndRefreshAuth() {
		t.Errorf("SetAccessAndRefreshAuth did not set AccessAndRefreshToken type")
	}
}

func TestVela_Authentication_IsTokenAuthExpired_ValidAuthToken(t *testing.T) {
	// setup types
	c, _ := NewClient("http://localhost:8080", "", nil)

	// run test
	c.Authentication.SetTokenAuth(TestTokenGood)

	expired, err := c.Authentication.IsTokenAuthExpired()
	if err != nil {
		t.Errorf("IsTokenAuthExpired returned err: %v", err)
	}

	if expired {
		t.Error("IsTokenAuthExpired returned expired for expired token")
	}
}

func TestVela_Authentication_IsTokenAuthExpired_ExpiredAuthToken(t *testing.T) {
	// setup types
	c, _ := NewClient("http://localhost:8080", "", nil)

	// run test
	c.Authentication.SetTokenAuth(TestTokenExpired)

	expired, err := c.Authentication.IsTokenAuthExpired()
	if err != nil {
		t.Errorf("IsTokenAuthExpired returned err: %v", err)
	}

	if !expired {
		t.Error("IsTokenAuthExpired did not return expired for expired token")
	}
}

func TestVela_Authentication_IsTokenAuthExpired_InvalidAuthToken(t *testing.T) {
	// setup types
	c, _ := NewClient("http://localhost:8080", "", nil)

	// run test
	c.Authentication.SetTokenAuth("someToken")

	expired, err := c.Authentication.IsTokenAuthExpired()
	if err != nil {
		t.Errorf("IsTokenAuthExpired returned err: %v", err)
	}

	if !expired {
		t.Error("IsTokenAuthExpired did not return expired for invalid token")
	}
}

func TestVela_Authentication_IsTokenAuthExpired_InvalidAuthType(t *testing.T) {
	// setup types
	c, _ := NewClient("http://localhost:8080", "", nil)

	// run test
	c.Authentication.SetPersonalAccessTokenAuth("somePersonalAccessToken")

	expired, err := c.Authentication.IsTokenAuthExpired()
	if err == nil {
		t.Error("IsTokenAuthExpired did not return err")
	}

	if !expired {
		t.Error("IsTokenAuthExpired did not return expired for invalid token type")
	}

	// run test
	c.Authentication.SetAccessAndRefreshAuth("someAccessToken", "someRefreshToken")

	expired, err = c.Authentication.IsTokenAuthExpired()
	if err == nil {
		t.Error("IsTokenAuthExpired did not return err")
	}

	if !expired {
		t.Error("IsTokenAuthExpired did not return expired for invalid token type")
	}
}

func TestVela_Authentication_RefreshAccessToken(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.TokenRefreshResp)

	var want library.Token
	_ = json.Unmarshal(data, &want)

	// run test
	resp, err := c.Authentication.RefreshAccessToken("refreshToken")

	if err != nil {
		t.Errorf("RefreshAccessToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("RefreshAccessToken returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if *c.Authentication.accessToken != want.GetToken() {
		t.Errorf("RefreshAccessToken didn't return token")
	}
}

func TestVela_Authentication_AuthenticateWithToken(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.TokenRefreshResp)

	var want library.Token
	_ = json.Unmarshal(data, &want)

	// run test
	at, resp, err := c.Authentication.AuthenticateWithToken("personalaccesstoken")

	if err != nil {
		t.Errorf("AuthenticateWithToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("AuthenticateWithToken returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if at != want.GetToken() {
		t.Errorf("AuthenticateWithToken didn't produce the right Access Token")
	}
}

func TestVela_Authentication_AuthenticateWithToken_NoToken(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Authentication.AuthenticateWithToken("")

	if err == nil {
		t.Errorf("AuthenticateWithToken should have returned error")
	}

	if resp != nil {
		t.Errorf("AuthenticateWithToken should be nil")
	}

	if c.Authentication.token != nil {
		t.Errorf("AuthenticateWithToken should not be set")
	}
}

func TestVela_Authentication_ExchangeTokens(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.TokenRefreshResp)

	var want library.Token
	_ = json.Unmarshal(data, &want)

	// create options
	opt := &OAuthExchangeOptions{
		Code:  "42",
		State: "411",
	}

	// hardcoded value in mock
	wantRefresh := "refresh"

	// run test
	at, rt, resp, err := c.Authentication.ExchangeTokens(opt)

	if err != nil {
		t.Errorf("ExchangeTokens returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ExchangeTokens returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if at != want.GetToken() {
		t.Errorf("ExchangeTokens didn't produce the right Access Token")
	}

	if rt != wantRefresh {
		t.Errorf("ExchangeTokens returned unexpected Refresh Token value")
	}

	if *c.Authentication.accessToken != want.GetToken() {
		t.Errorf("ExchangeTokens didn't produce the right Access Token")
	}

	if *c.Authentication.refreshToken != wantRefresh {
		t.Errorf("ExchangeTokens didn't produce the right Refresh Token")
	}
}

func TestVela_Authentication_ExchangeTokens_BadInput(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// create options
	opt := &OAuthExchangeOptions{}

	// run test
	_, _, resp, err := c.Authentication.ExchangeTokens(opt)

	if err == nil {
		t.Errorf("ExchangeTokens should have returned error: %v", err)
	}

	if resp != nil {
		t.Errorf("ExchangeTokens should not return resp")
	}

	if c.Authentication.accessToken != nil {
		t.Errorf("ExchangeTokens should not set Access Token")
	}

	if c.Authentication.refreshToken != nil {
		t.Errorf("ExchangeTokens should not set Refresh Token")
	}
}

func TestVela_Authentication_ValidateToken_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	c.Authentication.SetTokenAuth("foo")

	// run test
	resp, err := c.Authentication.ValidateToken()

	if err != nil {
		t.Errorf("ValidateToken returned error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ValidateToken returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestVela_Authentication_ValidateToken_NoToken(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	c.Authentication.SetTokenAuth("")

	// run test
	resp, err := c.Authentication.ValidateToken()

	if err == nil {
		t.Error("ValidateToken should have returned error")
	}

	if resp != nil {
		t.Error("ValidateToken response should be nil")
	}
}

func TestVela_Authentication_ValidateOAuthToken_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	c.Authentication.SetTokenAuth("foo")

	// run test
	resp, err := c.Authentication.ValidateOAuthToken()

	if err != nil {
		t.Errorf("ValidateOAuthToken returned error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ValidateOAuthToken returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestVela_Authentication_ValidateOAuthToken_NoToken(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	c.Authentication.SetTokenAuth("")

	// run test
	resp, err := c.Authentication.ValidateOAuthToken()

	if err == nil {
		t.Error("ValidateOAuthToken should have returned error")
	}

	if resp != nil {
		t.Error("ValidateOAuthToken response should be nil")
	}
}
