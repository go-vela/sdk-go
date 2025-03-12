package vela

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStorage_GetInfo_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)
	c.Authentication.SetPersonalAccessTokenAuth("token")
	data := []byte(server.StorageInfoResp)

	var want *api.StorageInfo

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Storage.GetInfo()
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
	_, resp, err := c.Storage.GetInfo()
	if err == nil {
		t.Errorf("GetInfo should have returned err %v", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("GetInfo returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}
