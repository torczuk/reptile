package state

import (
	"fmt"
	client "github.com/torczuk/reptile/protocol/client"
)

type ReplicaState struct {
	// ip addresses of replicas
	Configuration []string
	// repilica ip address stored in configuration
	MyAddress int
	// index of ip address
	ReplicaNum uint32
	// current view number
	ViewNum uint32
	// current status
	Status uint32
	// most recently received request
	OpNum uint32
	// log - opNum queue
	Log *Log
	// last committed opNum
	CommitNum uint32
	//client table, contains registered client and its last response
	ClientTable *ClientTable
}

type ClientTable struct {
	Mapping map[string]*client.ClientResponse
}

func NewClientTable() *ClientTable {
	return &ClientTable{Mapping: make(map[string]*client.ClientResponse)}
}

func NewReplicaState() *ReplicaState {
	log := NewLog()
	table := NewClientTable()
	return &ReplicaState{Log: log, ClientTable: table}
}

func (t *ClientTable) LastRequest(request *client.ClientRequest) (req *client.ClientResponse, err error) {
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

func (t *ClientTable) LastRequestNum(request *client.ClientRequest) uint32 {
	clientId := request.ClientId
	last := t.Mapping[clientId]
	if last == nil {
		return 0
	}
	return last.RequestNum
}

func (t *ClientTable) SaveRequest(request *client.ClientRequest, res *client.ClientResponse) {
	t.Mapping[request.ClientId] = res
}

func (t *ReplicaState) String() string {
	var header string
	if t.MyAddress == 0 {
		header = fmt.Sprintf("I'am primary, my ip %v", t.MyIp())
	} else {
		header = fmt.Sprintf("I'am backup, my ip %v", t.MyIp())
	}

	return fmt.Sprintf("%v\n\tservers: %v\n\tipNum: %v\n\tview: %v\n\tstatus: %v\n\topNum: %v\n\tcommitNum: %v\n",
		header, t.Configuration, t.MyAddress, t.ViewNum, t.Status, t.OpNum, t.CommitNum)
}

func (t *ReplicaState) MyIp() string {
	if len(t.Configuration) == 0 || len(t.Configuration) < t.MyAddress-1 {
		return ""
	}
	return t.Configuration[t.MyAddress]
}
