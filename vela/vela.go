// SPDX-License-Identifier: Apache-2.0

package vela

// Bool is a helper routine that allocates a new boolean
// value to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Bytes is a helper routine that allocates a new byte
// array value to store v and returns a pointer to it.
func Bytes(v []byte) *[]byte { return &v }

// Int is a helper routine that allocates a new integer
// value to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new 64 bit
// integer value to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string
// value to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Strings is a helper routine that allocates a new string
// array value to store v and returns a pointer to it.
func Strings(v []string) *[]string { return &v }
