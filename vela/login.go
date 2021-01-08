package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

// AuthorizationService handles user login actions
// against the server methods of the Vela API.
type AuthorizationService service

// TODO: change this to only login via PAT
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

func (svc *AuthorizationService) GetLoginURL(opt *LoginOpts) (string, error) {
	var err error

	l := "/login"

	if opt != nil && len(opt.Type) > 0 {
		l, err = addOptions(l, opt)
		if err != nil {
			return "", err
		}
	}

	loginURL, err := svc.client.buildURLForRequest(l)
	if err != nil {
		return "", err
	}

	return loginURL, nil
}
