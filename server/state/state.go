package state

import (
	"fmt"
	pb "github.com/torczuk/reptile/protocol/client"
	"strings"
	"sync"
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

	//private lock
	lock *sync.Mutex
}

func NewReplicaState() *ReplicaState {
	log := NewLog()
	table := NewClientTable()
	return &ReplicaState{Log: log, ClientTable: table, lock: &sync.Mutex{}}
}

func (t *ReplicaState) String() string {
	var header string
	if t.AmIPrimary() {
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

func (t *ReplicaState) OthersIp() []string {
	others := make([]string, 0)

	for i, address := range t.Configuration {
		if i != t.MyAddress {
			others = append(others, address)
		}
	}
	return others
}

func (t *ReplicaState) AmIPrimary() bool {
	ip := t.MyIp()
	return len(t.Configuration) > 0 && strings.Compare(ip, t.Configuration[0]) == 0
}

func (t *ReplicaState) RegisterRequest(request *pb.ClientRequest, operationRes string) *pb.ClientResponse {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.OpNum = t.Log.Add(request.ClientId, request.Operation)
	response := &pb.ClientResponse{RequestNum: request.RequestNum, Response: operationRes, OperationNum: t.OpNum}
	t.ClientTable.SaveRequest(request, response)
	return response
}

func (t *ReplicaState) Commit(operationNum int) (uint32, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	_, err := t.Log.Commit(operationNum)
	if err == nil {
		t.CommitNum = max(uint32(operationNum), t.CommitNum)
	}
	return t.CommitNum, err
}

func (t *ReplicaState) IsCommitted(operationNum int) bool {
	return t.Log.IsCommitted(operationNum)
}

func max(x, y uint32) uint32 {
	if x < y {
		return y
	}
	return x
}
