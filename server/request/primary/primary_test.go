package primary

import (
	"github.com/stretchr/testify/assert"
	"github.com/torczuk/reptile/server/request/client"
	"github.com/torczuk/reptile/server/state"
	"testing"
)

func TestExecute_LastClientRequestIsMemorized(t *testing.T) {
	request := client.NewClientRequest("exec", "client-1", uint32(10))
	replState := state.NewReplicaState()

	res, err := Execute(request, replState)
	assert.Nil(t, err)
	assert.Equal(t, client.NewClientResponse(request.RequestNum, "Response: exec", uint32(0)), res)
	assert.Equal(t, res, replState.ClientTable.Mapping["client-1"])
}

func TestExecute_LastClientRequestIsAppendedToLogAsUnCommittedOperation(t *testing.T) {
	request := client.NewClientRequest("exec", "client-1", uint32(10))
	replState := state.NewReplicaState()

	Execute(request, replState)
	assert.Contains(t, replState.Log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-1"})
}

func TestExecute_LastClientRequestIsAppendedToLogOnlyOnce(t *testing.T) {
	request := client.NewClientRequest("exec", "client-1", uint32(10))
	replState := state.NewReplicaState()

	Execute(request, replState)
	Execute(request, replState)
	assert.Contains(t, replState.Log.Sequence, &state.Operation{Committed: false, Operation: "exec", ClientId: "client-1"})
	assert.Len(t, replState.Log.Sequence, 1)
}
