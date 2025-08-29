// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	yml "go.yaml.in/yaml/v3"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/compiler/types/yaml/yaml"
	"github.com/go-vela/server/mock/server"
)

func TestPipeline_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.PipelineResp)

	var want api.Pipeline

	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Pipeline.Get(t.Context(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163")
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

	want := api.Pipeline{}

	// run test
	got, resp, err := c.Pipeline.Get(t.Context(), "github", "octocat", "0")
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

func TestPipeline_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.PipelinesResp)

	var want []api.Pipeline

	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Pipeline.GetAll(t.Context(), "github", "octocat", nil)
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GetAll returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetAll is %v, want %v", got, want)
	}
}

func TestPipeline_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.PipelineResp)

	var want api.Pipeline

	_ = json.Unmarshal(data, &want)

	req := api.Pipeline{
		Commit:  String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Ref:     String("refs/heads/main"),
		Type:    String("yaml"),
		Version: String("1"),
		Steps:   Bool(true),
	}

	// run test
	got, resp, err := c.Pipeline.Add(t.Context(), "github", "octocat", &req)
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Add returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Add add is %v, want %v", got, want)
	}
}

func TestPipeline_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.PipelineResp)

	var want api.Pipeline

	_ = json.Unmarshal(data, &want)

	req := api.Pipeline{
		Commit: String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Type:   String("yaml"),
	}

	// run test
	got, resp, err := c.Pipeline.Update(t.Context(), "github", "octocat", &req)
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Update returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Update is %v, want %v", got, want)
	}
}

func TestPipeline_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Pipeline{}

	req := api.Pipeline{
		Commit: String("0"),
	}

	// run test
	got, resp, err := c.Pipeline.Update(t.Context(), "github", "octocat", &req)
	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Update returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Update is %v, want %v", got, want)
	}
}

func TestPipeline_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Pipeline.Remove(t.Context(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163")
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Remove returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestPipeline_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Pipeline.Remove(t.Context(), "github", "octocat", "0")
	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Remove returned %v, want %v", resp.StatusCode, http.StatusOK)
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
	got, resp, err := c.Pipeline.Compile(t.Context(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)
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
	got, resp, err := c.Pipeline.Compile(t.Context(), "github", "octocat", "0", nil)
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
	got, resp, err := c.Pipeline.Expand(t.Context(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)
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
	got, resp, err := c.Pipeline.Expand(t.Context(), "github", "octocat", "0", nil)
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
	got, resp, err := c.Pipeline.Templates(t.Context(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)
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
	got, resp, err := c.Pipeline.Templates(t.Context(), "github", "octocat", "0", nil)
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
	_, resp, err := c.Pipeline.Validate(t.Context(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", nil)
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
	_, resp, err := c.Pipeline.Validate(t.Context(), "github", "octocat", "0", nil)
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

	// get a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Get(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the pipelines from the server
	pipelines, resp, err := c.Pipeline.GetAll(context.Background(), "github", "octocat", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipelines %+v", resp.StatusCode, pipelines)
}

func ExamplePipelineService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Pipeline{
		Commit:  String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Ref:     String("refs/heads/main"),
		Type:    String("yaml"),
		Version: String("1"),
		Steps:   Bool(true),
	}

	// Create the pipeline in the server
	pipeline, resp, err := c.Pipeline.Add(context.Background(), "github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Pipeline{
		Commit: String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Type:   String("yaml"),
	}

	// Update the step in the server
	pipeline, resp, err := c.Pipeline.Update(context.Background(), "github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the pipeline in the server
	pipeline, resp, err := c.Pipeline.Remove(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, pipeline)
}

func ExamplePipelineService_Compile() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// create default options for pipeline call
	opts := &PipelineOptions{
		Output: "yaml", // default
	}

	// compile a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Compile(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)

	// create ruledata options for pipeline call
	opts = &PipelineOptions{
		Output:  "yaml", // default
		Branch:  "main",
		Comment: "comment",
		Event:   "push",
		Repo:    "octocat",
		Status:  "success",
		Tag:     "v1.0.0",
		Target:  "production",
		Path:    []string{"path/to/file", "README.md"},
	}

	// compile a pipeline from a repo from the server
	pipeline, resp, err = c.Pipeline.Compile(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
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
	pipeline, resp, err := c.Pipeline.Expand(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)

	// create ruledata options for pipeline call
	opts = &PipelineOptions{
		Output:  "yaml", // default
		Branch:  "main",
		Comment: "comment",
		Event:   "push",
		Repo:    "octocat",
		Status:  "success",
		Tag:     "v1.0.0",
		Target:  "production",
		Path:    []string{"path/to/file", "README.md"},
	}

	// compile a pipeline from a repo from the server
	pipeline, resp, err = c.Pipeline.Expand(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
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
	pipeline, resp, err := c.Pipeline.Templates(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)

	// create ruledata options for pipeline call
	opts = &PipelineOptions{
		Output:  "yaml", // default
		Branch:  "main",
		Comment: "comment",
		Event:   "push",
		Repo:    "octocat",
		Status:  "success",
		Tag:     "v1.0.0",
		Target:  "production",
		Path:    []string{"path/to/file", "README.md"},
	}

	// compile a pipeline from a repo from the server
	pipeline, resp, err = c.Pipeline.Templates(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
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
		Output: "yaml", // default
	}

	// get templates for a pipeline from a repo from the server
	pipeline, resp, err := c.Pipeline.Validate(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)

	// create ruledata options for pipeline call
	opts = &PipelineOptions{
		Output:  "yaml", // default
		Branch:  "main",
		Comment: "comment",
		Event:   "push",
		Repo:    "octocat",
		Status:  "success",
		Tag:     "v1.0.0",
		Target:  "production",
		Path:    []string{"path/to/file", "README.md"},
	}

	// compile a pipeline from a repo from the server
	pipeline, resp, err = c.Pipeline.Validate(context.Background(), "github", "octocat", "48afb5bdc41ad69bf22588491333f7cf71135163", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for pipeline %+v", resp.StatusCode, pipeline)
}
