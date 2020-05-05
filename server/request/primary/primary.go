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

	res, cliErr := table.LastRequest(request)
	if cliErr != nil {
		return nil, cliErr
	}
	if res == nil {
		operationRes := fmt.Sprintf("Response: %s", request.Operation)
		res = replState.RegisterRequest(request, operationRes)
	}
	return res, nil
}

func Log(replState *state.ReplicaState, stream client.Reptile_LogServer) (err error) {
	log := replState.Log
	defer stream.Context().Done()
	for _, op := range log.Sequence {
		if op.Committed {
			err := stream.Send(&client.ClientLog{Log: op.Operation, ClientId: op.ClientId})
			if err != nil {
				logger.Printf("error when log: %v", err)
				break
			}
		}
	}
	return err
}

func NotifyReplica(replicaIp string, prepare *pb.PrepareReplica, c chan *pb.PrepareOk) {
	reptileCli := reptile.NewReptileClient(replicaIp)
	res, err := reptileCli.Prepare(prepare)
	if err == nil {
		c <- res
	}
}

func ExecuteRequest(request *client.ClientRequest, replState *state.ReplicaState) (res *client.ClientResponse, err error) {
	res, err = Execute(request, replState)
	if err != nil {
		logger.Printf("error when executing request: %v", err)
		return nil, err
	}

	prepare := NewPrepareReplica(res.OperationNum, request, replState)

	ips := replState.OthersIp()
	//wait for all
	c := make(chan *pb.PrepareOk, len(ips))
	//send to all replicas
	for _, ip := range ips {
		logger.Printf("preparing replica %v", ip)
		go NotifyReplica(ip, prepare, c)
	}
	//wait for all responses
	for i := 0; i < len(ips); i++ {
		<-c
	}
	_, err = replState.Commit(int(res.OperationNum))
	return res, nil
}

func NewPrepareReplica(operationNum uint32, request *client.ClientRequest, replState *state.ReplicaState) *pb.PrepareReplica {
	return &pb.PrepareReplica{
		View:            replState.ViewNum,
		ClientOperation: request.Operation,
		ClientId:        request.ClientId,
		ClientReqNum:    request.RequestNum,
		OperationNum:    operationNum,
		CommitNum:       replState.CommitNum}
}
