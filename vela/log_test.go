// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
)

func TestLog_GetService_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.LogResp)

	var want api.Log
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Log.GetService("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Log get is %v, want %v", got, want)
	}
}

func TestLog_GetService_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Log{}

	// run test
	got, resp, err := c.Log.GetService("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Log get is %v, want %v", got, want)
	}
}

func TestLog_AddService_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.LogResp)

	var want api.Log
	_ = json.Unmarshal(data, &want)

	req := api.Log{
		Data: Bytes([]byte("Hello, World")),
	}

	// run test
	resp, err := c.Log.AddService("github", "octocat", 1, 1, &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_UpdateService_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	req := api.Log{
		Data: Bytes([]byte("Hello, World Manny")),
	}

	// run test
	resp, err := c.Log.UpdateService("github", "octocat", 1, 1, &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_UpdateService_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	req := api.Log{
		Data: Bytes([]byte("Hello, World Manny")),
	}

	// run test
	resp, err := c.Log.UpdateService("github", "not-found", 1, 0, &req)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_RemoveService_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Log.RemoveService("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_RemoveService_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Log.RemoveService("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_GetStep_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.LogResp)

	var want api.Log
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Log.GetStep("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Log get is %v, want %v", got, want)
	}
}

func TestLog_GetStep_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Log{}

	// run test
	got, resp, err := c.Log.GetStep("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Log get is %v, want %v", got, want)
	}
}

func TestLog_AddStep_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.LogResp)

	var want api.Log
	_ = json.Unmarshal(data, &want)

	req := api.Log{
		Data: Bytes([]byte("Hello, World")),
	}

	// run test
	resp, err := c.Log.AddStep("github", "octocat", 1, 1, &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_UpdateStep_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	req := api.Log{
		Data: Bytes([]byte("Hello, World Manny")),
	}

	// run test
	resp, err := c.Log.UpdateStep("github", "octocat", 1, 1, &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_UpdateStep_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	req := api.Log{
		Data: Bytes([]byte("Hello, World Manny")),
	}

	// run test
	resp, err := c.Log.UpdateStep("github", "not-found", 1, 0, &req)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_RemoveStep_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Log.RemoveStep("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestLog_RemoveStep_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Log.RemoveStep("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Log returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func ExampleLogService_GetService() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a log from the server
	log, resp, err := c.Log.GetService("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for log %+v", resp.StatusCode, log)
}

func ExampleLogService_AddService() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Log{
		Data: Bytes([]byte("Hello World")),
	}

	// Create the log in the server
	resp, err := c.Log.AddService("github", "octocat", 1, 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d for log", resp.StatusCode)
}

func ExampleLogService_UpdateService() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Log{
		Data: Bytes([]byte("Hello World")),
	}

	// Update the log in the server
	resp, err := c.Log.UpdateService("github", "octocat", 1, 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for log", resp.StatusCode)
}

func ExampleLogService_RemoveService() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the log in the server
	log, resp, err := c.Log.RemoveService("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for log %+v", resp.StatusCode, log)
}

func ExampleLogService_GetStep() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a log from the server
	log, resp, err := c.Log.GetStep("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for log %+v", resp.StatusCode, log)
}

func ExampleLogService_AddStep() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Log{
		Data: Bytes([]byte("Hello World")),
	}

	// Create the log in the server
	resp, err := c.Log.AddStep("github", "octocat", 1, 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d for log", resp.StatusCode)
}

func ExampleLogService_UpdateStep() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Log{
		Data: Bytes([]byte("Hello World")),
	}

	// Update the log in the server
	resp, err := c.Log.UpdateStep("github", "octocat", 1, 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d for log", resp.StatusCode)
}

func ExampleLogService_RemoveStep() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the log in the server
	log, resp, err := c.Log.RemoveStep("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for log %+v", resp.StatusCode, log)
}
