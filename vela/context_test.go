// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"context"
	"testing"
)

func TestVela_FromContext(t *testing.T) {
	// setup types
	want := &Client{}

	// setup context
	ctx := context.Background()
	ctx = context.WithValue(ctx, key, want)

	// run test
	got := FromContext(ctx)

	if got != want {
		t.Errorf("FromContext is %v, want %v", got, want)
	}
}

func TestVela_FromContext_InvalidType(t *testing.T) {
	// setup context
	ctx := context.Background()
	ctx = context.WithValue(ctx, key, "foobar")

	// run test
	got := FromContext(ctx)

	if got != nil {
		t.Errorf("FromContext is %v, want nil", got)
	}
}

func TestVela_FromContext_Empty(t *testing.T) {
	// setup context
	ctx := context.Background()

	// run test
	got := FromContext(ctx)

	if got != nil {
		t.Errorf("FromContext is %v, want nil", got)
	}
}

func TestVela_ToContext(t *testing.T) {
	// setup types
	want := &Client{}

	// setup context
	ctx := context.Background()
	ctx = ToContext(ctx, want)

	// run test
	got := ctx.Value(key).(*Client)

	if got != want {
		t.Errorf("ToContext is %v, want %v", got, want)
	}
}
