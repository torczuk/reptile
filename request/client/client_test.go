package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/torczuk/reptile/state"
	"testing"
)

func TestCreateClientRequest_CreateRequest(t *testing.T) {
	expectedRequest := &state.ClientRequest{
		Operation: "exec-operation", ClientId: "client-id-1", RequestNum: 1,
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