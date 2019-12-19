package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// AuthorizationService handles user login actions
// against the server methods of the Vela API.
type AuthorizationService service

// Login constructs a build with the provided details.
func (svc *AuthorizationService) Login(l *library.Login) (*library.Login, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/login")

	// library Login type we want to return
	v := new(library.Login)

	// send request using client
	resp, err := svc.client.Call("POST", u, l, v)

	return v, resp, err
}
