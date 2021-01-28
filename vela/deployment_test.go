// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-vela/mock/server"
	"github.com/go-vela/types/library"

	"github.com/gin-gonic/gin"
)

func TestDeployment_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DeploymentResp)

	var want library.Deployment
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Deployment.Get("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Get returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Get is %v, want %v", got, want)
	}
}

func TestDeployment_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := library.Deployment{}

	// run test
	got, resp, err := c.Deployment.Get("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Get returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Get is %v, want %v", got, want)
	}
}

func TestDeployment_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DeploymentsResp)

	var want []library.Deployment
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Deployment.GetAll("github", "octocat", nil)

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

func TestDeployment_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DeploymentResp)

	var want library.Deployment
	_ = json.Unmarshal(data, &want)

	req := library.Deployment{
		Commit:      String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Ref:         String("refs/heads/master"),
		Task:        String("vela-deploy"),
		Target:      String("production"),
		Description: String("Deployment request from Vela"),
	}

	// run test
	got, resp, err := c.Deployment.Add("github", "octocat", &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Add returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Add is %v, want %v", got, want)
	}
}

func ExampleDeploymentService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a build from the server
	deployment, resp, err := c.Deployment.Get("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for deployment %+v", resp.StatusCode, deployment)
}

func ExampleDeploymentService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the deployments from the server
	deployments, resp, err := c.Deployment.GetAll("github", "octocat", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for deployments %+v", resp.StatusCode, deployments)
}

func ExampleDeploymentService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := library.Deployment{
		Commit:      String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Ref:         String("refs/heads/master"),
		Task:        String("vela-deploy"),
		Target:      String("production"),
		Description: String("Deployment request from Vela"),
	}

	// Create the deployment in the server
	deployment, resp, err := c.Deployment.Add("github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for deployment %+v", resp.StatusCode, deployment)
}
