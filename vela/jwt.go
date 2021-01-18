// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// IsTokenExpired will parse the expiration of the the given
// token and return a boolean depending on whether the is
// expired given the delta.
func IsTokenExpired(token string) bool {
	minTimeLeft := 10 * time.Second

	// if the token is empty, we treat it as expired
	if len(token) == 0 {
		return true
	}

	// parse the token, we just want to check expiration -
	// the server will handle verification
	t, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return true
	}

	// get the claims
	c, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return true
	}

	// if there is no expiration set, it's expired
	if _, ok := c["exp"]; !ok {
		return true
	}

	// check the expiration
	expiration := time.Unix(int64(c["exp"].(float64)), 0)

	// get the difference
	timeLeft := time.Until(expiration)

	// return whether we are within the delta time padding
	return timeLeft.Seconds() <= minTimeLeft.Seconds()
}
