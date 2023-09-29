// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
)

// contextKey defines the key type for
// storing database types in a context.
type contextKey int

// key defines the key type for
// storing the Vela client type in a context.
const key contextKey = iota

// FromContext returns the Client associated with this context.
func FromContext(c context.Context) *Client {
	// get client value from context
	v := c.Value(key)
	if v == nil {
		return nil
	}

	// cast client value to expected Client type
	cli, ok := v.(*Client)
	if !ok {
		return nil
	}

	return cli
}

// ToContext adds the Client to the context.
func ToContext(c context.Context, cli *Client) context.Context {
	return context.WithValue(c, key, cli)
}
