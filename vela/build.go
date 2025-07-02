// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"fmt"

	api "github.com/go-vela/server/api/types"
)

// BuildService handles retrieving builds from
// the server methods of the Vela API.
type BuildService service

// BuildListOptions specifies the optional parameters to the
// Build.GetAll method.
type BuildListOptions struct {
	Branch string `url:"branch,omitempty"`
	Event  string `url:"event,omitempty"`
	Status string `url:"status,omitempty"`
	Before int64  `url:"before,omitempty"`
	After  int64  `url:"after,omitempty"`

	ListOptions
}

// RequestTokenOptions specifies the required parameters to the
// Build.GetIDRequestToken method.
type RequestTokenOptions struct {
	Image    string `url:"image,omitempty"`
	Request  string `url:"request,omitempty"`
	Commands bool   `url:"commands,omitempty"`
}

// IDTokenOptions specifies the required parameters to the
// Build.GetIDToken method.
type IDTokenOptions struct {
	Audience []string `url:"audience,omitempty"`
}

// Get returns the provided build.
func (svc *BuildService) Get(org, repo string, build int64) (*api.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, build)

	// API Build type we want to return
	v := new(api.Build)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetBuildExecutable returns the executable for the provided build.
func (svc *BuildService) GetBuildExecutable(org, repo string, build int64) (*api.BuildExecutable, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/executable", org, repo, build)

	// API Build type we want to return
	v := new(api.BuildExecutable)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetAll returns a list of all builds.
func (svc *BuildService) GetAll(org, repo string, opt *BuildListOptions) (*[]api.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds", org, repo)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice API Build type we want to return
	v := new([]api.Build)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// GetLogs returns the provided build logs.
func (svc *BuildService) GetLogs(org, repo string, build int64, opt *ListOptions) (*[]api.Log, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/logs", org, repo, build)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice database Log type we want to return
	v := new([]api.Log)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Add constructs a build with the provided details.
func (svc *BuildService) Add(b *api.Build) (*api.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds", b.GetRepo().GetOrg(), b.GetRepo().GetName())

	// API Build type we want to return
	v := new(api.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, b, v)

	return v, resp, err
}

// Update modifies a build with the provided details.
func (svc *BuildService) Update(b *api.Build) (*api.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", b.GetRepo().GetOrg(), b.GetRepo().GetName(), b.GetNumber())

	// API Build type we want to return
	v := new(api.Build)

	// send request using client
	resp, err := svc.client.Call("PUT", u, b, v)

	return v, resp, err
}

// Remove deletes the provided build.
func (svc *BuildService) Remove(org, repo string, build int) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, build)

	// string type we want to return
	v := new(string)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}

// Restart takes the build provided and restarts it.
func (svc *BuildService) Restart(org, repo string, build int64) (*api.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d", org, repo, build)

	// API Build type we want to return
	v := new(api.Build)

	// send request using client
	resp, err := svc.client.Call("POST", u, nil, v)

	return v, resp, err
}

// Cancel takes the build provided and cancels it.
func (svc *BuildService) Cancel(org, repo string, build int64) (*api.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/cancel", org, repo, build)

	// API Build type we want to return
	v := new(api.Build)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}

// Approve takes the build provided and approves it as an admin.
func (svc *BuildService) Approve(org, repo string, build int64) (*Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/approve", org, repo, build)

	return svc.client.Call("POST", u, nil, nil)
}

// GetBuildToken returns an auth token for updating build resources.
func (svc *BuildService) GetBuildToken(org, repo string, build int64) (*api.Token, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/token", org, repo, build)

	// API Token type we want to return
	t := new(api.Token)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, t)

	return t, resp, err
}

// GetIDRequestToken returns an id request token for integrating with build OIDC.
func (svc *BuildService) GetIDRequestToken(org, repo string, build int64, opt *RequestTokenOptions) (*api.Token, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/id_request_token", org, repo, build)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// API Token type we want to return
	t := new(api.Token)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, t)

	return t, resp, err
}

// GetIDToken returns an ID token corresponding to the request token during a build.
func (svc *BuildService) GetIDToken(org, repo string, build int, opt *IDTokenOptions) (*api.Token, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/repos/%s/%s/builds/%d/id_token", org, repo, build)

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// API Token type we want to return
	t := new(api.Token)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, t)

	return t, resp, err
}
