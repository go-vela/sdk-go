// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
)

func TestBuild_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildResp)

	var want api.Build
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.Get("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build get is %v, want %v", got, want)
	}
}

func TestBuild_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Build{}

	// run test
	got, resp, err := c.Build.Get("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build get is %v, want %v", got, want)
	}
}

func TestBuildExecutable_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildExecutableResp)

	var want api.BuildExecutable
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.GetBuildExecutable("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build get is %v, want %v", got, want)
	}
}

func TestBuildExecutable_Get_500(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.BuildExecutable{}

	// run test
	got, resp, err := c.Build.GetBuildExecutable("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusInternalServerError)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build get is %v, want %v", got, want)
	}
}

func TestBuild_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildsResp)

	var want []api.Build
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.GetAll("github", "octocat", nil)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build getall is %v, want %v", got, want)
	}
}

func TestBuild_GetLogs_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildLogsResp)

	var want []api.Log
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.GetLogs("github", "octocat", 1, nil)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build getlogs is %v, want %v", got, want)
	}
}

func TestBuild_GetLogs_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := []api.Log{}

	// run test
	got, resp, err := c.Build.GetLogs("github", "octocat", 0, nil)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if len(*got) != len(want) {
		t.Errorf("Build getlogs is %v, want %v", len(*got), len(want))
	}
}

func TestBuild_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildResp)

	var want api.Build
	_ = json.Unmarshal(data, &want)

	req := api.Build{
		Number: Int64(1),
		Repo: &api.Repo{
			Org:  String("github"),
			Name: String("octocat"),
		},
		Parent:       Int64(1),
		Event:        String("push"),
		Status:       String("created"),
		Error:        String(""),
		Enqueued:     Int64(1563474077),
		Created:      Int64(1563474076),
		Started:      Int64(1563474077),
		Finished:     Int64(0),
		Deploy:       String(""),
		Clone:        String("https://github.com/github/octocat.git"),
		Source:       String("https://github.com/github/octocat/abcdefghi123456789"),
		Title:        String("push received from https://github.com/github/octocat"),
		Message:      String("First commit..."),
		Commit:       String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Sender:       String("OctoKitty"),
		Author:       String("OctoKitty"),
		Email:        String("octokitty@github.com"),
		Link:         String("https://vela.example.company.com/github/octocat/1"),
		Branch:       String("main"),
		Ref:          String("refs/heads/main"),
		BaseRef:      String(""),
		Host:         String("example.company.com"),
		Runtime:      String("docker"),
		Distribution: String("linux"),
	}

	// run test
	got, resp, err := c.Build.Add(&req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Config returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build add is %v, want %v", got, want)
	}
}

func TestBuild_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildResp)

	var want api.Build
	_ = json.Unmarshal(data, &want)

	req := api.Build{
		Number: Int64(1),
		Repo: &api.Repo{
			Org:  String("github"),
			Name: String("octocat"),
		},
		Parent: Int64(1),
		Event:  String("push"),
		Status: String("running"),
	}

	// run test
	got, resp, err := c.Build.Update(&req)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Config returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build update is %v, want %v", got, want)
	}
}

func TestBuild_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Build{}

	req := api.Build{
		Number: Int64(0),
		Repo: &api.Repo{
			Org:  String("github"),
			Name: String("octocat"),
		},
		Parent: Int64(1),
		Event:  String("push"),
		Status: String("running"),
	}

	// run test
	got, resp, err := c.Build.Update(&req)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Config returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build update is %v, want %v", got, want)
	}
}

func TestBuild_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Build.Remove("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestBuild_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Build.Remove("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestBuild_Restart_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildResp)

	var want api.Build
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.Restart("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build restart is %v, want %v", got, want)
	}
}

func TestBuild_Restart_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := api.Build{}

	// run test
	got, resp, err := c.Build.Restart("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Build restart is %v, want %v", got, want)
	}
}

func TestBuild_Cancel_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, got, err := c.Build.Cancel("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if got.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", got.StatusCode, http.StatusOK)
	}
}

func TestBuild_Cancel_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Build.Cancel("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestBuild_Approve_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	got, err := c.Build.Approve("github", "octocat", 1)

	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	if got.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", got.StatusCode, http.StatusOK)
	}
}

func TestBuild_Approve_403(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	resp, err := c.Build.Approve("github", "octocat", 0)

	if err == nil {
		t.Errorf("New returned err: %v", err)
	}

	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestBuild_GetBuildToken_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.BuildTokenResp)

	var want api.Token
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.GetBuildToken("github", "octocat", 1)

	if err != nil {
		t.Errorf("GetBuildToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetBuildToken is %v, want %v", got, want)
	}
}

func TestBuild_GetBuildToken_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	var want api.Token

	// run test
	got, resp, err := c.Build.GetBuildToken("github", "octocat", 0)

	if err != nil {
		t.Errorf("GetBuildToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetBuildToken is %v, want %v", got, want)
	}
}

func TestBuild_GetBuildToken_400(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	var want api.Token

	// run test
	got, resp, err := c.Build.GetBuildToken("github", "octocat", 2)

	if err != nil {
		t.Errorf("GetBuildToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusBadRequest)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetBuildToken is %v, want %v", got, want)
	}
}

func TestBuild_GetIDRequestToken_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.IDTokenRequestTokenResp)

	var want api.Token
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.GetIDRequestToken("github", "octocat", 1, nil)

	if err != nil {
		t.Errorf("GetIDRequestToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetIDRequestToken is %v, want %v", got, want)
	}
}

func TestBuild_GetIDRequestToken_400(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	var want api.Token

	// run test
	got, resp, err := c.Build.GetIDRequestToken("github", "octocat", 0, nil)

	if err != nil {
		t.Errorf("GetIDRequestToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusBadRequest)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetIdRequestToken is %v, want %v", got, want)
	}
}

func TestBuild_GetIDToken_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.IDTokenResp)

	var want api.Token
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Build.GetIDToken("github", "octocat", 1, nil)

	if err != nil {
		t.Errorf("GetIDToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetIDToken is %v, want %v", got, want)
	}
}

func TestBuild_GetIDToken_400(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	var want api.Token

	// run test
	got, resp, err := c.Build.GetIDToken("github", "octocat", 0, nil)

	if err != nil {
		t.Errorf("GetIDToken returned err: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Build returned %v, want %v", resp.StatusCode, http.StatusBadRequest)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("GetIdToken is %v, want %v", got, want)
	}
}

func ExampleBuildService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a build from the server
	build, resp, err := c.Build.Get("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for build %+v", resp.StatusCode, build)
}

func ExampleBuildService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the builds from the server
	builds, resp, err := c.Build.GetAll("github", "octocat", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for builds %+v", resp.StatusCode, builds)
}

func ExampleBuildService_GetLogs() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get for a build from the server
	logs, resp, err := c.Build.GetLogs("github", "octocat", 1, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for build logs %+v", resp.StatusCode, logs)
}

func ExampleBuildService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Build{
		Number:       Int64(1),
		Repo:         &api.Repo{Org: String("github"), Name: String("octocat")},
		Parent:       Int64(1),
		Event:        String("push"),
		Status:       String("created"),
		Error:        String(""),
		Enqueued:     Int64(time.Now().UTC().Unix()),
		Created:      Int64(time.Now().UTC().Unix()),
		Started:      Int64(0),
		Finished:     Int64(0),
		Deploy:       String(""),
		Clone:        String("https://github.com/go-vela/server.git"),
		Source:       String("https://github.com/go-vela/server/abcdefghi123456789"),
		Title:        String(""),
		Message:      String(""),
		Commit:       String("abcdefghi123456789"),
		Sender:       String("someone"),
		Author:       String("someone"),
		Email:        String("someone@example.com"),
		Link:         String("https://vela.example.company.com/go-vela/server/1"),
		Branch:       String("main"),
		Ref:          String(""),
		BaseRef:      String(""),
		Host:         String("example.company.com"),
		Runtime:      String("docker"),
		Distribution: String("linux"),
	}

	// Create the build in the server
	build, resp, err := c.Build.Add(&req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for build %+v", resp.StatusCode, build)
}

func ExampleBuildService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Build{
		Repo:   &api.Repo{Org: String("github"), Name: String("octocat")},
		Status: String("error"),
		Error:  String(""),
	}

	// Update the step in the server
	build, resp, err := c.Build.Update(&req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for build %+v", resp.StatusCode, build)
}

func ExampleBuildService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the build in the server
	build, resp, err := c.Build.Remove("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, build)
}

func ExampleBuildService_Restart() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Restart the build in the server
	build, resp, err := c.Build.Restart("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for step %+v", resp.StatusCode, build)
}

func ExampleBuildService_Cancel() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Cancel the build in the server
	_, resp, err := c.Build.Cancel("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for github/octocat/1", resp.StatusCode)
}

func ExampleBuildService_GetBuildToken() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get token for a build from the server
	token, resp, err := c.Build.GetBuildToken("github", "octocat", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for build token %+v", resp.StatusCode, token)
}
