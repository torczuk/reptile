package state

import "fmt"

type ReplicaState struct {
	// ip addresses of replicas
	Configuration []string
	// index of ip address
	ReplicaNum int
	// current view number
	ViewNum int
	// current status
	Status int
	// most recently received request
	OpNum int
	// log - opNum queue
	Log []int
	// last committed opNum
	CommitNum int
	//client table, contains registered client and its last response
	ClientTable *ClientTable
}

type ClientTable struct {
	Mapping map[string]*ClientResponse
}

type ClientRequest struct {
	Operation  string
	ClientId   string
	RequestNum int
}

type ClientResponse struct {
	RequestNum int
	Response   []byte
}

func (t *ClientTable) LastRequest(request *ClientRequest) (req *ClientResponse, err error) {
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

func (t *ClientTable) LastRequestNum(request *ClientRequest) int {
	clientId := request.ClientId
	last := t.Mapping[clientId]
	if last == nil {
		return 0
	}
	return last.RequestNum
}

func (t *ClientTable) SaveRequest(request *ClientRequest, res *ClientResponse) {
	t.Mapping[request.ClientId] = res
}
