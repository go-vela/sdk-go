// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Bytes is a helper routine that allocates a new byte array value
// to store v and returns a pointer to it.
func Bytes(v []byte) *[]byte { return &v }
