package backup

import (
	"fmt"
	pbs "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/request/client"
	"github.com/torczuk/reptile/server/state"
	logger "log"
)

func Prepare(request *pbs.PrepareReplica, replState *state.ReplicaState) (res *pbs.PrepareOk, err error) {
	logger.Printf("preparing request %v", request)
	replState.OpNum = replState.OpNum + 1

	cliReq := client.NewClientRequest(request.ClientOperation, request.ClientId, request.ClientReqNum)
	cliRes := client.NewClientResponse(cliReq.RequestNum, fmt.Sprintf("Response: %v", cliReq.Operation), replState.OpNum)

	replState.ClientTable.SaveRequest(cliReq, cliRes)
	replState.Log.Add(request.ClientId, request.ClientOperation)

	return &pbs.PrepareOk{View: replState.ViewNum, OperationNum: request.OperationNum, ReplicaNum: uint32(replState.MyAddress)}, nil
}
