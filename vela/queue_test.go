// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types/library"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdmin_Worker_GetQueueCreds_200(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)
	c.Authentication.SetPersonalAccessTokenAuth("token")
	data := []byte(server.QueueRegistrationResp)

	var want *library.QueueRegistration

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Queue.GetQueueCreds()
	if err != nil {
		t.Errorf("GetQueueCreds returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("GetQueueCreds returned %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetQueueCreds() mismatch (-want +got):\n%s", diff)
	}
}

func TestAdmin_Worker_GetQueueCreds_401(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Queue.GetQueueCreds()
	if err == nil {
		t.Errorf("getQueueCreds should have returned err %v", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("getQueueCreds returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}
