// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"encoding/json"
	"github.com/go-vela/server/database"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/mock/server"
	"github.com/go-vela/types/library"
)

func TestAdmin_Build_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildsResp)

	var want []library.Build
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Build.GetAll(nil)

	if err != nil {
		t.Errorf("Build returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build getall is %v, want %v", got, want)
	}
}

func TestAdmin_Build_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildResp)

	var want library.Build
	_ = json.Unmarshal(data, &want)

	req := library.Build{
		Number: Int(1),
		Parent: Int(1),
		Event:  String("push"),
		Status: String("running"),
	}

	// run test
	got, resp, err := c.Admin.Build.Update(&req)

	if err != nil {
		t.Errorf("Build returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build update is %v, want %v", got, want)
	}
}

func TestAdmin_Deployment_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DeploymentsResp)

	var want []library.Deployment
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Deployment.GetAll(nil)

	if err != nil {
		t.Errorf("Deployment returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Deployment returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Deployment getall is %v, want %v", got, want)
	}
}

func TestAdmin_Deployment_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DeploymentResp)

	var want library.Deployment
	_ = json.Unmarshal(data, &want)

	req := library.Deployment{
		Commit:      String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Ref:         String("refs/heads/master"),
		Task:        String("vela-deploy"),
		Target:      String("production"),
		Description: String("Deployment request from Vela"),
	}

	// run test
	got, resp, err := c.Admin.Deployment.Update(&req)

	if err != nil {
		t.Errorf("Deployment returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Deployment returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Deployment update is %v, want %v", got, want)
	}
}

func TestAdmin_Hook_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.HooksResp)

	var want []library.Hook
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Hook.GetAll(nil)

	if err != nil {
		t.Errorf("Hook returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Hook returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Hook getall is %v, want %v", got, want)
	}
}

func TestAdmin_Hook_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.HookResp)

	var want library.Hook
	_ = json.Unmarshal(data, &want)

	req := library.Hook{
		Number: Int(1),
		Event:  String("push"),
		Status: String("success"),
	}

	// run test
	got, resp, err := c.Admin.Hook.Update(&req)

	if err != nil {
		t.Errorf("Hook returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Hook returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Hook update is %v, want %v", got, want)
	}
}

func TestAdmin_Repo_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ReposResp)

	var want []library.Repo
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Repo.GetAll(nil)

	if err != nil {
		t.Errorf("Repo returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo getall is %v, want %v", got, want)
	}
}

func TestAdmin_Repo_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.RepoResp)

	var want library.Repo
	_ = json.Unmarshal(data, &want)

	req := library.Repo{
		Private:     Bool(true),
		Trusted:     Bool(true),
		Active:      Bool(true),
		AllowPull:   Bool(true),
		AllowPush:   Bool(true),
		AllowDeploy: Bool(true),
		AllowTag:    Bool(true),
	}

	// run test
	got, resp, err := c.Admin.Repo.Update(&req)

	if err != nil {
		t.Errorf("Repo returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Repo returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Repo update is %v, want %v", got, want)
	}
}

func TestAdmin_Secret_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SecretsResp)

	var want []library.Secret
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Secret.GetAll(nil)

	if err != nil {
		t.Errorf("Secret returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret getall is %v, want %v", got, want)
	}
}

func TestAdmin_Secret_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SecretResp)

	var want library.Secret
	_ = json.Unmarshal(data, &want)

	req := library.Secret{
		Name:   String("foo"),
		Value:  String("bar"),
		Events: &[]string{"barf", "foob"},
	}
	// run test
	got, resp, err := c.Admin.Secret.Update(&req)

	if err != nil {
		t.Errorf("Secret returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Secret returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Secret update is %v, want %v", got, want)
	}
}

func TestAdmin_Service_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ServicesResp)

	var want []library.Service
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Service.GetAll(nil)

	if err != nil {
		t.Errorf("Service returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Service returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Service getall is %v, want %v", got, want)
	}
}

func TestAdmin_Service_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ServiceResp)

	var want library.Service
	_ = json.Unmarshal(data, &want)

	req := library.Service{
		Number:   Int(1),
		Status:   String("finished"),
		Started:  Int64(1563475419),
		Finished: Int64(1563475419),
	}

	// run test
	got, resp, err := c.Admin.Service.Update(&req)

	if err != nil {
		t.Errorf("Service returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Service returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Service update is %v, want %v", got, want)
	}
}

func TestAdmin_Step_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.StepsResp)

	var want []library.Step
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Step.GetAll(nil)

	if err != nil {
		t.Errorf("Step returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Step returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Step getall is %v, want %v", got, want)
	}
}

func TestAdmin_Step_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.StepResp)

	var want library.Step
	_ = json.Unmarshal(data, &want)

	req := library.Step{
		Number:   Int(1),
		Status:   String("finished"),
		Started:  Int64(1563475419),
		Finished: Int64(1563475419),
	}

	// run test
	got, resp, err := c.Admin.Step.Update(&req)

	if err != nil {
		t.Errorf("Step returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Step returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Step update is %v, want %v", got, want)
	}
}

func TestAdmin_User_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UsersResp)

	var want []library.User
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.User.GetAll(nil)

	if err != nil {
		t.Errorf("User returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("User returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User getall is %v, want %v", got, want)
	}
}

func TestAdmin_User_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UserResp)

	var want library.User
	_ = json.Unmarshal(data, &want)

	req := library.User{
		Name: String("octocat"),
	}

	// run test
	got, resp, err := c.Admin.User.Update(&req)

	if err != nil {
		t.Errorf("User returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("User returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("User update is %v, want %v", got, want)
	}
}

func TestAdmin_Build_Queue_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildQueueResp)

	var want *[]database.BuildQueue
	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Admin.Build.GetQueue(&GetQueueOptions{})
	if err != nil {
		t.Errorf("GetQueue returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GetQueue returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetQueue() mismatch (-want +got):\n%s", diff)
	}
}
