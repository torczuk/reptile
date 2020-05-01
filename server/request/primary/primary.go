package primary

import (
	"fmt"
	client "github.com/torczuk/reptile/protocol/client"
	pb "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/reptile"
	"github.com/torczuk/reptile/server/state"
	logger "log"
)

func Execute(request *client.ClientRequest, replState *state.ReplicaState) (res *client.ClientResponse, err error) {
	table := replState.ClientTable
	log := replState.Log

	res, cliErr := table.LastRequest(request)
	if cliErr != nil {
		return nil, cliErr
	}
	if res == nil {
		echo := fmt.Sprintf("Response: %s", request.Operation)
		res = &client.ClientResponse{RequestNum: request.RequestNum, Response: echo}
		table.SaveRequest(request, res)
		log.Add(request.ClientId, request.Operation)
	}
	return res, nil
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

func NotifyReplica(replica string, prepare *pb.PrepareReplica) (*pb.PrepareOk, error) {
	reptile := reptile.NewReptileClient(replica)
	return reptile.Prepare(prepare)
}

func ExecuteRequest(request *client.ClientRequest, replState *state.ReplicaState) (*client.ClientResponse, error) {
	res, err := Execute(request, replState)
	if err != nil {
		logger.Printf("error when executing request: %v", err)
		return nil, err
	}

	clientReqNum := uint32(len(replState.Log.Sequence))
	prepare := &pb.PrepareReplica{View: replState.ViewNum, ClientOperation: request.Operation, ClientId: request.ClientId, ClientReqNum: clientReqNum}

	ips := replState.OthersIp()
	for _, ip := range ips {
		logger.Printf("preparing replica %v", ip)
		NotifyReplica(ip, prepare)
	}
	return res, nil
}
