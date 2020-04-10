package state

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
