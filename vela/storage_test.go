package vela

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
	"github.com/google/go-cmp/cmp"
)

func TestStorage_GetInfo_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)
	data := []byte(server.StorageInfoResp)
	c.Authentication.SetPersonalAccessTokenAuth("token")

	var want *api.StorageInfo

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Storage.GetInfo(t.Context())
	if err != nil {
		t.Errorf("GetInfo returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("GetInfo returned %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetInfo mismatch (-want +got):\n%s", diff)
	}
}

func TestStorage_GetInfo_401(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Storage.GetInfo(t.Context())
	if err == nil {
		t.Errorf("GetInfo should have returned err %v", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("GetInfo returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}
