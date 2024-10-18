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

func TestService_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ServiceResp)

	var want api.Service
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Svc.Get("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Service returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Service get is %v, want %v", got, want)
	}
}

func TestService_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Service{}

	// run test
	got, resp, err := c.Svc.Get("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Service returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Service get is %v, want %v", got, want)
	}
}

func TestService_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ServicesResp)

	var want []api.Service
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Svc.GetAll("github", "octocat", 1, nil)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Service returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Service getall is %v, want %v", got, want)
	}
}

func TestService_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ServiceResp)

	var want api.Service
	_ = json.Unmarshal(data, &want)

	req := api.Service{
		Number:   Int(1),
		Name:     String("clone"),
		Status:   String("created"),
		Error:    String(""),
		ExitCode: Int(0),
		Created:  Int64(1563475419),
		Started:  Int64(0),
		Finished: Int64(0),
	}

	// run test
	got, resp, err := c.Svc.Add("github", "octocat", 1, &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Service returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Service create is %v, want %v", got, want)
	}
}

func TestService_Update_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ServiceResp)

	var want api.Service
	_ = json.Unmarshal(data, &want)

	req := api.Service{
		Number:   Int(1),
		Status:   String("finished"),
		Started:  Int64(1563475419),
		Finished: Int64(1563475419),
	}

	// run test
	got, resp, err := c.Svc.Update("github", "octocat", 1, &req)

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

func TestService_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Service{}

	req := api.Service{
		Number:   Int(0),
		Status:   String("finished"),
		Started:  Int64(1563475419),
		Finished: Int64(1563475419),
	}

	// run test
	got, resp, err := c.Svc.Update("github", "not-found", 0, &req)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Service get is %v, want %v", got, want)
	}
}

func TestService_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Svc.Remove("github", "octocat", 1, 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Service returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestService_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Svc.Remove("github", "octocat", 1, 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func ExampleSvcService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a service from the server
	service, resp, err := c.Svc.Get("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for service %+v", resp.StatusCode, service)
}

func ExampleSvcService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the services from the server
	services, resp, err := c.Svc.GetAll("github", "octocat", 1, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for services %+v", resp.StatusCode, services)
}

func ExampleSvcService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Service{
		Number:   Int(1),
		Name:     String("clone"),
		Status:   String("pending"),
		Error:    String(""),
		ExitCode: Int(0),
		Created:  Int64(time.Now().UTC().Unix()),
		Started:  Int64(0),
		Finished: Int64(0),
	}

	// Create the service in the server
	service, resp, err := c.Svc.Add("github", "octocat", 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for service %+v", resp.StatusCode, service)
}

func ExampleSvcService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Service{
		Status: String("error"),
		Error:  String("Something in the runtime broke"),
	}

	// Update the service in the server
	service, resp, err := c.Svc.Update("github", "octocat", 1, &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for service %+v", resp.StatusCode, service)
}

func ExampleSvcService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the service in the server
	service, resp, err := c.Svc.Remove("github", "octocat", 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for service %+v", resp.StatusCode, service)
}
