package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClientRequest_CreateRequest(t *testing.T) {
	expectedRequest := &ClientRequest{
		operation: "exec-operation", clientId: "client-id-1", requestNum: 1,
	}

	request, err := CreateRequest("REQUEST exec-operation client-id-1 1")

	assert.Nil(t, err)
	assert.Equal(t, expectedRequest, request)
}

func TestCreateClientRequest_InvalidRequest(t *testing.T) {
	req, err := CreateRequest("REQUEST exec-operation client-id-1")

	assert.Nil(t, req)
	assert.EqualError(t, err, "wrong req: [REQUEST exec-operation client-id-1]")
}

func TestCreateClientRequest_InvalidRequestType(t *testing.T) {
	req, err := CreateRequest("ABC exec-operation client-id-1 1")

	assert.Nil(t, req)
	assert.EqualError(t, err, "wrong req: [ABC exec-operation client-id-1 1]")
}

func TestCreateClientRequest_InvalidRequestNum(t *testing.T) {
	req, err := CreateRequest("REQUEST exec-operation client-id-1 x")

	assert.Nil(t, req)
	assert.EqualError(t, err, "wrong req num: [REQUEST exec-operation client-id-1 x]")
}

func TestLastRequestNum_FirstRequest(t *testing.T) {
	req := &ClientRequest{operation: "1+1", clientId: "client-id-1", requestNum: 1}
	table := &ClientTable{mapping: make(map[string]*ClientReponse)}

	last := LastRequestNum(req, table)

	assert.Equal(t, last, 0)
}

func TestLastRequest_FirstRequest(t *testing.T) {
	req := &ClientRequest{operation: "1+1", clientId: "client-id-1", requestNum: 1}
	table := &ClientTable{mapping: make(map[string]*ClientReponse)}

	last, err := LastRequest(req, table)

	assert.Nil(t, err)
	assert.Nil(t, last)
}

func TestLastRequest_ReqNumSameThanLast(t *testing.T) {
	req := &ClientRequest{operation: "1+1", clientId: "client-id-1", requestNum: 1}
	res := &ClientReponse{requestNum: 1, response: []byte("2")}

	mapping := make(map[string]*ClientReponse)
	mapping["client-id-1"] = res
	table := &ClientTable{mapping: mapping}

	last, err := LastRequest(req, table)

	assert.Nil(t, err)
	assert.Equal(t, last, res)
}

func TestLastRequest_BadRequest_ReqNumLessThanLast(t *testing.T) {
	req := &ClientRequest{operation: "1+1", clientId: "client-id-1", requestNum: 1}
	res := &ClientReponse{requestNum: 2, response: []byte("2")}

	mapping := make(map[string]*ClientReponse)
	mapping["client-id-1"] = res
	table := &ClientTable{mapping: mapping}

	last, err := LastRequest(req, table)

	assert.Equal(t, err, fmt.Errorf("last req num: %v, current was: [%#v]", 2, req))
	assert.Nil(t, last)
}

func TestLastRequest_NextRequest(t *testing.T) {
	req := &ClientRequest{operation: "1+1", clientId: "client-id-1", requestNum: 2}
	res := &ClientReponse{requestNum: 1, response: []byte("2")}

	mapping := make(map[string]*ClientReponse)
	mapping["client-id-1"] = res
	table := &ClientTable{mapping: mapping}

	last, err := LastRequest(req, table)

	assert.Nil(t, err)
	assert.Nil(t, last)
}