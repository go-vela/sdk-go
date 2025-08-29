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

func TestDashboard_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DashCardResp)

	var want api.DashCard

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Errorf("unable to unmarshal data: %v", err)
	}

	// run test
	got, resp, err := c.Dashboard.Get(t.Context(), "c976470d-34c1-49b2-9a98-1035871c576b")
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

func TestDashboard_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.DashCard{}

	// run test
	got, resp, err := c.Dashboard.Get(t.Context(), "0")
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

func TestDashboard_GetAllUser_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DashCardsResp)

	var want []api.DashCard

	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Dashboard.GetAllUser(t.Context())
	if err != nil {
		t.Errorf("GetAllUser returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GetAllUser returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetAllUser is %v, want %v", got, want)
	}
}

func TestDashboard_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DashboardResp)

	var want api.Dashboard

	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Dashboard.Add(t.Context(), &want)
	if err != nil {
		t.Errorf("Add returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Add returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Add is %v, want %v", got, want)
	}
}

func TestDashboard_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DashboardResp)

	var want api.Dashboard

	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Dashboard.Update(t.Context(), &want)
	if err != nil {
		t.Errorf("Update returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Update returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Update is %v, want %v", got, want)
	}
}
