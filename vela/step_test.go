// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
)

func TestStep_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.StepResp)

	var want api.Step
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Step.Get("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Step returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Step get is %v, want %v", got, want)
	}
}

func TestStep_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Step{}

	// run test
	got, resp, err := c.Step.Get("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Step returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Step get is %v, want %v", got, want)
	}
}

func TestStep_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.StepsResp)

	var want []api.Step
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Step.GetAll("github", "octocat", 1, nil)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Step returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Step getall is %v, want %v", got, want)
	}
}

func TestStep_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.StepResp)

	var want api.Step
	_ = json.Unmarshal(data, &want)

	req := api.Step{
		Number:       Int32(1),
		Name:         String("clone"),
		Status:       String("created"),
		Error:        String(""),
		ExitCode:     Int32(0),
		Created:      Int64(1563475419),
		Started:      Int64(0),
		Finished:     Int64(0),
		Host:         String("example.company.com"),
		Runtime:      String("docker"),
		Distribution: String("linux"),
	}

	// run test
	got, resp, err := c.Step.Add("github", "octocat", 1, &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Step returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Step create is %v, want %v", got, want)
	}
}

func TestStep_Update_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.StepResp)

	var want api.Step
	_ = json.Unmarshal(data, &want)

	req := api.Step{
		Number:   Int32(1),
		Status:   String("finished"),
		Started:  Int64(1563475419),
		Finished: Int64(1563475419),
	}

	// run test
	got, resp, err := c.Step.Update("github", "octocat", 1, &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo create is %v, want %v", got, want)
	}
}

func TestStep_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Step{}

	req := api.Step{
		Number:   Int32(0),
		Status:   String("finished"),
		Started:  Int64(1563475419),
		Finished: Int64(1563475419),
	}

	// run test
	got, resp, err := c.Step.Update("github", "not-found", 0, &req)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Step get is %v, want %v", got, want)
	}
}

func TestStep_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Step.Remove("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Step returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestStep_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Step.Remove("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func ExampleStepService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a step from the server
	step, resp, err := c.Step.Get("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, step)
}

func ExampleStepService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the steps from the server
	steps, resp, err := c.Step.GetAll("github", "octocat", 1, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for steps %+v", resp.StatusCode, steps)
}

func ExampleStepService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Step{
		Number:       Int32(1),
		Name:         String("clone"),
		Status:       String("pending"),
		Error:        String(""),
		ExitCode:     Int32(0),
		Created:      Int64(time.Now().UTC().Unix()),
		Started:      Int64(0),
		Finished:     Int64(0),
		Host:         String("example.company.com"),
		Runtime:      String("docker"),
		Distribution: String("linux"),
	}

	// Create the step in the server
	step, resp, err := c.Step.Add("github", "octocat", 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, step)
}

func ExampleStepService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Step{
		Status: String("error"),
		Error:  String("Something in the runtime broke"),
	}

	// Update the step in the server
	step, resp, err := c.Step.Update("github", "octocat", 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, step)
}

func ExampleStepService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the step in the server
	step, resp, err := c.Step.Remove("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, step)
}
