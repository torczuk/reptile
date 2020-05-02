package backup

import (
	"github.com/stretchr/testify/assert"
	pbc "github.com/torczuk/reptile/protocol/client"
	pbs "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/state"
	"testing"
)

func Test_PrepareBackup_RequestOk(t *testing.T) {
	replState := state.NewReplicaState()
	request := &pbs.PrepareReplica{View: 1, ClientOperation: "exec", ClientId: "client-id-1", ClientReqNum: uint32(1), OperationNum: uint32(2), CommitNum: 0}

	res, err := Prepare(request, replState)
	assert.Nil(t, err)
	assert.Equal(t, &pbs.PrepareOk{View: replState.ViewNum, OperationNum: request.OperationNum, ReplicaNum: uint32(replState.MyAddress)}, res)
}

func Test_PrepareBackup_AppendToLog(t *testing.T) {
	replState := state.NewReplicaState()
	request := &pbs.PrepareReplica{View: 1, ClientOperation: "exec", ClientId: "client-id-1", ClientReqNum: uint32(1), OperationNum: uint32(2), CommitNum: 0}

	Prepare(request, replState)
	assert.Contains(t, replState.Log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-id-1"})
}

func Test_PrepareBackup_MemorizeRequest(t *testing.T) {
	replState := state.NewReplicaState()
	request := &pbs.PrepareReplica{View: 1, ClientOperation: "exec", ClientId: "client-id-1", ClientReqNum: uint32(1), OperationNum: uint32(2), CommitNum: 0}

	Prepare(request, replState)
	assert.Equal(t, &pbc.ClientResponse{Response: "Response: exec", RequestNum: uint32(1), OperationNum: uint32(1)}, replState.ClientTable.Mapping["client-id-1"])
}