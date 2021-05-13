// Copyright © 2021 Kaleido, Inc.
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

package engine

import (
	"fmt"
	"testing"

	"github.com/kaleido-io/firefly/internal/blockchain"
	"github.com/kaleido-io/firefly/internal/fftypes"
	"github.com/kaleido-io/firefly/mocks/aggregatormocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionUpdate(t *testing.T) {
	ma := &aggregatormocks.Aggregator{}
	e := &engine{
		aggregator: ma,
	}
	ma.On("TransactionUpdate", "node1", fftypes.TransactionStatusConfirmed, "protoid", "", mock.Anything).Return(fmt.Errorf("pop"))
	err := e.TransactionUpdate("node1", fftypes.TransactionStatusConfirmed, "protoid", "", nil)
	assert.EqualError(t, err, "pop")
}

func TestSequencedBroadcastBatch(t *testing.T) {
	ma := &aggregatormocks.Aggregator{}
	e := &engine{
		aggregator: ma,
	}
	ma.On("SequencedBroadcastBatch", mock.Anything, "0x12345", "protoid", mock.Anything).Return(fmt.Errorf("pop"))
	err := e.SequencedBroadcastBatch(&blockchain.BroadcastBatch{}, "0x12345", "protoid", nil)
	assert.EqualError(t, err, "pop")
}