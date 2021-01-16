// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/mock/server"
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

func TestVela_Authentication_RefreshAccessToken(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.TokenRefreshResp)

	var want library.Login
	_ = json.Unmarshal(data, &want)

	// run test
	resp, err := c.Authentication.RefreshAccessToken("refreshToken")

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("RefreshAccessToken returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if *c.Authentication.accessToken != want.GetToken() {
		t.Errorf("RefreshAccessToken didn't return token")
	}
}
