// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"

	"github.com/go-vela/types/library"
)

type (
	// AdminService handles retrieving resources from
	// the server methods of the Vela API.
	AdminService struct {
		Build      *AdminBuildService
		Deployment *AdminDeploymentService
		Hook       *AdminHookService
		Repo       *AdminRepoService
		Secret     *AdminSecretService
		Service    *AdminSvcService
		Step       *AdminStepService
		User       *AdminUserService
	}

	// AdminBuildService handles retrieving admin builds from
	// the server methods of the Vela API.
	AdminBuildService service

	// AdminDeploymentService handles retrieving admin deployments from
	// the server methods of the Vela API.
	AdminDeploymentService service

	// AdminHookService handles retrieving admin hooks from
	// the server methods of the Vela API.
	AdminHookService service

	// AdminRepoService handles retrieving admin repos from
	// the server methods of the Vela API.
	AdminRepoService service

	// AdminSecretService handles retrieving admin secrets from
	// the server methods of the Vela API.
	AdminSecretService service

	// AdminSvcService handles retrieving admin services from
	// the server methods of the Vela API.
	AdminSvcService service

	// AdminStepService handles retrieving admin steps from
	// the server methods of the Vela API.
	AdminStepService service

	// AdminUserService handles retrieving admin users from
	// the server methods of the Vela API.
	AdminUserService service
)

// GetAll returns a list of all builds.
// nolint
func (svc *AdminBuildService) GetAll(opt *ListOptions) (*[]library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/builds")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Build type we want to return
	v := new([]library.Build)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a build with the provided details.
// nolint
func (svc *AdminBuildService) Update(b *library.Build) (*library.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/build")

	// library Build type we want to return
	v := new(library.Build)

	// send request using client
	resp, err := svc.client.Call("PUT", u, b, v)

	return v, resp, err
}

// GetAll returns a list of all deployments.
// nolint
func (svc *AdminDeploymentService) GetAll(opt *ListOptions) (*[]library.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/deployments")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Deployment type we want to return
	v := new([]library.Deployment)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a deployment with the provided details.
// nolint
func (svc *AdminDeploymentService) Update(d *library.Deployment) (*library.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/deployment")

	// library Deployment type we want to return
	v := new(library.Deployment)

	// send request using client
	resp, err := svc.client.Call("PUT", u, d, v)

	return v, resp, err
}

// GetAll returns a list of all hooks.
// nolint
func (svc *AdminHookService) GetAll(opt *ListOptions) (*[]library.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/hooks")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Hook type we want to return
	v := new([]library.Hook)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a hook with the provided details.
// nolint
func (svc *AdminHookService) Update(h *library.Hook) (*library.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/hook")

	// library Hook type we want to return
	v := new(library.Hook)

	// send request using client
	resp, err := svc.client.Call("PUT", u, h, v)

	return v, resp, err
}

// GetAll returns a list of all repos.
// nolint
func (svc *AdminRepoService) GetAll(opt *ListOptions) (*[]library.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/repos")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Repo type we want to return
	v := new([]library.Repo)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a repo with the provided details.
// nolint
func (svc *AdminRepoService) Update(r *library.Repo) (*library.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/repo")

	// library Repo type we want to return
	v := new(library.Repo)

	// send request using client
	resp, err := svc.client.Call("PUT", u, r, v)

	return v, resp, err
}

// GetAll returns a list of all secrets.
// nolint
func (svc *AdminSecretService) GetAll(opt *ListOptions) (*[]library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/secrets")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Secret type we want to return
	v := new([]library.Secret)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a secret with the provided details.
// nolint
func (svc *AdminSecretService) Update(s *library.Secret) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/secret")

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// GetAll returns a list of all services.
// nolint
func (svc *AdminSvcService) GetAll(opt *ListOptions) (*[]library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/services")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Service type we want to return
	v := new([]library.Service)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a service with the provided details.
// nolint
func (svc *AdminSvcService) Update(s *library.Service) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/service")

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// GetAll returns a list of all steps.
// nolint
func (svc *AdminStepService) GetAll(opt *ListOptions) (*[]library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/steps")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library Step type we want to return
	v := new([]library.Step)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a step with the provided details.
// nolint
func (svc *AdminStepService) Update(s *library.Step) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/step")

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// GetAll returns a list of all users.
// nolint
func (svc *AdminUserService) GetAll(opt *ListOptions) (*[]library.User, *Response, error) {
	// set the API endpoint path we send the request to
	u := fmt.Sprintf("/api/v1/admin/users")

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// slice library User type we want to return
	v := new([]library.User)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a user with the provided details.
// nolint
func (svc *AdminUserService) Update(u *library.User) (*library.User, *Response, error) {
	// set the API endpoint path we send the request to
	url := fmt.Sprintf("/api/v1/admin/user")

	// library User type we want to return
	v := new(library.User)

	// send request using client
	resp, err := svc.client.Call("PUT", url, u, v)

	return v, resp, err
}
