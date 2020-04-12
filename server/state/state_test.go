package state

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	pro "github.com/torczuk/reptile/protocol"
	"testing"
)

func TestLastRequestNum_FirstRequest(t *testing.T) {
	req := &pro.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	table := &ClientTable{Mapping: make(map[string]*pro.ClientResponse)}

	last := table.LastRequestNum(req)

	assert.Equal(t, last, 0)
}

func TestLastRequest_FirstRequest(t *testing.T) {
	req := &pro.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	table := &ClientTable{Mapping: make(map[string]*pro.ClientResponse)}

	last, err := table.LastRequest(req)

	assert.Nil(t, err)
	assert.Nil(t, last)
}

func TestLastRequest_ReqNumSameThanLast(t *testing.T) {
	req := &pro.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	res := &pro.ClientResponse{RequestNum: 1, Response: "2"}

	mapping := make(map[string]*pro.ClientResponse)
	mapping["client-id-1"] = res
	table := &ClientTable{Mapping: mapping}

	last, err := table.LastRequest(req)

	assert.Nil(t, err)
	assert.Equal(t, last, res)
}

func TestLastRequest_BadRequest_ReqNumLessThanLast(t *testing.T) {
	req := &pro.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	res := &pro.ClientResponse{RequestNum: 2, Response: "2"}

	mapping := make(map[string]*pro.ClientResponse)
	mapping["client-id-1"] = res
	table := &ClientTable{Mapping: mapping}

	last, err := table.LastRequest(req)

	assert.Equal(t, err, fmt.Errorf("last req num: %v, current was: [%#v]", 2, req))
	assert.Nil(t, last)
}

func TestLastRequest_NextRequest(t *testing.T) {
	req := &pro.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 2}
	res := &pro.ClientResponse{RequestNum: 1, Response: "2"}

	mapping := make(map[string]*pro.ClientResponse)
	mapping["client-id-1"] = res
	table := &ClientTable{Mapping: mapping}

	last, err := table.LastRequest(req)

	assert.Nil(t, err)
	assert.Nil(t, last)
}
