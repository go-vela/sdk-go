// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"errors"
	"fmt"
	"strings"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/api/types/settings"
	"github.com/go-vela/types"
	"github.com/go-vela/types/library"
)

type (
	// AdminService handles retrieving resources from
	// the server methods of the Vela API.
	AdminService struct {
		Build      *AdminBuildService
		Clean      *AdminCleanService
		Deployment *AdminDeploymentService
		Hook       *AdminHookService
		OIDC       *AdminOIDCService
		Repo       *AdminRepoService
		Secret     *AdminSecretService
		Service    *AdminSvcService
		Step       *AdminStepService
		User       *AdminUserService
		Worker     *AdminWorkerService
		Settings   *AdminSettingsService
	}

	// AdminBuildService handles retrieving admin builds from
	// the server methods of the Vela API.
	AdminBuildService service

	// AdminCleanService handles cleaning resources using
	// the server methods of the Vela API.
	AdminCleanService service

	// AdminDeploymentService handles retrieving admin deployments from
	// the server methods of the Vela API.
	AdminDeploymentService service

	// AdminHookService handles retrieving admin hooks from
	// the server methods of the Vela API.
	AdminHookService service

	// AdminOIDCService handles key rotation for OpenID Connect.
	AdminOIDCService service

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

	// AdminWorkerService handles managing admin worker functionality
	// from the server methods of the Vela API.
	AdminWorkerService service

	// AdminSettingsService handles managing admin settings functionality
	// from the server methods of the Vela API.
	AdminSettingsService service
)

// GetQueueOptions specifies the optional parameters to the
// AdminBuildService.GetQueue method.
type GetQueueOptions struct {
	// Unix timestamp.
	// Returns only the builds created since the timestamp.
	// Default: 24 hours ago
	After string `url:"after,omitempty"`

	ListOptions
}

// CleanOptions specifies the optional parameters to the
// Clean.Clean method.
type CleanOptions struct {
	Before int64 `url:"before,omitempty"`
}

// Update modifies a build with the provided details.
func (svc *AdminBuildService) Update(b *api.Build) (*api.Build, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/build"

	// library Build type we want to return
	v := new(api.Build)

	// send request using client
	resp, err := svc.client.Call("PUT", u, b, v)

	return v, resp, err
}

// Clean sets build resources older than a specified time to a proper canceled / finished state with the provided message.
func (svc *AdminCleanService) Clean(e *types.Error, opt *CleanOptions) (*string, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/clean"

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	v := new(string)

	resp, err := svc.client.Call("PUT", u, e, v)

	return v, resp, err
}

// GetQueue returns the list of builds in pending and running status.
func (svc *AdminBuildService) GetQueue(opt *GetQueueOptions) (*[]api.QueueBuild, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/builds/queue"

	// add optional arguments if supplied
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	// BuildQueue type we want to return
	v := new([]api.QueueBuild)

	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies a deployment with the provided details.
func (svc *AdminDeploymentService) Update(d *library.Deployment) (*library.Deployment, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/deployment"

	// library Deployment type we want to return
	v := new(library.Deployment)

	// send request using client
	resp, err := svc.client.Call("PUT", u, d, v)

	return v, resp, err
}

// Update modifies a hook with the provided details.
func (svc *AdminHookService) Update(h *library.Hook) (*library.Hook, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/hook"

	// library Hook type we want to return
	v := new(library.Hook)

	// send request using client
	resp, err := svc.client.Call("PUT", u, h, v)

	return v, resp, err
}

// Update modifies a repo with the provided details.
func (svc *AdminRepoService) Update(r *api.Repo) (*api.Repo, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/repo"

	// library Repo type we want to return
	v := new(api.Repo)

	// send request using client
	resp, err := svc.client.Call("PUT", u, r, v)

	return v, resp, err
}

// Update modifies a secret with the provided details.
func (svc *AdminSecretService) Update(s *library.Secret) (*library.Secret, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/secret"

	// library Secret type we want to return
	v := new(library.Secret)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Update modifies a service with the provided details.
func (svc *AdminSvcService) Update(s *library.Service) (*library.Service, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/service"

	// library Service type we want to return
	v := new(library.Service)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Update modifies a step with the provided details.
func (svc *AdminStepService) Update(s *library.Step) (*library.Step, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/step"

	// library Step type we want to return
	v := new(library.Step)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Update modifies a user with the provided details.
func (svc *AdminUserService) Update(u *api.User) (*api.User, *Response, error) {
	// set the API endpoint path we send the request to
	url := "/api/v1/admin/user"

	// library User type we want to return
	v := new(api.User)

	// send request using client
	resp, err := svc.client.Call("PUT", url, u, v)

	return v, resp, err
}

// Get retrieves the active platform settings.
func (svc *AdminSettingsService) Get() (*settings.Platform, *Response, error) {
	// set the API endpoint path we send the request to
	//nolint:goconst // ignore
	u := "/api/v1/admin/settings"

	// api Settings type we want to return
	v := new(settings.Platform)

	// send request using client
	resp, err := svc.client.Call("GET", u, nil, v)

	return v, resp, err
}

// Update modifies platform settings with the provided details.
func (svc *AdminSettingsService) Update(s *settings.Platform) (*settings.Platform, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/settings"

	// api Settings type we want to return
	v := new(settings.Platform)

	// send request using client
	resp, err := svc.client.Call("PUT", u, s, v)

	return v, resp, err
}

// Restore returns the platform settings to the server's environment-provided defaults.
func (svc *AdminSettingsService) Restore() (*settings.Platform, *Response, error) {
	// set the API endpoint path we send the request to
	u := "/api/v1/admin/settings"

	// api Settings type we want to return
	v := new(settings.Platform)

	// send request using client
	resp, err := svc.client.Call("DELETE", u, nil, v)

	return v, resp, err
}

// RegisterToken generates a worker registration token with the provided details.
func (svc *AdminWorkerService) RegisterToken(hostname string) (*library.Token, *Response, error) {
	// validate input
	if strings.EqualFold(hostname, "") {
		return nil, nil, errors.New("bad request, no hostname provided")
	}

	// set the API endpoint path we send the request to
	url := fmt.Sprintf("/api/v1/admin/workers/%s/register", hostname)

	// library Token type we want to return
	t := new(library.Token)

	// send request using client
	resp, err := svc.client.Call("POST", url, nil, t)

	return t, resp, err
}

// RotateOIDCKeys sends a request to rotate the private keys used for creating ID tokens.
func (svc *AdminOIDCService) RotateOIDCKeys() (*string, *Response, error) {
	// set the API endpoint path we send the request to
	url := "/api/v1/admin/rotate_oidc_keys"

	v := new(string)

	// send request using client
	resp, err := svc.client.Call("POST", url, nil, v)

	return v, resp, err
}
