// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"testing"
)

func TestVela_Authentication_SetTokenAuth(t *testing.T) {
	// setup types
	c, _ := NewClient("http://localhost:8080", nil)

	// run test
	c.Authentication.SetTokenAuth("someToken")

	if !c.Authentication.HasAuth() {
		t.Errorf("SetTokenAuth did not set an authentication type")
	}

	if !c.Authentication.HasTokenAuth() {
		t.Errorf("SetTokenAuth did not set AuthenticationToken type")
	}
}
