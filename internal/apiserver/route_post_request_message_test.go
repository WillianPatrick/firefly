// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger-labs/firefly/mocks/privatemessagingmocks"
	"github.com/hyperledger-labs/firefly/pkg/fftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostRequestMessage(t *testing.T) {
	o, r := newTestAPIServer()
	mpm := &privatemessagingmocks.Manager{}
	o.On("PrivateMessaging").Return(mpm)
	input := &fftypes.MessageInOut{
		Group: &fftypes.InputGroup{
			Members: []fftypes.MemberInput{
				{Identity: "org1"},
			},
		},
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/request/message", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mpm.On("RequestReply", mock.Anything, "ns1", mock.AnythingOfType("*fftypes.MessageInOut")).
		Return(&fftypes.MessageInOut{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
}

func TestPostRequestMessageMissingGroup(t *testing.T) {
	o, r := newTestAPIServer()
	mpm := &privatemessagingmocks.Manager{}
	o.On("PrivateMessaging").Return(mpm)
	input := fftypes.MessageInOut{}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("POST", "/api/v1/namespaces/ns1/request/message", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	mpm.On("RequestReply", mock.Anything, "ns1", mock.AnythingOfType("*fftypes.MessageInOut")).
		Return(&fftypes.MessageInOut{}, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 400, res.Result().StatusCode)
}
