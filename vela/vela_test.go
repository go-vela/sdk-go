// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"reflect"
	"testing"
)

func TestVela_Bool(t *testing.T) {
	// setup types
	_bool := false

	want := &_bool

	// run test
	got := Bool(_bool)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Bool is %v, want %v", got, want)
	}
}

func TestVela_Bytes(t *testing.T) {
	// setup types
	_bytes := []byte("foo")

	want := &_bytes

	// run test
	got := Bytes(_bytes)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Bytes is %v, want %v", got, want)
	}
}

func TestVela_Int(t *testing.T) {
	// setup types
	_int := 1

	want := &_int

	// run test
	got := Int(_int)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Int is %v, want %v", got, want)
	}
}

func TestVela_Int64(t *testing.T) {
	// setup types
	_int64 := int64(1)

	want := &_int64

	// run test
	got := Int64(_int64)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Int64 is %v, want %v", got, want)
	}
}

func TestVela_String(t *testing.T) {
	// setup types
	_string := "foo"

	want := &_string

	// run test
	got := String(_string)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

func TestVela_Strings(t *testing.T) {
	// setup types
	_strings := []string{"foo"}

	want := &_strings

	// run test
	got := Strings(_strings)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Strings is %v, want %v", got, want)
	}
}
