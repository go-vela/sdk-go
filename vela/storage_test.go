// SPDX-License-Identifier: Apache-2.0

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

func TestStorage_GetPresignedPutURL_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)
	data := []byte(server.PresignedPutResp)

	c.Authentication.SetBuildTokenAuth("buildToken", "scmToken", 0, "foo/bar", 1)

	var want *api.PresignURL

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Build.GetPresignedPutURL(t.Context(), "file.txt", "foo", "bar", 1)
	if err != nil {
		t.Errorf("GetPresignedPutURL returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("GetPresignedPutURL returned %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetPresignedPutURL mismatch (-want +got):\n%s", diff)
	}
}

func TestStorage_GetPresignedPutURL_401(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Build.GetPresignedPutURL(t.Context(), "file.txt", "foo", "bar", 1)
	if err == nil {
		t.Errorf("GetPresignedPutURL should have returned err %v", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("GetPresignedPutURL returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}
