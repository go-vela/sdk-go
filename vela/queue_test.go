// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types/library"
	"github.com/google/go-cmp/cmp"
)

func TestQueue_GetInfo_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)
	c.Authentication.SetPersonalAccessTokenAuth("token")
	data := []byte(server.QueueInfoResp)

	var want *library.QueueInfo

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Queue.GetInfo()
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

func TestQueue_GetInfo_401(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Queue.GetInfo()
	if err == nil {
		t.Errorf("GetInfo should have returned err %v", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("GetInfo returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}
