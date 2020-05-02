package state

import (
	"fmt"
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

func NewReplicaState() *ReplicaState {
	log := NewLog()
	table := NewClientTable()
	return &ReplicaState{Log: log, ClientTable: table}
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

func (t *ReplicaState) OthersIp() []string {
	others := make([]string, 0)

	for i, address := range t.Configuration {
		if i != t.MyAddress {
			others = append(others, address)
		}
	}
	return others
}
