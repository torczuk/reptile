package primary

import (
	"fmt"
	client "github.com/torczuk/reptile/protocol/client"
	"github.com/torczuk/reptile/server/state"
	logger "log"
)

func Execute(request *client.ClientRequest, replState *state.ReplicaState) (req *client.ClientResponse, err error) {
	table := replState.ClientTable
	log := replState.Log

	cliRes, cliErr := table.LastRequest(request)
	if cliErr != nil {
		return nil, cliErr
	}
	if cliRes == nil {
		echo := fmt.Sprintf("Response: %s", request.Operation)
		cliRes = &client.ClientResponse{RequestNum: request.RequestNum, Response: echo}
		table.SaveRequest(request, cliRes)
		log.Add(request.ClientId, request.Operation)
	}
	return cliRes, nil
}

func Log(replState *state.ReplicaState, stream client.Reptile_LogServer) (err error) {
	log := replState.Log
	defer stream.Context().Done()
	for _, op := range log.Sequence {
		//if op.Committed {
		err := stream.Send(&client.ClientLog{Log: op.Operation, ClientId: op.ClientId})
		if err != nil {
			logger.Printf("error when log: %v", err)
			break
		}
		//}
	}
	return err
}
