// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"crypto/rsa"
	"io/ioutil"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestIsTokenExpired(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "expired token",
			args: args{
				token: makeSampleToken(jwt.MapClaims{"exp": float64(time.Now().Unix() - 100)}),
			},
			want: true,
		},
		{
			name: "good token",
			args: args{
				token: makeSampleToken(jwt.MapClaims{"exp": float64(time.Now().Unix() + 100)}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTokenExpired(tt.args.token); got != tt.want {
				t.Errorf("IsTokenExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeSampleToken(c jwt.Claims) string {
	key := loadKey("testdata/sample_key")
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	s, e := t.SignedString(key)

	if e != nil {
		return ""
	}

	return s
}

func loadKey(loc string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(loc)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}
