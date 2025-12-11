// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/go-vela/server/mock/server"
)

func TestSCM_Sync_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.SCM.Sync(t.Context(), "github", "octocat")
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestSCM_Sync_404(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.SCM.Sync(t.Context(), "github", "not-found")
	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}
}

func TestSCM_SyncAll_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.SCM.SyncAll(t.Context(), "github")
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestSCM_SyncAll_404(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.SCM.SyncAll(t.Context(), "not-found")
	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}
}

func ExampleSCMService_Sync() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Change orgship of the repo in the server
	repo, resp, err := c.SCM.Sync(context.Background(), "github", "octocat")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, repo)
}

func ExampleSCMService_SyncAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Change orgship of the repo in the server
	org, resp, err := c.SCM.SyncAll(context.Background(), "github")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for repo %+v", resp.StatusCode, org)
}
