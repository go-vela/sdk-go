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
	"github.com/go-vela/server/api/types/actions"
	"github.com/go-vela/server/mock/server"
)

func TestRepo_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.RepoResp)

	var want api.Repo
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Repo.Get("github", "octocat")

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo get is %v, want %v", got, want)
	}
}

func TestRepo_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Repo{}

	// run test
	got, resp, err := c.Repo.Get("github", "not-found")

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo get is %v, want %v", got, want)
	}
}

func TestRepo_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ReposResp)

	var want []api.Repo
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Repo.GetAll(nil)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo getall is %v, want %v", got, want)
	}
}

func TestRepo_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.RepoResp)

	var want api.Repo
	_ = json.Unmarshal(data, &want)

	req := api.Repo{
		Org:         String("github"),
		Name:        String("octocat"),
		FullName:    String("github/octocat"),
		Link:        String("https://github.com/github/octocat"),
		Clone:       String("https://github.com/github/octocat.git"),
		Branch:      String("main"),
		Timeout:     Ptr(60),
		Visibility:  String("public"),
		Private:     Bool(false),
		Trusted:     Bool(false),
		Active:      Bool(true),
		AllowEvents: testEvents(),
	}

	// run test
	got, resp, err := c.Repo.Add(&req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo add is %v, want %v", got, want)
	}
}

func TestRepo_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.RepoResp)

	var want api.Repo
	_ = json.Unmarshal(data, &want)

	req := api.Repo{
		Private:     Bool(true),
		Trusted:     Bool(true),
		Active:      Bool(true),
		AllowEvents: testEvents(),
	}

	// run test
	got, resp, err := c.Repo.Update("github", "octocat", &req)

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

func TestRepo_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Repo{}

	req := api.Repo{
		Private:     Bool(true),
		Trusted:     Bool(true),
		Active:      Bool(true),
		AllowEvents: testEvents(),
	}

	// run test
	got, resp, err := c.Repo.Update("github", "not-found", &req)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo update is %v, want %v", got, want)
	}
}

func TestRepo_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Repo.Remove("github", "octocat")

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestRepo_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Repo.Remove("github", "not-found")

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestRepo_Repair_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Repo.Repair("github", "octocat")

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestRepo_Repair_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Repo.Repair("github", "not-found")

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestRepo_Chown_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Repo.Chown("github", "octocat")

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestRepo_Chown_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Repo.Chown("github", "not-found")

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}
}

func ExampleRepoService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a repo from the server
	repo, resp, err := c.Repo.Get("github", "octocat")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, repo)
}

func ExampleRepoService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the repos from the server
	repos, resp, err := c.Repo.GetAll(nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repos %+v", resp.StatusCode, repos)
}

func ExampleRepoService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Repo{
		Org:         String("github"),
		Name:        String("octocat"),
		FullName:    String("github/octocat"),
		Link:        String("https://github.com/github/octocat"),
		Clone:       String("https://github.com/github/octocat.git"),
		Branch:      String("main"),
		Timeout:     Ptr(60),
		Visibility:  String("public"),
		Private:     Bool(false),
		Trusted:     Bool(false),
		Active:      Bool(true),
		AllowEvents: testEvents(),
	}

	// Create the repo in the server
	repo, resp, err := c.Repo.Add(&req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, repo)
}

func ExampleRepoService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Repo{
		AllowEvents: testEvents(),
	}

	// Update the repo in the server
	repo, resp, err := c.Repo.Update("github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, repo)
}

func ExampleRepoService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the repo in the server
	repo, resp, err := c.Repo.Remove("github", "octocat")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, repo)
}

func ExampleRepoService_Repair() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Repair the repo in the server
	repo, resp, err := c.Repo.Repair("github", "octocat")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, repo)
}

func ExampleRepoService_Chown() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Change orgship of the repo in the server
	repo, resp, err := c.Repo.Chown("github", "octocat")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, repo)
}

func testEvents() *api.Events {
	return &api.Events{
		Push: &actions.Push{
			Branch:       Bool(true),
			Tag:          Bool(true),
			DeleteBranch: Bool(true),
			DeleteTag:    Bool(true),
		},
		PullRequest: &actions.Pull{
			Opened:      Bool(true),
			Edited:      Bool(true),
			Synchronize: Bool(true),
			Reopened:    Bool(true),
		},
		Deployment: &actions.Deploy{
			Created: Bool(true),
		},
		Comment: &actions.Comment{
			Created: Bool(true),
			Edited:  Bool(true),
		},
		Schedule: &actions.Schedule{
			Run: Bool(true),
		},
	}
}
