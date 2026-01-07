// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"go.yaml.in/yaml/v3"

	"github.com/go-vela/sdk-go/version"
	api "github.com/go-vela/server/api/types"
)

const (
	userAgent = "vela-sdk-go"
)

type (
	// Client is a client that manages communication with the Vela API.
	Client struct {
		// HTTP client used to communicate with the Vela API.
		client *http.Client

		// Base URL for Vela API requests.
		baseURL *url.URL

		// User agent used when communicating with the Vela API.
		UserAgent string

		// Vela service for authentication.
		Admin          *AdminService
		Authentication *AuthenticationService
		Authorization  *AuthorizationService
		Build          *BuildService
		Dashboard      *DashboardService
		Deployment     *DeploymentService
		Hook           *HookService
		Log            *LogService
		Pipeline       *PipelineService
		Repo           *RepoService
		SCM            *SCMService
		Schedule       *ScheduleService
		Secret         *SecretService
		Step           *StepService
		Svc            *SvcService
		User           *UserService
		Worker         *WorkerService
		Queue          *QueueService
	}

	service struct {
		client *Client
	}

	// ListOptions represents the optional parameters to various List methods that
	// support pagination.
	ListOptions struct {
		// For paginated result sets, page of results to retrieve.
		Page int `url:"page,omitempty"`

		// For paginated result sets, the number of results to include per page.
		PerPage int `url:"per_page,omitempty"`
	}

	// OAuthExchangeOptions represents the required
	// parameters to exchange for tokens.
	OAuthExchangeOptions struct {
		Code  string `url:"code,omitempty"`
		State string `url:"state,omitempty"`
	}

	// LoginOptions represents the optional parameters
	// to launch the login process.
	LoginOptions struct {
		Type string `url:"type,omitempty"`
		Port string `url:"port,omitempty"`
	}
)

// NewClient returns a new Vela API client.
// baseURL has to be the HTTP endpoint of the Vela API.
// If no httpClient is provided, then the http.DefaultClient will be used.
func NewClient(baseURL, id string, httpClient *http.Client) (*Client, error) {
	// use http.DefaultClient if no client is provided
	if httpClient == nil {
		httpClient = http.DefaultClient
		httpClient.Timeout = time.Second * 15
	}

	// we must have a url provided to create the client
	if len(baseURL) == 0 {
		return nil, fmt.Errorf("no Vela baseURL provided")
	}

	// parse url provided for the client
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// prepare the user agent string
	ua := fmt.Sprintf("%s/%s", userAgent, version.Version.String())

	// if an ID was given, use it in the user agent string
	if len(id) > 0 {
		ua = fmt.Sprintf("%s (%s)", ua, id)
	}

	// create initial client fields
	c := &Client{
		client:    httpClient,
		baseURL:   url,
		UserAgent: ua,
	}

	// instantiate all client services
	c.Authentication = &AuthenticationService{client: c}
	c.Authorization = &AuthorizationService{client: c}
	c.Admin = &AdminService{
		&AdminBuildService{client: c},
		&AdminCleanService{client: c},
		&AdminDeploymentService{client: c},
		&AdminHookService{client: c},
		&AdminOIDCService{client: c},
		&AdminRepoService{client: c},
		&AdminSecretService{client: c},
		&AdminSvcService{client: c},
		&AdminStepService{client: c},
		&AdminUserService{client: c},
		&AdminWorkerService{client: c},
		&AdminSettingsService{client: c},
	}
	c.Build = &BuildService{client: c}
	c.Dashboard = &DashboardService{client: c}
	c.Deployment = &DeploymentService{client: c}
	c.Hook = &HookService{client: c}
	c.Log = &LogService{client: c}
	c.Pipeline = &PipelineService{client: c}
	c.Repo = &RepoService{client: c}
	c.SCM = &SCMService{client: c}
	c.Schedule = &ScheduleService{client: c}
	c.Secret = &SecretService{client: c}
	c.Step = &StepService{client: c}
	c.Svc = &SvcService{client: c}
	c.User = &UserService{client: c}
	c.Worker = &WorkerService{client: c}
	c.Queue = &QueueService{client: c}

	return c, nil
}

// SetTimeout sets the timeout for the http client.
func (c *Client) SetTimeout(d time.Duration) {
	c.client.Timeout = d
}

// buildURLForRequest will build the URL (as a string) that will be called.
// It does several cleaning tasks for us.
func (c *Client) buildURLForRequest(urlStr string) (string, error) {
	// capture base url from client for string
	u := c.baseURL.String()

	// If there is no / at the end, add one.
	if !strings.HasSuffix(u, "/") {
		u += "/"
	}

	// remove "/" prefix from url
	urlStr = strings.TrimPrefix(urlStr, "/")

	// parse trimmed url string
	rel, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	u += rel.String()

	return u, nil
}

// addAuthentication adds the necessary authentication to the request.
func (c *Client) addAuthentication(ctx context.Context, req *http.Request) error {
	// token that will be sent with the request depending on auth type
	token := ""

	// handle access + refresh tokens
	// refresh access token if needed
	if c.Authentication.HasAccessAndRefreshAuth() {
		currentAccess, err := c.Authentication.getAccessToken()
		if err != nil {
			return err
		}

		currentRefresh, err := c.Authentication.getRefreshToken()
		if err != nil {
			return err
		}

		isExpired := IsTokenExpired(currentAccess)
		if isExpired {
			logrus.Debug("access token has expired")

			isRefreshExpired := IsTokenExpired(currentRefresh)
			if isRefreshExpired {
				return fmt.Errorf("your tokens have expired - please log in again with 'vela login'")
			}

			logrus.Debug("fetching new access token with existing refresh token")

			// send API call to refresh the access token to Vela
			//
			// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#AuthenticationService.RefreshAccessToken
			_, err = c.Authentication.RefreshAccessToken(ctx, currentRefresh)
			if err != nil {
				return err
			}
		}

		// refresh could have produced new access token
		updatedAccess, err := c.Authentication.getAccessToken()
		if err != nil {
			return err
		}

		token = updatedAccess
	}

	// handle personal access token
	if c.Authentication.HasPersonalAccessTokenAuth() {
		// send API call to exchange token for access token to Vela
		//
		// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#AuthenticationService.AuthenticateWithToken
		at, _, err := c.Authentication.AuthenticateWithToken(ctx, *c.Authentication.personalAccessToken)
		if err != nil {
			return err
		}

		token = at
	}

	// handle plain token
	if c.Authentication.HasTokenAuth() {
		token = *c.Authentication.token
	}

	if c.Authentication.HasBuildTokenAuth() {
		token = *c.Authentication.token

		scmTkn := *c.Authentication.scmToken
		if len(scmTkn) == 0 {
			return fmt.Errorf("scm token has no value")
		}

		if c.Authentication.IsSCMTokenExpired() {
			splitR := strings.Split(*c.Authentication.buildRepo, "/")
			if len(splitR) != 2 {
				return fmt.Errorf("invalid build repo format")
			}

			org := splitR[0]
			repo := splitR[1]
			build := *c.Authentication.buildNumber

			_, err := c.Authentication.RefreshInstallToken(ctx, org, repo, build)
			if err != nil {
				return err
			}
		}

		req.Header.Add("Token", *c.Authentication.scmToken)
	}

	// make sure token is not empty
	if len(token) == 0 {
		return fmt.Errorf("token has no value")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return nil
}

// addOptions adds the parameters in opt as url query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	// return url if option is a pointer but is also nil
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	// parse url provided for the options
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	// add query values to url
	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	// safely encode url with query values
	u.RawQuery = qs.Encode()

	return u.String(), nil
}

// NewRequest creates an API request.
// A relative URL can be provided in url,
// in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(ctx context.Context, method, url string, body any) (*http.Request, error) {
	// build url for request
	u, err := c.buildURLForRequest(url)
	if err != nil {
		return nil, err
	}

	// variable to store http request
	var req *http.Request

	// handle body based on body type
	switch body := body.(type) {
	// io.ReadCloser is used for streaming endpoints
	case io.ReadCloser:
		req, err = http.NewRequestWithContext(ctx, method, u, body)
		if err != nil {
			return nil, err
		}
	// default assumes JSON body
	default:
		var buf io.ReadWriter
		if body != nil {
			// buffer to store request body
			buf = new(bytes.Buffer)

			// encode request body into buffer for request
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}

		// create new http request from built url and body
		req, err = http.NewRequestWithContext(ctx, method, u, buf)
		if err != nil {
			return nil, err
		}
	}

	// apply authentication to request if client is set
	if c.Authentication.HasAuth() {
		err = c.addAuthentication(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	// add the user agent for the request
	req.Header.Add("User-Agent", c.UserAgent)

	// apply default header for content-type
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// Response represents an Vela API response.
// This wraps the standard http.Response returned from Vela.
type Response struct {
	*http.Response

	// Values hold basic information pertaining to how to paginate
	// through response results
	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.populatePageValues()

	return response
}

// populatePageValues parses the HTTP Link response headers and populates the
// various pagination link values in the Response.
func (r *Response) populatePageValues() {
	if links, ok := r.Header["Link"]; ok && len(links) > 0 {
		for _, link := range strings.Split(links[0], ",") {
			segments := strings.Split(strings.TrimSpace(link), ";")

			// link must at least have href and rel
			if len(segments) < 2 {
				continue
			}

			// ensure href is properly formatted
			if !strings.HasPrefix(segments[0], "<") || !strings.HasSuffix(segments[0], ">") {
				continue
			}

			// try to pull out page parameter
			url, err := url.Parse(segments[0][1 : len(segments[0])-1])
			if err != nil {
				continue
			}

			page := url.Query().Get("page")
			if page == "" {
				continue
			}

			for _, segment := range segments[1:] {
				switch strings.TrimSpace(segment) {
				case `rel="next"`:
					r.NextPage, _ = strconv.Atoi(page)
				case `rel="prev"`:
					r.PrevPage, _ = strconv.Atoi(page)
				case `rel="first"`:
					r.FirstPage, _ = strconv.Atoi(page)
				case `rel="last"`:
					r.LastPage, _ = strconv.Atoi(page)
				}
			}
		}
	}
}

// Call is a combined function for Client.NewRequest and Client.Do.
//
// Most API methods are quite the same.
// Get the URL, apply options, make a request, and get the response.
// Without adding special headers or something.
// To avoid a big amount of code duplication you can Client.Call.
//
// method is the HTTP method you want to call.
// url is the URL you want to call.
// body is the HTTP body.
// respType is the type that the HTTP response will resolve to.
//
// For more information read https://github.com/google/go-github/issues/234
func (c *Client) Call(ctx context.Context, method, url string, body, respType interface{}) (*Response, error) {
	// create new request from parameters
	req, err := c.NewRequest(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	// send request with client
	resp, err := c.Do(req, respType)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// CallWithHeaders is a combined function for Client.NewRequest and Client.Do.
//
// Most API methods are quite the same.
// Get the URL, apply options, make a request, and get the response.
// Without adding special headers or something.
// To avoid a big amount of code duplication you can Client.Call.
//
// method is the HTTP method you want to call.
// url is the URL you want to call.
// body is the HTTP body.
// respType is the type that the HTTP response will resolve to.
// headers is a map of HTTP headers.
//
// For more information read https://github.com/google/go-github/issues/234
func (c *Client) CallWithHeaders(ctx context.Context, method, url string, body, respType interface{}, headers map[string]string) (*Response, error) {
	// create new request from parameters
	req, err := c.NewRequest(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	// add header key or overwrite key with new values
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// send request with client
	resp, err := c.Do(req, respType)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by respType,
// or returned as an error if an API error has occurred.
// If respType implements the io.Writer interface, the raw response body will
// be written to respType, without attempting to first decode it.
func (c *Client) Do(req *http.Request, respType interface{}) (*Response, error) {
	// send request with client
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer closing response body
	defer resp.Body.Close()

	// wrap response
	response := &Response{Response: resp}

	// check response for errors
	err = CheckResponse(resp)
	if err != nil {
		// if error is present, we still return the response so the caller
		// may inspect it further for debugging and troubleshooting
		return response, err
	}

	// if return object is provided
	if respType != nil {
		// copy response body if object implements io.Writer interface
		if w, ok := respType.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return response, err
			}
		} else {
			// copy all bytes from response body
			body, err := io.ReadAll(resp.Body)
			// ensure response body is not empty so the user may inspect
			// it further for debugging and troubleshooting
			resp.Body = io.NopCloser(bytes.NewBuffer(body))

			if err != nil {
				// if error is present, we still return the response so the caller
				// may inspect it further for debugging and troubleshooting
				return response, err
			}

			// check if the content type is YAML (or deprecated x-yaml)
			if strings.Contains(resp.Header.Get("Content-Type"), "application/yaml") ||
				strings.Contains(resp.Header.Get("Content-Type"), "application/x-yaml") {
				// unmarshal the body as YAML to the return object
				_ = yaml.Unmarshal(body, respType)
			} else {
				// unmarshal the body as JSON to the return object
				_ = json.Unmarshal(body, respType)
			}
		}
	}

	return response, err
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
func CheckResponse(r *http.Response) error {
	// return no error if successful response code
	if c := r.StatusCode; http.StatusOK <= c && c <= 299 {
		return nil
	}

	// custom response type
	resp := api.Error{}

	// read all bytes from response body
	b, _ := io.ReadAll(r.Body)

	// unmarshal bytes into custom response type
	err := json.Unmarshal(b, &resp)
	if err != nil {
		//nolint:nilerr // ignore returning nil
		return nil
	}

	return fmt.Errorf("%v", *resp.Message)
}
