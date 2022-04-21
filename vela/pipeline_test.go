// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types/yaml"

	yml "github.com/buildkite/yaml"
)

func TestPipeline_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.PipelineResp)

	var want yaml.Build
	_ = yml.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Pipeline.Get("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err != nil {
		t.Errorf("Get returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Get returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Get is %v, want %v", got, want)
	}
}

func TestPipeline_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := yaml.Build{}

	// run test
	got, resp, err := c.Pipeline.Get("github", "not-found", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err == nil {
		t.Errorf("Get returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Get returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Get is %v, want %v", got, want)
	}
}

func TestPipeline_Compile_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.CompileResp)

	var want yaml.Build
	_ = yml.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Pipeline.Compile("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err != nil {
		t.Errorf("Compile returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Compile returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Compile is %v, want %v", got, want)
	}
}

func TestPipeline_Compile_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := yaml.Build{}

	// run test
	got, resp, err := c.Pipeline.Compile("github", "not-found", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err == nil {
		t.Errorf("Compile returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Compile returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Compile is %v, want %v", got, want)
	}
}

func TestPipeline_Expand_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ExpandResp)

	var want yaml.Build
	_ = yml.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Pipeline.Expand("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err != nil {
		t.Errorf("Expand returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expand returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Expand is %v, want %v", got, want)
	}
}

func TestPipeline_Expand_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := yaml.Build{}

	// run test
	got, resp, err := c.Pipeline.Expand("github", "not-found", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err == nil {
		t.Errorf("Expand returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expand returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Expand is %v, want %v", got, want)
	}
}

func TestPipeline_Templates_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.TemplateResp)

	want := make(map[string]*yaml.Template)
	_ = yml.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Pipeline.Templates("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err != nil {
		t.Errorf("Templates returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Templates returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Templates is %v, want %v", got, want)
	}
}

func TestPipeline_Templates_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := make(map[string]*yaml.Template)

	// run test
	got, resp, err := c.Pipeline.Templates("github", "not-found", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err == nil {
		t.Errorf("Templates returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Templates returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Templates is %v, want %v", got, want)
	}
}

func TestPipeline_Validate_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Pipeline.Validate("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Validate returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestPipeline_Validate_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Pipeline.Validate("github", "not-found", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)

	if err == nil {
		t.Errorf("Validate returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Validate returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func ExamplePipelineService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// create options for pipeline call
	opts := &PipelineOptions{
		Output: "yaml", // default
	}

	// get a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Get("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_Compile() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// create options for pipeline call
	opts := &PipelineOptions{
		Output: "yaml", // default
	}

	// compile a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Compile("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_Expand() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// create options for pipeline call
	opts := &PipelineOptions{
		Output: "yaml", // default
	}

	// expand templates for a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Expand("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_Templates() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// create options for pipeline call
	opts := &PipelineOptions{
		Output: "yaml", // default
	}

	// get templates for a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Templates("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_Validate() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// create options for pipeline call
	opts := &PipelineOptions{
		Output:   "yaml", // default
		Template: true,   // default
	}

	// get templates for a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Validate("github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}
