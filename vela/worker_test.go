// SPDX-License-Identifier: Apache-2.0

package vela

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-vela/server/mock/server"
	"github.com/go-vela/types/library"

	"github.com/gin-gonic/gin"
)

func TestWorker_Get_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.WorkerResp)

	var want library.Worker
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Worker.Get("worker_1")

	if err != nil {
		t.Errorf("Worker get returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Worker get returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Worker get is %v, want %v", got, want)
	}
}

func TestWorker_Get_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := library.Worker{}

	// run test
	got, resp, err := c.Worker.Get("0")

	if err == nil {
		t.Errorf("Worker get returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Worker get returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Worker get is %v, want %v", got, want)
	}
}

func TestWorker_GetAll_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.WorkersResp)

	var want []library.Worker
	_ = json.Unmarshal(data, &want)

	// run test
	got, resp, err := c.Worker.GetAll()

	if err != nil {
		t.Errorf("Worker get all returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Worker returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Worker getall is %v, want %v", got, want)
	}
}

func TestWorker_Add_201(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.AddWorkerResp)

	var want library.Token
	_ = json.Unmarshal(data, &want)

	req := library.Worker{
		ID:       Int64(1),
		Hostname: String("worker_1"),
		Address:  String("http://vela:8080"),
		Routes: Strings([]string{
			"large",
			"docker",
			"large:docker",
		}),
		Active:        Bool(true),
		LastCheckedIn: Int64(1602612590),
	}

	// run test
	got, resp, err := c.Worker.Add(&req)

	if err != nil {
		t.Errorf("Worker add returned err: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Worker add returned %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Worker add is %v, want %v", got, want)
	}
}

func TestWorker_RefreshAuth_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.AddWorkerResp)

	var want library.Token
	_ = json.Unmarshal(data, &want)

	worker := "worker_1"

	// run test
	got, resp, err := c.Worker.RefreshAuth(worker)

	if err != nil {
		t.Errorf("Worker RefreshAuth returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Worker RefreshAuth returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Worker RefreshAuth is %v, want %v", got, want)
	}
}

func TestWorker_RefreshAuth_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	worker := "0"

	// run test
	_, resp, err := c.Worker.RefreshAuth(worker)

	if err == nil {
		t.Error("Worker RefreshAuth should have returned err")
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Worker RefreshAuth returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}
}

func TestWorker_Update_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	data := []byte(server.WorkerResp)

	var want library.Worker
	_ = json.Unmarshal(data, &want)

	req := library.Worker{
		Active: Bool(true),
	}

	// run test
	got, resp, err := c.Worker.Update("worker_1", &req)

	if err != nil {
		t.Errorf("Worker update returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Worker returned %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Worker create is %v, want %v", got, want)
	}
}

func TestWorker_Update_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	want := library.Worker{}

	req := library.Worker{
		Active: Bool(true),
	}

	// run test
	got, resp, err := c.Worker.Update("0", &req)

	if err == nil {
		t.Errorf("Worker update returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Worker returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("Worker update is %v, want %v", got, want)
	}
}

func TestWorker_Remove_200(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Worker.Remove("worker_1")

	if err != nil {
		t.Errorf("Worker remove returned err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Worker remove returned %v, want %v", resp.StatusCode, http.StatusOK)
	}
}

func TestWorker_Remove_404(t *testing.T) {
	// setup context
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())
	c, _ := NewClient(s.URL, "", nil)

	// run test
	_, resp, err := c.Worker.Remove("0")

	if err == nil {
		t.Errorf("Worker remove returned err: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Worker remove returned %v, want %v", resp.StatusCode, http.StatusNotFound)
	}
}

func ExampleWorkerService_Get() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get a worker from the server
	worker, resp, err := c.Worker.Get("worker_1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for worker %+v", resp.StatusCode, worker)
}

func ExampleWorkerService_GetAll() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Get all the workers from the server
	workers, resp, err := c.Worker.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for workers %+v", resp.StatusCode, workers)
}

func ExampleWorkerService_Add() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := library.Worker{
		ID:       Int64(1),
		Hostname: String("worker_1"),
		Address:  String("http://vela:8080"),
		Routes: Strings([]string{
			"large",
			"docker",
			"large:docker",
		}),
		Active:        Bool(true),
		LastCheckedIn: Int64(1602612590),
	}

	// Create the worker in the server
	worker, resp, err := c.Worker.Add(&req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for worker %+v", resp.StatusCode, worker)
}

func ExampleWorkerService_RefreshAuth() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	worker := "worker_1"

	// Refresh a worker token with the server
	_, resp, err := c.Worker.RefreshAuth(worker)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for worker %+v", resp.StatusCode, worker)
}

func ExampleWorkerService_Update() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	req := library.Worker{
		Active: Bool(false),
	}

	// Update the worker in the server
	worker, resp, err := c.Worker.Update("worker_1", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for worker %+v", resp.StatusCode, worker)
}

func ExampleWorkerService_Remove() {
	// Create a new vela client for interacting with server
	c, _ := NewClient("http://localhost:8080", "", nil)

	// Set new token in existing client
	c.Authentication.SetPersonalAccessTokenAuth("token")

	// Remove the worker in the server
	worker, resp, err := c.Worker.Remove("worker_1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Received response code %d, for worker %+v", resp.StatusCode, worker)
}
