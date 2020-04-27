package primary

import (
	"github.com/stretchr/testify/assert"
	client "github.com/torczuk/reptile/protocol/client"
	"github.com/torczuk/reptile/server/state"
	"testing"
)

func TestExecute_LastClientRequestIsMemorized(t *testing.T) {
	request := &client.ClientRequest{Operation: "exec", ClientId: "client-1", RequestNum: uint32(10)}
	replState := state.NewReplicaState()

	res, err := Execute(request, replState)
	assert.Nil(t, err)
	assert.Equal(t, &client.ClientResponse{RequestNum: request.RequestNum, Response: "Response: exec"}, res)
	assert.Equal(t, res, replState.ClientTable.Mapping["client-1"])
}

func TestExecute_LastClientRequestIsAppendedToLogAsUnCommittedOperation(t *testing.T) {
	request := &client.ClientRequest{Operation: "exec", ClientId: "client-1", RequestNum: uint32(10)}
	replState := state.NewReplicaState()

	Execute(request, replState)
	assert.Contains(t, replState.Log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-1"})
}

func TestExecute_LastClientRequestIsAppendedToLogOnlyOnce(t *testing.T) {
	request := &client.ClientRequest{Operation: "exec", ClientId: "client-1", RequestNum: uint32(10)}
	replState := state.NewReplicaState()

	Execute(request, replState)
	Execute(request, replState)
	assert.Contains(t, replState.Log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-1"})
	assert.Len(t, replState.Log.Sequence, 1)
}
