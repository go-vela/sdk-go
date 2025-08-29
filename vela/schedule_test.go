// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
)

func TestSchedule_Get(t *testing.T) {
	s := httptest.NewServer(server.FakeHandler())

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Errorf("unable to create test client: %v", err)
	}

	var schedule api.Schedule

	err = json.Unmarshal([]byte(server.ScheduleResp), &schedule)
	if err != nil {
		t.Errorf("unable to create test schedule: %v", err)
	}

	type args struct {
		org      string
		repo     string
		schedule string
	}

	tests := []struct {
		failure  bool
		name     string
		args     args
		want     *api.Schedule
		wantResp int
	}{
		{
			failure: false,
			name:    "success with 200",
			args: args{
				org:      "github",
				repo:     "octocat",
				schedule: "foo",
			},
			want:     &schedule,
			wantResp: http.StatusOK,
		},
		{
			failure: true,
			name:    "failure with 404",
			args: args{
				org:      "github",
				repo:     "octocat",
				schedule: "not-found",
			},
			want:     new(api.Schedule),
			wantResp: http.StatusNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotResp, err := c.Schedule.Get(test.args.org, test.args.repo, test.args.schedule)

			if test.failure {
				if err == nil {
					t.Errorf("Get for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Get for %s returned err: %v", test.name, err)
			}

			if !reflect.DeepEqual(gotResp.StatusCode, test.wantResp) {
				t.Errorf("Get for %s is %v, want %v", test.name, gotResp.StatusCode, test.wantResp)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Get for %s is %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestSchedule_GetAll(t *testing.T) {
	t.Skip() // server.SchedulesResp is a poorly formatted string. TODO: fix in v0.24

	s := httptest.NewServer(server.FakeHandler())

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Errorf("unable to create test client: %v", err)
	}

	var schedules []api.Schedule

	err = json.Unmarshal([]byte(server.SchedulesResp), &schedules)
	if err != nil {
		t.Errorf("unable to create test schedules: %v", err)
	}

	type args struct {
		org  string
		repo string
		opts *ListOptions
	}

	tests := []struct {
		failure  bool
		name     string
		args     args
		want     []api.Schedule
		wantResp int
	}{
		{
			failure: false,
			name:    "success with 200",
			args: args{
				org:  "github",
				repo: "octocat",
				opts: nil,
			},
			want:     schedules,
			wantResp: http.StatusOK,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotResp, err := c.Schedule.GetAll(test.args.org, test.args.repo, test.args.opts)

			if test.failure {
				if err == nil {
					t.Errorf("GetAll for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("GetAll for %s returned err: %v", test.name, err)
			}

			if !reflect.DeepEqual(gotResp.StatusCode, test.wantResp) {
				t.Errorf("GetAll for %s is %v, want %v", test.name, gotResp.StatusCode, test.wantResp)
			}

			if !reflect.DeepEqual(*got, test.want) {
				t.Errorf("GetAll for %s is %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestSchedule_Add(t *testing.T) {
	s := httptest.NewServer(server.FakeHandler())

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Errorf("unable to create test client: %v", err)
	}

	var schedule api.Schedule

	err = json.Unmarshal([]byte(server.ScheduleResp), &schedule)
	if err != nil {
		t.Errorf("unable to create test schedule: %v", err)
	}

	type args struct {
		org      string
		repo     string
		schedule *api.Schedule
	}

	tests := []struct {
		failure  bool
		name     string
		args     args
		want     *api.Schedule
		wantResp int
	}{
		{
			failure: false,
			name:    "success with 201",
			args: args{
				org:  "github",
				repo: "octocat",
				schedule: &api.Schedule{
					Active: Bool(true),
					Name:   String("foo"),
					Entry:  String("@weekly"),
				},
			},
			want:     &schedule,
			wantResp: http.StatusCreated,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotResp, err := c.Schedule.Add(test.args.org, test.args.repo, test.args.schedule)

			if test.failure {
				if err == nil {
					t.Errorf("Add for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Add for %s returned err: %v", test.name, err)
			}

			if !reflect.DeepEqual(gotResp.StatusCode, test.wantResp) {
				t.Errorf("Add for %s is %v, want %v", test.name, gotResp.StatusCode, test.wantResp)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Add for %s is %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestSchedule_Update(t *testing.T) {
	s := httptest.NewServer(server.FakeHandler())

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Errorf("unable to create test client: %v", err)
	}

	var schedule api.Schedule

	err = json.Unmarshal([]byte(server.ScheduleResp), &schedule)
	if err != nil {
		t.Errorf("unable to create test schedule: %v", err)
	}

	type args struct {
		org      string
		repo     string
		schedule *api.Schedule
	}

	tests := []struct {
		failure  bool
		name     string
		args     args
		want     *api.Schedule
		wantResp int
	}{
		{
			failure: false,
			name:    "success with 200",
			args: args{
				org:  "github",
				repo: "octocat",
				schedule: &api.Schedule{
					Active: Bool(true),
					Name:   String("foo"),
					Entry:  String("@weekly"),
				},
			},
			want:     &schedule,
			wantResp: http.StatusOK,
		},
		{
			failure: true,
			name:    "failure with 404",
			args: args{
				org:  "github",
				repo: "octocat",
				schedule: &api.Schedule{
					Name: String("not-found"),
				},
			},
			want:     new(api.Schedule),
			wantResp: http.StatusNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotResp, err := c.Schedule.Update(test.args.org, test.args.repo, test.args.schedule)

			if test.failure {
				if err == nil {
					t.Errorf("Update for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Update for %s returned err: %v", test.name, err)
			}

			if !reflect.DeepEqual(gotResp.StatusCode, test.wantResp) {
				t.Errorf("Update for %s is %v, want %v", test.name, gotResp.StatusCode, test.wantResp)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Update for %s is %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func TestSchedule_Remove(t *testing.T) {
	s := httptest.NewServer(server.FakeHandler())

	c, err := NewClient(s.URL, "", nil)
	if err != nil {
		t.Errorf("unable to create test client: %v", err)
	}

	type args struct {
		org      string
		repo     string
		schedule string
	}

	tests := []struct {
		failure  bool
		name     string
		args     args
		want     *string
		wantResp int
	}{
		{
			failure: false,
			name:    "success with 200",
			args: args{
				org:      "github",
				repo:     "octocat",
				schedule: "foo",
			},
			want:     String("schedule foo deleted"),
			wantResp: http.StatusOK,
		},
		{
			failure: true,
			name:    "failure with 404",
			args: args{
				org:      "github",
				repo:     "octocat",
				schedule: "not-found",
			},
			want:     String("Schedule not-found does not exist"),
			wantResp: http.StatusNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotResp, err := c.Schedule.Remove(test.args.org, test.args.repo, test.args.schedule)

			if test.failure {
				if err == nil {
					t.Errorf("Remove for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Remove for %s returned err: %v", test.name, err)
			}

			if !reflect.DeepEqual(gotResp.StatusCode, test.wantResp) {
				t.Errorf("Remove for %s is %v, want %v", test.name, gotResp.StatusCode, test.wantResp)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Remove for %s is %v, want %v", test.name, got, test.want)
			}
		})
	}
}

func ExampleScheduleService_Get() {
	// create a new vela client for interacting with server
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		fmt.Println(err)
	}

	// set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// get a schedule from a repo in the server
	schedule, resp, err := c.Schedule.Get("github", "octocat", "nightly")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("received response code %d, for schedule %+v", resp.StatusCode, schedule)
}

func ExampleScheduleService_GetAll() {
	// create a new vela client for interacting with server
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		fmt.Println(err)
	}

	// set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// get all the schedules from a repo in the server
	schedules, resp, err := c.Schedule.GetAll("github", "octocat", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("received response code %d, for schedules %+v", resp.StatusCode, schedules)
}

func ExampleScheduleService_Add() {
	// create a new vela client for interacting with server
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		fmt.Println(err)
	}

	// set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Schedule{
		Active: Bool(true),
		Name:   String("nightly"),
		Entry:  String("0 0 * * *"),
	}

	// create the schedule in the server
	schedule, resp, err := c.Schedule.Add("github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("received response code %d, for schedule %+v", resp.StatusCode, schedule)
}

func ExampleScheduleService_Update() {
	// create a new vela client for interacting with server
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		fmt.Println(err)
	}

	// set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := api.Schedule{
		Active: Bool(false),
		Name:   String("nightly"),
		Entry:  String("0 0 * * *"),
	}

	// update the schedule in the server
	schedule, resp, err := c.Schedule.Update("github", "octocat", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("received response code %d, for schedule %+v", resp.StatusCode, schedule)
}

func ExampleScheduleService_Remove() {
	// create a new vela client for interacting with server
	c, err := NewClient("http://localhost:8080", "", nil)
	if err != nil {
		fmt.Println(err)
	}

	// set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// remove the schedule from the server
	schedule, resp, err := c.Schedule.Remove("github", "octocat", "nightly")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("received response code %d, for step %+v", resp.StatusCode, schedule)
}
