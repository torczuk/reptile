package backup

import (
	"fmt"
	server "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/request/client"
	"github.com/torczuk/reptile/server/state"
	logger "log"
)

func Prepare(request *server.PrepareReplica, replState *state.ReplicaState) (res *server.PrepareOk, err error) {
	logger.Printf("preparing request %v", request)
	replState.OpNum = replState.OpNum + 1

	cliReq := client.NewClientRequest(request.ClientOperation, request.ClientId, request.ClientReqNum)
	cliRes := client.NewClientResponse(cliReq.RequestNum, fmt.Sprintf("Response: %v", cliReq.Operation), replState.OpNum)

	replState.ClientTable.SaveRequest(cliReq, cliRes)
	replState.Log.Add(request.ClientId, request.ClientOperation)

	return &server.PrepareOk{View: replState.ViewNum, OperationNum: request.OperationNum, ReplicaNum: uint32(replState.MyAddress)}, nil
}

func HeartBean(request *server.HeartBeat, replState *state.ReplicaState) (*server.HeartBeat, error) {
	_, err := replState.Commit(int(request.CommitNum))
	return &server.HeartBeat{CommitNum: request.CommitNum}, err
}
