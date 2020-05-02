package state

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	client "github.com/torczuk/reptile/protocol/client"
	"testing"
)

func Test_LastRequestNum_FirstRequest(t *testing.T) {
	req := &client.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	table := NewClientTable()

	last := table.LastRequestNum(req)

	assert.Equal(t, last, uint32(0))
}

func Test_LastRequest_FirstRequest(t *testing.T) {
	req := &client.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	table := NewClientTable()

	last, err := table.LastRequest(req)

	assert.Nil(t, err)
	assert.Nil(t, last)
}

func Test_LastRequest_ReqNumSameThanLast(t *testing.T) {
	req := &client.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	res := &client.ClientResponse{RequestNum: 1, Response: "2"}

	mapping := make(map[string]*client.ClientResponse)
	mapping["client-id-1"] = res
	table := &ClientTable{Mapping: mapping}

	last, err := table.LastRequest(req)

	assert.Nil(t, err)
	assert.Equal(t, last, res)
}

func Test_LastRequest_BadRequest_ReqNumLessThanLast(t *testing.T) {
	req := &client.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 1}
	res := &client.ClientResponse{RequestNum: 2, Response: "2"}

	mapping := make(map[string]*client.ClientResponse)
	mapping["client-id-1"] = res
	table := &ClientTable{Mapping: mapping}

	last, err := table.LastRequest(req)

	assert.Equal(t, err, fmt.Errorf("last req num: %v, current was: [%#v]", 2, req))
	assert.Nil(t, last)
}

func Test_LastRequest_NextRequest(t *testing.T) {
	req := &client.ClientRequest{Operation: "1+1", ClientId: "client-id-1", RequestNum: 2}
	res := &client.ClientResponse{RequestNum: 1, Response: "2"}

	mapping := make(map[string]*client.ClientResponse)
	mapping["client-id-1"] = res
	table := &ClientTable{Mapping: mapping}

	last, err := table.LastRequest(req)

	assert.Nil(t, err)
	assert.Nil(t, last)
}
