// SPDX-License-Identifier: Apache-2.0

package vela

// Bool is a helper routine that allocates a new boolean
// value to store v and returns a pointer to it.
//
//go:fix inline
func Bool(v bool) *bool { return new(v) }

// Bytes is a helper routine that allocates a new byte
// array value to store v and returns a pointer to it.
//
//go:fix inline
func Bytes(v []byte) *[]byte { return new(v) }

// Int is a helper routine that allocates a new integer
// value to store v and returns a pointer to it.
//
//go:fix inline
func Int(v int) *int { return new(v) }

// Int32 is a helper routine that allocates a new 32 bit
// integer value to store v and returns a pointer to it.
//
//go:fix inline
func Int32(v int32) *int32 { return new(v) }

// Int64 is a helper routine that allocates a new 64 bit
// integer value to store v and returns a pointer to it.
//
//go:fix inline
func Int64(v int64) *int64 { return new(v) }

// UInt64 is a helper routine that allocates a new unsigned 64 bit
// integer value to store v and returns a pointer to it.
//
//go:fix inline
func UInt64(v uint64) *uint64 { return new(v) }

// String is a helper routine that allocates a new string
// value to store v and returns a pointer to it.
//
//go:fix inline
func String(v string) *string { return new(v) }

// Strings is a helper routine that allocates a new string
// array value to store v and returns a pointer to it.
//
//go:fix inline
func Strings(v []string) *[]string { return new(v) }
