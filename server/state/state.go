package state

import (
	"fmt"
	pb "github.com/torczuk/reptile/protocol"
)

type ReplicaState struct {
	// ip addresses of replicas
	Configuration []string
	// index of ip address
	ReplicaNum uint32
	// current view number
	ViewNum uint32
	// current status
	Status uint32
	// most recently received request
	OpNum uint32
	// log - opNum queue
	Log []uint32
	// last committed opNum
	CommitNum uint32
	//client table, contains registered client and its last response
	ClientTable *ClientTable
}

type ClientTable struct {
	Mapping map[string]*pb.ClientResponse
}

func (t *ClientTable) LastRequest(request *pb.ClientRequest) (req *pb.ClientResponse, err error) {
	clientId := request.ClientId
	last := t.LastRequestNum(request)
	current := request.RequestNum

	if last == current {
		return t.Mapping[clientId], nil
	} else if last > current {
		return nil, fmt.Errorf("last req num: %v, current was: [%#v]", last, request)
	} else {
		return nil, nil
	}
}

func (t *ClientTable) LastRequestNum(request *pb.ClientRequest) uint32 {
	clientId := request.ClientId
	last := t.Mapping[clientId]
	if last == nil {
		return 0
	}
	return last.RequestNum
}

func (t *ClientTable) SaveRequest(request *pb.ClientRequest, res *pb.ClientResponse) {
	t.Mapping[request.ClientId] = res
}
