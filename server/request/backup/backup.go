package backup

import (
	"fmt"
	pbc "github.com/torczuk/reptile/protocol/client"
	pbs "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/state"
)

func Prepare(request *pbs.PrepareReplica, replState *state.ReplicaState) (res *pbs.PrepareOk, err error) {
	replState.OpNum = replState.OpNum + 1

	cliReq := &pbc.ClientRequest{Operation: request.ClientOperation, ClientId: request.ClientId, RequestNum: request.ClientReqNum}
	cliRes := &pbc.ClientResponse{Response: fmt.Sprintf("Response: %v", cliReq.Operation), RequestNum: cliReq.RequestNum}

	replState.ClientTable.SaveRequest(cliReq, cliRes)
	replState.Log.Add(request.ClientId, request.ClientOperation)

	return &pbs.PrepareOk{View: replState.ViewNum, OperationNum: request.OperationNum, ReplicaNum: uint32(replState.MyAddress)}, nil
}
