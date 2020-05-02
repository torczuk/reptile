package state

import (
	"fmt"
	client "github.com/torczuk/reptile/protocol/client"
)

type ClientTable struct {
	Mapping map[string]*client.ClientResponse
}

func NewClientTable() *ClientTable {
	return &ClientTable{Mapping: make(map[string]*client.ClientResponse)}
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
