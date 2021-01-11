// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"encoding/json"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// IsTokenExpired will parse the expiration of the the given
// token and return a boolean depending on whether the is
// expired given the delta.
func IsTokenExpired(token string) bool {
	minTimeLeft := 10 * time.Second

	if len(token) == 0 {
		return true
	}

	t, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return true
	}

	c, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return true
	}

	var expiration time.Time
	switch e := c["exp"].(type) {
	case float64:
		expiration = time.Unix(int64(e), 0)
	case json.Number:
		v, _ := e.Int64()
		expiration = time.Unix(v, 0)
	}

	timeLeft := time.Until(expiration)

	return timeLeft.Seconds() <= minTimeLeft.Seconds()
}
