package backup

import (
	"github.com/stretchr/testify/assert"
	pbc "github.com/torczuk/reptile/protocol/client"
	pbs "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/state"
	"testing"
)

func TestPrepareBackup_RequestOk(t *testing.T) {
	table := &state.ClientTable{Mapping: make(map[string]*pbc.ClientResponse)}
	log := &state.Log{Sequence: make([]*state.Operation, 0)}
	state := &state.ReplicaState{ClientTable: table, Log: log}
	request := &pbs.PrepareReplica{View: 1, ClientOperation: "exec", ClientId: "client-id-1", ClientReqNum: uint32(1), OperationNum: uint32(2), CommitNum: 0}

	res, err := Prepare(request, state)
	assert.Nil(t, err)
	assert.Equal(t, &pbs.PrepareOk{View: state.ViewNum, OperationNum: request.OperationNum, ReplicaNum: uint32(state.MyAddress)}, res)
}

func TestPrepareBackup_AppendToLog(t *testing.T) {
	table := &state.ClientTable{Mapping: make(map[string]*pbc.ClientResponse)}
	log := &state.Log{Sequence: make([]*state.Operation, 0)}
	replState := &state.ReplicaState{ClientTable: table, Log: log}
	request := &pbs.PrepareReplica{View: 1, ClientOperation: "exec", ClientId: "client-id-1", ClientReqNum: uint32(1), OperationNum: uint32(2), CommitNum: 0}

	Prepare(request, replState)
	assert.Contains(t, log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-id-1"})
}

func TestPrepareBackup_MemorizeRequest(t *testing.T) {
	table := &state.ClientTable{Mapping: make(map[string]*pbc.ClientResponse)}
	log := &state.Log{Sequence: make([]*state.Operation, 0)}
	replState := &state.ReplicaState{ClientTable: table, Log: log}
	request := &pbs.PrepareReplica{View: 1, ClientOperation: "exec", ClientId: "client-id-1", ClientReqNum: uint32(1), OperationNum: uint32(2), CommitNum: 0}

	Prepare(request, replState)
	assert.Equal(t, &pbc.ClientResponse{Response: "Response: exec", RequestNum: uint32(1)}, table.Mapping["client-id-1"])
}
