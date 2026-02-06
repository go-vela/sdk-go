package vela

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
)

func TestStorage_GetSTSCreds_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)
	data := []byte(server.StorageSTSResp)
	c.Authentication.SetPersonalAccessTokenAuth("token")

	var want *api.STSCreds

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Build.GetSTSCreds(t.Context(), "foo", "bar", 1)
	if err != nil {
		t.Errorf("GetSTSCreds returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("GetSTSCreds returned %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetSTSCreds mismatch (-want +got):\n%s", diff)
	}
}

func TestStorage_GetSTSCreds_401(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Build.GetSTSCreds(t.Context(), "foo", "bar", 1)
	if err == nil {
		t.Errorf("GetSTSCreds should have returned err %v", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("GetSTSCreds returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}
