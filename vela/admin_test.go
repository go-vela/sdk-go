// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/api/types/actions"
	"github.com/go-vela/server/api/types/settings"
	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types"
	"github.com/go-vela/types/library"
)

func TestAdmin_Build_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildResp)

	var want api.Build
	_ = json.Unmarshal(data, &want)

	req := api.Build{
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

func TestAdmin_Clean_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := server.CleanResourcesResp

	req := types.Error{
		Message: String("msg"),
	}

	// run test
	got, resp, err := c.Admin.Clean.Clean(&req, nil)

	if err != nil {
		t.Errorf("Clean returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Clean returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Clean is %v, want %v", got, want)
	}
}

func TestAdmin_Clean_Error(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	req := types.Error{
		Message: String("msg"),
	}

	opt500 := CleanOptions{
		Before: 1,
	}

	opt401 := CleanOptions{
		Before: 2,
	}

	// run tests
	_, resp, _ := c.Admin.Clean.Clean(&req, &opt500)

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Clean returned %v, want %v", resp.StatusCode, http.StatusInternalServerError)
	}

	_, resp, _ = c.Admin.Clean.Clean(&req, &opt401)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Clean returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}

func TestAdmin_Deployment_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.DeploymentResp)

	var want api.Deployment
	_ = json.Unmarshal(data, &want)

	req := api.Deployment{
		Commit:      String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Ref:         String("refs/heads/main"),
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

func TestAdmin_Hook_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.HookResp)

	var want api.Hook
	_ = json.Unmarshal(data, &want)

	req := api.Hook{
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

func TestAdmin_Repo_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.RepoResp)

	var want api.Repo
	_ = json.Unmarshal(data, &want)

	req := api.Repo{
		Private: Bool(true),
		Trusted: Bool(true),
		Active:  Bool(true),
		AllowEvents: &api.Events{
			Push: &actions.Push{
				Branch:       Bool(true),
				Tag:          Bool(true),
				DeleteBranch: Bool(true),
				DeleteTag:    Bool(true),
			},
			PullRequest: &actions.Pull{
				Opened:      Bool(true),
				Edited:      Bool(true),
				Synchronize: Bool(true),
				Reopened:    Bool(true),
			},
			Deployment: &actions.Deploy{
				Created: Bool(true),
			},
			Comment: &actions.Comment{
				Created: Bool(true),
				Edited:  Bool(true),
			},
			Schedule: &actions.Schedule{
				Run: Bool(true),
			},
		},
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

func TestAdmin_Secret_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SecretResp)

	var want api.Secret
	_ = json.Unmarshal(data, &want)

	req := api.Secret{
		Name:        String("foo"),
		Value:       String("bar"),
		AllowEvents: testEvents(),
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

func TestAdmin_Service_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.ServiceResp)

	var want api.Service
	_ = json.Unmarshal(data, &want)

	req := api.Service{
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

func TestAdmin_Step_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.StepResp)

	var want api.Step
	_ = json.Unmarshal(data, &want)

	req := api.Step{
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

func TestAdmin_User_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UserResp)

	var want api.User
	_ = json.Unmarshal(data, &want)

	req := api.User{
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

	var want *[]api.QueueBuild

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

func TestAdmin_Worker_RegistrationToken_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.RegisterTokenResp)

	var want *library.Token

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	hostname := "foo"

	// run test
	got, resp, err := c.Admin.Worker.RegisterToken(hostname)
	if err != nil {
		t.Errorf("RegisterToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("RegisterToken returned %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("RegisterToken() mismatch (-want +got):\n%s", diff)
	}
}

func TestAdmin_Worker_RegistrationToken_NoHostname(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// bad hostname
	hostname := ""

	// run test
	_, _, err := c.Admin.Worker.RegisterToken(hostname)
	if err == nil {
		t.Error("RegisterToken should have returned err")
	}
}

func TestAdmin_Settings_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.SettingsResp)

	var want *settings.Platform

	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Error(err)
	}

	// run test
	got, resp, err := c.Admin.Settings.Get()
	if err != nil {
		t.Errorf("Settings.Get returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Settings.Get returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Settings.Get() mismatch (-want +got):\n%s", diff)
	}
}

func TestAdmin_Settings_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.UpdateSettingsResp)

	var want settings.Platform
	_ = json.Unmarshal(data, &want)

	req := settings.Platform{}

	// run test
	got, resp, err := c.Admin.Settings.Update(&req)

	if err != nil {
		t.Errorf("Settings.Update returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Settings.Update returned response code %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Settings.Update returned %v, want %v", got, want)
	}
}

func TestAdmin_Settings_Restore_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.RestoreSettingsResp)

	var want settings.Platform
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Admin.Settings.Restore()

	if err != nil {
		t.Errorf("Settings.Restore returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Settings.Restore returned response code %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Settings.Restore returned %v, want %v", got, want)
	}
}

func TestAdmin_OIDC_RotateKeys_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := "keys rotated successfully"

	// run test
	got, resp, err := c.Admin.OIDC.RotateOIDCKeys()
	if err != nil {
		t.Errorf("RotateOIDCKeys returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("RotateOIDCKeys returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if diff := cmp.Diff(&want, got); diff != "" {
		t.Errorf("RotateOIDCKeys() mismatch (-want +got):\n%s", diff)
	}
}

func TestAdmin_OIDC_RotateKeys_Unauthorized(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	c.Authentication.SetTokenAuth("invalid")

	// run test
	_, resp, err := c.Admin.OIDC.RotateOIDCKeys()
	if err == nil {
		t.Error("RotateOIDCKeys should have returned err")
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("RotateOIDCKeys returned %v, want %v", resp.StatusCode, http.StatusUnauthorized)
	}
}
