// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import (
	"fmt"
	"testing"
)

func TestAuthorizationService_GetLoginURL(t *testing.T) {
	// setup types
	addr := "http://localhost:8080"
	client, _ := NewClient(addr, "vela", nil)
	badClient, _ := NewClient("", "vela", nil)

	type fields struct {
		client *Client
	}
	type args struct {
		opt *LoginOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "basic",
			fields:  fields{client: client},
			args:    args{opt: &LoginOptions{Type: "", Port: ""}},
			want:    fmt.Sprintf("%s/login", addr),
			wantErr: false,
		},
		{
			name:    "cli",
			fields:  fields{client: client},
			args:    args{opt: &LoginOptions{Type: "cli", Port: "123"}},
			want:    fmt.Sprintf("%s/login?port=123&type=cli", addr),
			wantErr: false,
		},
		{
			name:    "web",
			fields:  fields{client: client},
			args:    args{opt: &LoginOptions{Type: "web", Port: ""}},
			want:    fmt.Sprintf("%s/login?type=web", addr),
			wantErr: false,
		},
		{
			name:    "basic bad",
			fields:  fields{client: badClient},
			args:    args{opt: &LoginOptions{Type: "", Port: ""}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &AuthorizationService{
				client: tt.fields.client,
			}
			got, err := svc.GetLoginURL(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizationService.GetLoginURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthorizationService.GetLoginURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
