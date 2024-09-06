// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
)

func TestUser_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UserResp)

	var want api.User
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.User.Get("octocat")

	if err != nil {
		t.Errorf("User Get returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("User Get returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User Get is %v, want %v", got, want)
	}
}

func TestUser_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.User{}

	// run test
	got, resp, err := c.User.Get("not-found")

	if err == nil {
		t.Errorf("User Get should have returned err")
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("User Get returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User Get is %v, want %v", got, want)
	}
}

func TestUser_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UserResp)

	var want api.User
	_ = json.Unmarshal(data, &want)

	req := api.User{
		Admin: Bool(true),
	}

	// run test
	got, resp, err := c.User.Update("octocat", &req)

	if err != nil {
		t.Errorf("User Update returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("User Update returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User Update is %v, want %v", got, want)
	}
}

func TestUser_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.User{}

	req := api.User{
		Admin: Bool(true),
	}

	// run test
	got, resp, err := c.User.Update("not-found", &req)

	if err == nil {
		t.Errorf("User Update should have returned err")
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("User Update returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User Update is %v, want %v", got, want)
	}
}

func TestCurrentUser_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UserResp)

	var want api.User
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.User.GetCurrent()

	if err != nil {
		t.Errorf("User GetCurrent returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("User GetCurrent returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User GetCurrent is %v, want %v", got, want)
	}
}

func TestUser_GetCurrent_401(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	c.Authentication.SetTokenAuth("invalid")

	want := api.User{}

	// run test
	got, resp, err := c.User.GetCurrent()

	if err == nil {
		t.Errorf("User GetCurrent should have returned err")
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("User GetCurrent returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User GetCurrent is %v, want %v", got, want)
	}
}

func TestUser_UpdateCurrent_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UserResp)

	var want api.User
	_ = json.Unmarshal(data, &want)

	favorites := []string{"github/octocat"}

	req := api.User{
		Favorites: &favorites,
	}

	// run test
	got, resp, err := c.User.UpdateCurrent(&req)

	if err != nil {
		t.Errorf("User UpdateCurrent returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("User UpdateCurrent returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User UpdateCurrent is %v, want %v", got, want)
	}
}

func TestUser_UpdateCurrent_401(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	c.Authentication.SetTokenAuth("invalid")

	want := api.User{}

	favorites := []string{"github/octocat"}

	req := api.User{
		Favorites: &favorites,
	}

	// run test
	got, resp, err := c.User.UpdateCurrent(&req)

	if err == nil {
		t.Errorf("User UpdateCurrent should have returned err")
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("User UpdateCurrent returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User UpdateCurrent is %v, want %v", got, want)
	}
}
