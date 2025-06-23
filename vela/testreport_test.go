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

func TestReport_Add_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)
	c.Authentication.SetPersonalAccessTokenAuth("token")
	data := []byte(server.TestReportResp)

	var want *api.TestReport

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.TestReport.Add("org", "repo", 1)
	if err != nil {
		t.Errorf("Add returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Add returned %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Add mismatch (-want +got):\n%s", diff)
	}
}

//func TestReport_Add_401(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//
//	s := httptest.NewServer(server.FakeHandler())
//	c, _ := NewClient(s.URL, "", nil)
//
//	// run test
//	_, resp, err := c.Storage.GetInfo()
//	if err == nil {
//		t.Errorf("GetInfo should have returned err %v", resp.StatusCode)
//	}
//	if resp.StatusCode != http.StatusUnauthorized {
//		t.Errorf("GetInfo returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
//	}
//}
