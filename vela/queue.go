// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package vela

import "github.com/go-vela/types/library"

// QueueService handles retrieving queue info from
// the server methods of the Vela API.
type QueueService service

// GetQueueCreds fetches queue credentials based valid registration token.
func (qvc *QueueService) GetInfo() (*library.QueueInfo, *Response, error) {
	// set the API endpoint path we send the request to
	url := "/api/v1/queue/info"

	// library Token type we want to return
	t := new(library.QueueInfo)

	// send request using client
	resp, err := qvc.client.Call("GET", url, nil, t)

	return t, resp, err
}
