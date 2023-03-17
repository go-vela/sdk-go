// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-vela/sdk-go/version"
	"github.com/go-vela/server/mock/server"
)

func TestVela_NewClient(t *testing.T) {
	// setup types
	addr := "http://localhost:8080"

	url, err := url.Parse(addr)
	if err != nil {
		t.Errorf("Unable to parse url: %v", err)
	}

	want := &Client{
		client:    http.DefaultClient,
		baseURL:   url,
		UserAgent: fmt.Sprintf("%s/%s", "vela-sdk-go", version.Version.String()),
	}
	want.Authentication = &AuthenticationService{client: want}
	want.Authorization = &AuthorizationService{client: want}
	want.Admin = &AdminService{
		&AdminBuildService{client: want},
		&AdminDeploymentService{client: want},
		&AdminHookService{client: want},
		&AdminRepoService{client: want},
		&AdminSecretService{client: want},
		&AdminSvcService{client: want},
		&AdminStepService{client: want},
		&AdminUserService{client: want},
		&AdminWorkerService{client: want},
	}
	want.Build = &BuildService{client: want}
	want.Deployment = &DeploymentService{client: want}
	want.Hook = &HookService{client: want}
	want.Log = &LogService{client: want}
	want.Pipeline = &PipelineService{client: want}
	want.Repo = &RepoService{client: want}
	want.SCM = &SCMService{client: want}
	want.Secret = &SecretService{client: want}
	want.Step = &StepService{client: want}
	want.Svc = &SvcService{client: want}
	want.Worker = &WorkerService{client: want}

	// run test
	got, err := NewClient(addr, "", nil)
	if err != nil {
		t.Errorf("NewClient returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewClient is %v, want %v", got, want)
	}
}

func TestVela_NewClient_EmptyUrl(t *testing.T) {
	// run test
	got, err := NewClient("", "", nil)
	if err == nil {
		t.Errorf("NewClient should have returned err")
	}

	if got != nil {
		t.Errorf("NewClient is %v, want nil", got)
	}
}

func TestVela_NewClient_UserAgent(t *testing.T) {
	// setup types
	addr := "http://localhost:8080"

	want := fmt.Sprintf("%s/%s (%s)", userAgent, version.Version.String(), "vela")

	// run test
	got, err := NewClient(addr, "vela", nil)
	if err != nil {
		t.Errorf("NewClient returned err: %v", err)
	}

	if got.UserAgent != want {
		t.Errorf("NewClient is %v, want %v", got, want)
	}
}

func TestVela_NewClient_BadUrl(t *testing.T) {
	// run test
	got, err := NewClient("!@#$%^&*()", "", nil)
	if err == nil {
		t.Errorf("NewClient should have returned err")
	}

	if got != nil {
		t.Errorf("NewClient is %v, want nil", got)
	}
}

func TestVela_SetTimeout(t *testing.T) {
	// setup types
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	tests := []struct {
		input time.Duration
		want  time.Duration
	}{
		// use the default timeout
		{
			want: 15 * time.Second,
		},
		// set a custom timeout
		{
			input: 73 * time.Minute,
			want:  73 * time.Minute,
		},
	}

	for _, tc := range tests {
		// if not using the default timeout, then set custom timeout
		if tc.input != 0 {
			t.Log(tc.input)
			c.SetTimeout(tc.input)
		}

		got := c.client.Timeout

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("SetTimeout is %v, want %v", got, tc.want)
		}
	}
}

func TestVela_buildURLForRequest_NoSlash(t *testing.T) {
	// setup types
	want := "http://localhost:8080/test"

	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("test")
	if err != nil {
		t.Errorf("buildURLForRequest returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildURLForRequest is %v, want %v", got, want)
	}
}

func TestVela_buildURLForRequest_PrefixSlash(t *testing.T) {
	// setup types
	want := "http://localhost:8080/test"

	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("/test")
	if err != nil {
		t.Errorf("buildURLForRequest returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildURLForRequest is %v, want %v", got, want)
	}
}

func TestVela_buildURLForRequest_SuffixSlash(t *testing.T) {
	// setup types
	want := "http://localhost:8080/test/"

	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("test/")
	if err != nil {
		t.Errorf("buildURLForRequest returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("buildURLForRequest is %v, want %v", got, want)
	}
}

func TestVela_buildURLForRequest_BadUrl(t *testing.T) {
	// setup types
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	got, err := c.buildURLForRequest("!@#$%^&*()")
	if err == nil {
		t.Errorf("buildURLForRequest should have returned err")
	}

	if len(got) > 0 {
		t.Errorf("buildURLForRequest is %v, want \"\"", got)
	}
}

func TestVela_addAuthentication(t *testing.T) {
	// setup types
	want := "Bearer foobar"

	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	r, err := http.NewRequestWithContext(context.Background(), "GET", "http://localhost:8080/health", nil)
	if err != nil {
		t.Errorf("Unable to create new request: %v", err)
	}

	// run test
	c.Authentication.SetTokenAuth("foobar")

	err = c.addAuthentication(r)
	if err != nil {
		t.Error("addAuthentication should not have errored")
	}

	got := r.Header.Get("Authorization")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addAuthentication is %v, want %v", got, want)
	}
}

func TestVela_addAuthentication_AccessAndRefresh_GoodToken(t *testing.T) {
	// setup types
	testToken := TestTokenGood
	want := fmt.Sprintf("Bearer %s", testToken)

	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	r, err := http.NewRequestWithContext(context.Background(), "GET", "http://localhost:8080/health", nil)
	if err != nil {
		t.Errorf("Unable to create new request: %v", err)
	}

	// run test
	c.Authentication.SetAccessAndRefreshAuth(testToken, "bar")

	err = c.addAuthentication(r)
	if err != nil {
		t.Error("addAuthentication should not have errored")
	}

	got := r.Header.Get("Authorization")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addAuthentication is %v, want %v", got, want)
	}
}

func TestVela_addAuthentication_AccessAndRefresh_ExpiredTokens(t *testing.T) {
	// setup types
	testToken := TestTokenExpired

	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	r, err := http.NewRequestWithContext(context.Background(), "GET", "http://localhost:8080/health", nil)
	if err != nil {
		t.Errorf("Unable to create new request: %v", err)
	}

	// run test
	c.Authentication.SetAccessAndRefreshAuth(testToken, testToken)

	err = c.addAuthentication(r)
	if err == nil {
		t.Error("addAuthentication should have errored with expired tokens")
	}
}

func TestVela_addAuthentication_AccessAndRefresh_ExpiredAccessGoodRefresh(t *testing.T) {
	// setup types
	want := "Bearer header.payload.signature"
	s := httptest.NewServer(server.FakeHandler())

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	r, err := http.NewRequestWithContext(context.Background(), "GET", fmt.Sprintf("%s/health", s.URL), nil)
	if err != nil {
		t.Errorf("Unable to create new request: %v", err)
	}

	// run test
	c.Authentication.SetAccessAndRefreshAuth(TestTokenExpired, TestTokenGood)

	err = c.addAuthentication(r)
	if err != nil {
		t.Error("addAuthentication should not have errored")
	}

	got := r.Header.Get("Authorization")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addAuthentication is %v, want %v", got, want)
	}
}

func TestVela_Call_BadMethod(t *testing.T) {
	// setup types
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	// run test
	_, err = c.Call("!@#$%^&*()", "/health", nil, nil)
	if err == nil {
		t.Errorf("Call should have returned err")
	}
}

func TestClient_CallWithHeaders(t *testing.T) {
	type args struct {
		method  string
		u       string
		body    interface{}
		v       interface{}
		headers map[string]string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"happy path",
			args{"GET", "/health", nil, nil, nil},
			false,
		},
		{
			"custom header",
			args{"GET", "/health", nil, nil, map[string]string{"Content-Type": "application/octet-stream"}},
			false,
		},
		{
			"bad method",
			args{"$(#*@&$", "/health", nil, nil, nil},
			true,
		},
	}

	s := httptest.NewServer(server.FakeHandler())

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.CallWithHeaders(tt.args.method, tt.args.u, tt.args.body, tt.args.v, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CallWithHeaders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestVela_NewRequest(t *testing.T) {
	rc := io.NopCloser(strings.NewReader("Hello, world!"))

	type input struct {
		method   string
		endpoint string
		body     interface{}
	}

	tests := []struct {
		name    string
		input   input
		failure bool
		want    *http.Request
	}{
		{
			name:    "happy path",
			input:   input{method: "GET", endpoint: "/health", body: nil},
			failure: false,
		},
		{
			name:    "bad method",
			input:   input{method: "!@#$%^&*()", endpoint: "/health", body: nil},
			failure: true,
		},
		{
			name:    "bad endpoint",
			input:   input{method: "GET", endpoint: "!@#$%^&*()", body: nil},
			failure: true,
		},
		{
			name:    "stream",
			input:   input{method: "GET", endpoint: "/health", body: rc},
			failure: false,
		},
		{
			name:    "stream bad method",
			input:   input{method: "!@#$%^&*()", endpoint: "/health", body: rc},
			failure: true,
		},
	}

	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		t.Errorf("Unable to create new client: %v", err)
	}

	c.Authentication.SetTokenAuth("foobar")

	// run test
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup types
			if !test.failure {
				if test.input.body == nil {
					test.want, err = http.NewRequest(
						test.input.method,
						fmt.Sprintf("http://localhost:8080%s", test.input.endpoint),
						nil,
					)
				} else {
					test.want, err = http.NewRequest(
						test.input.method,
						fmt.Sprintf("http://localhost:8080%s", test.input.endpoint),
						test.input.body.(io.ReadCloser),
					)
				}
				if err != nil {
					t.Errorf("Unable to create new request: %v", err)
				}
				test.want.Header.Add("Content-Type", "application/json")
				test.want.Header.Add("Authorization", "Bearer foobar")
				test.want.Header.Add("User-Agent", c.UserAgent)
			}

			got, err := c.NewRequest(test.input.method, test.input.endpoint, test.input.body)

			if test.failure {
				if err == nil {
					t.Errorf("NewRequest should have returned err")
				}

				return
			}

			if err != nil {
				t.Errorf("NewRequest returned err: %v", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewRequest is %v, want %v", got, test.want)
			}
		})
	}
}

type options struct {
	ShowAll bool `url:"all"`
	Page    int  `url:"page"`
}

func TestVela_addOptions(t *testing.T) {
	// setup types
	want := "http://localhost:8080?all=true&page=1"
	options := options{ShowAll: true, Page: 1}

	// run test
	got, err := addOptions("http://localhost:8080", options)
	if err != nil {
		t.Errorf("addOptions returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}

func TestVela_addOptions_BadOptions(t *testing.T) {
	// setup types
	want := "http://localhost:8080"

	// run test
	got, err := addOptions("http://localhost:8080", 87)
	if err == nil {
		t.Errorf("addOptions should have returned err")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}

func TestVela_addOptions_BadUrl(t *testing.T) {
	// setup types
	want := "!@#$%^&*()"
	options := options{ShowAll: true, Page: 1}

	// run test
	got, err := addOptions("!@#$%^&*()", options)
	if err == nil {
		t.Errorf("addOptions should have returned err")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}

func TestVela_addOptions_NilOptions(t *testing.T) {
	// setup types
	want := "http://localhost:8080"

	// run test
	got, err := addOptions("http://localhost:8080", nil)
	if err != nil {
		t.Errorf("addOptions returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}

func TestResponse_populatePageValues(t *testing.T) {
	// setup types
	r := http.Response{
		Header: http.Header{
			"Link": {`<https://vela.company.com/api/v1/repos?per_page=1&page=1>; rel="first",` +
				` <https://vela.company.com/api/v1/repos?per_page=1&page=2>; rel="prev",` +
				` <https://vela.company.com/api/v1/repos?per_page=1&page=4>; rel="next",` +
				` <https://vela.company.com/api/v1/repos?per_page=1&page=5>; rel="last"`,
			},
		},
	}

	// run test
	response := newResponse(&r)
	if got, want := response.FirstPage, 1; got != want {
		t.Errorf("response.FirstPage: %v, want %v", got, want)
	}

	if got, want := response.PrevPage, 2; want != got {
		t.Errorf("response.PrevPage: %v, want %v", got, want)
	}

	if got, want := response.NextPage, 4; want != got {
		t.Errorf("response.NextPage: %v, want %v", got, want)
	}

	if got, want := response.LastPage, 5; want != got {
		t.Errorf("response.LastPage: %v, want %v", got, want)
	}
}

func TestResponse_populatePageValues_invalid(t *testing.T) {
	// setup types
	r := http.Response{
		Header: http.Header{
			"Link": {`<https://vela.company.com/api/v1/repos/?page=1>,` +
				`<https://vela.company.com/api/v1/repos/?page=foo>; rel="first",` +
				`https://vela.company.com/api/v1/repos/?page=1; rel="prev",` +
				`<https://vela.company.com/api/v1/repos/>; rel="next",` +
				`<https://vela.company.com/api/v1/repos/?page=>; rel="last"`,
			},
		},
	}

	// run test
	response := newResponse(&r)
	if got, want := response.FirstPage, 0; got != want {
		t.Errorf("response.FirstPage: %v, want %v", got, want)
	}

	if got, want := response.PrevPage, 0; got != want {
		t.Errorf("response.PrevPage: %v, want %v", got, want)
	}

	if got, want := response.NextPage, 0; got != want {
		t.Errorf("response.NextPage: %v, want %v", got, want)
	}

	if got, want := response.LastPage, 0; got != want {
		t.Errorf("response.LastPage: %v, want %v", got, want)
	}
}
