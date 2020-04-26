package primary

import (
	"github.com/stretchr/testify/assert"
	client "github.com/torczuk/reptile/protocol/client"
	"github.com/torczuk/reptile/server/state"
	"testing"
)

func TestExecute_LastClientRequestIsMemorized(t *testing.T) {
	request := &client.ClientRequest{Operation: "exec", ClientId: "client-1", RequestNum: uint32(10)}
	table := &state.ClientTable{Mapping: make(map[string]*client.ClientResponse)}
	log := &state.Log{Sequence: make([]*state.Operation, 0)}
	state := &state.ReplicaState{ClientTable: table, Log: log}

	res, err := Execute(request, state)
	assert.Nil(t, err)
	assert.Equal(t, &client.ClientResponse{RequestNum: request.RequestNum, Response: "Response: exec"}, res)
	assert.Equal(t, res, table.Mapping["client-1"])
}

func TestExecute_LastClientRequestIsAppendedToLogAsUnCommittedOperation(t *testing.T) {
	request := &client.ClientRequest{Operation: "exec", ClientId: "client-1", RequestNum: uint32(10)}
	table := &state.ClientTable{Mapping: make(map[string]*client.ClientResponse)}
	log := &state.Log{Sequence: make([]*state.Operation, 0)}
	replState := &state.ReplicaState{ClientTable: table, Log: log}

	Execute(request, replState)
	assert.Contains(t, log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-1"})
}

func TestExecute_LastClientRequestIsAppendedToLogOnlyOnce(t *testing.T) {
	request := &client.ClientRequest{Operation: "exec", ClientId: "client-1", RequestNum: uint32(10)}
	table := &state.ClientTable{Mapping: make(map[string]*client.ClientResponse)}
	log := &state.Log{Sequence: make([]*state.Operation, 0)}
	replState := &state.ReplicaState{ClientTable: table, Log: log}

	Execute(request, replState)
	Execute(request, replState)
	assert.Contains(t, log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-1"})
	assert.Len(t, log.Sequence, 1)
}
