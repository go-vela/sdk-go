// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
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

	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types/library"

	"github.com/gin-gonic/gin"
)

func TestHook_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.HookResp)

	var want library.Hook
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Hook.Get("github", "octocat", 1)

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

func TestHook_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := library.Hook{}

	// run test
	got, resp, err := c.Hook.Get("github", "octocat", 0)

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

func TestHook_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.HooksResp)

	var want []library.Hook
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Hook.GetAll("github", "octocat", nil)

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

func TestHook_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.HookResp)

	var want library.Hook
	_ = json.Unmarshal(data, &want)

	req := library.Hook{
		Number:   Int(1),
		SourceID: String("c8da1302-07d6-11ea-882f-4893bca275b8"),
		Event:    String("push"),
		Status:   String("created"),
		Error:    String(""),
		Created:  Int64(1563474076),
		Link:     String("https://github.com/github/octocat/settings/hooks/1"),
		Branch:   String("master"),
		Host:     String("github.com"),
	}

	// run test
	got, resp, err := c.Hook.Add("github", "octocat", &req)

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

func TestHook_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.HookResp)

	var want library.Hook
	_ = json.Unmarshal(data, &want)

	req := library.Hook{
		Number: Int(1),
		Event:  String("push"),
		Status: String("success"),
	}

	// run test
	got, resp, err := c.Hook.Update("github", "octocat", &req)

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

func TestHook_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := library.Hook{}

	req := library.Hook{
		Number: Int(0),
		Event:  String("push"),
		Status: String("running"),
	}

	// run test
	got, resp, err := c.Hook.Update("github", "octocat", &req)

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

func TestHook_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Hook.Remove("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Remove returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestHook_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Hook.Remove("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Remove returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func ExampleHookService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a hook from the server
	hook, resp, err := c.Hook.Get("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for hook %+v", resp.StatusCode, hook)
}

func ExampleHookService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the hooks from the server
	hooks, resp, err := c.Hook.GetAll("github", "octocat", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for hooks %+v", resp.StatusCode, hooks)
}

func ExampleHookService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := library.Hook{
		Number:   Int(1),
		SourceID: String("c8da1302-07d6-11ea-882f-4893bca275b8"),
		Event:    String("push"),
		Status:   String("created"),
		Error:    String(""),
		Created:  Int64(1563474076),
		Link:     String("https://github.com/github/octocat/settings/hooks/1"),
		Branch:   String("master"),
		Host:     String("github.com"),
	}

	// Create the hook in the server
	hook, resp, err := c.Hook.Add("github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for hook %+v", resp.StatusCode, hook)
}

func ExampleHookService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := library.Hook{
		Status: String("error"),
		Error:  String(""),
	}

	// Update the step in the server
	hook, resp, err := c.Hook.Update("github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for hook %+v", resp.StatusCode, hook)
}

func ExampleHookService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the hook in the server
	hook, resp, err := c.Hook.Remove("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, hook)
}
