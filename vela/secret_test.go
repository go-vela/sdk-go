// SPDX-License-Identifier: Apache-2.0

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

func TestSecret_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SecretResp)

	var want library.Secret
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Secret.Get("native", "repo", "github", "octocat", "foo")

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret get is %v, want %v", got, want)
	}
}

func TestSecret_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := library.Secret{}

	// run test
	got, resp, err := c.Secret.Get("native", "repo", "github", "not-found", "not-found")

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret get is %v, want %v", got, want)
	}
}

func TestSecret_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SecretsResp)

	var want []library.Secret
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Secret.GetAll("native", "repo", "github", "octocat", nil)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret get is %v, want %v", got, want)
	}
}

func TestSecret_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SecretResp)

	var want library.Secret
	_ = json.Unmarshal(data, &want)

	req := library.Secret{
		Org:         String("github"),
		Repo:        String("octocat"),
		Name:        String("foo"),
		Value:       String("bar"),
		Images:      &[]string{"foo", "bar"},
		AllowEvents: testLibraryEvents(),
	}

	// run test
	got, resp, err := c.Secret.Add("native", "repo", "github", "octocat", &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret add is %v, want %v", got, want)
	}
}

func TestSecret_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SecretResp)

	var want library.Secret
	_ = json.Unmarshal(data, &want)

	req := library.Secret{
		Name:        String("foo"),
		Value:       String("bar"),
		AllowEvents: testLibraryEvents(),
	}

	// run test
	got, resp, err := c.Secret.Update("native", "repo", "github", "octocat", &req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret get is %v, want %v", got, want)
	}
}

func TestSecret_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := library.Secret{}

	req := library.Secret{
		Name:        String("foo"),
		Value:       String("bar"),
		AllowEvents: testLibraryEvents(),
	}

	// run test
	got, resp, err := c.Secret.Update("native", "repo", "github", "not-found", &req)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret get is %v, want %v", got, want)
	}
}

func TestSecret_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Secret.Remove("native", "repo", "github", "octocat", "foo")

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestSecret_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Secret.Remove("native", "repo", "github", "not-found", "not-found")

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func ExampleSecretService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get the secret from the server
	secret, resp, err := c.Secret.Get("native", "repo", "github", "octocat", "foo")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for secret %+v", resp.StatusCode, secret)
}

func ExampleSecretService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the secrets from the server
	secrets, resp, err := c.Secret.GetAll("native", "repo", "github", "octocat", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for secrets %+v", resp.StatusCode, secrets)
}

func ExampleSecretService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := library.Secret{
		Name:        String("foo"),
		Value:       String("bar"),
		Images:      &[]string{"foo", "bar"},
		AllowEvents: testLibraryEvents(),
	}

	// Create the secret in the server
	secret, resp, err := c.Secret.Add("native", "repo", "github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for secret %+v", resp.StatusCode, secret)
}

func ExampleSecretService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := library.Secret{
		Name:        String("foo"),
		Value:       String("bar"),
		AllowEvents: testLibraryEvents(),
	}

	// Update the secret in the server
	secret, resp, err := c.Secret.Update("native", "repo", "github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for secret %+v", resp.StatusCode, secret)
}

func ExampleSecretService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the secret in the server
	secret, resp, err := c.Secret.Remove("native", "repo", "github", "octocat", "foo")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for secret %+v", resp.StatusCode, secret)
}
