package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/torczuk/reptile/client"
	pb "github.com/torczuk/reptile/protocol/client"
	"testing"
)

func TestFirstClientRequest(t *testing.T) {
	c := &client.ReptileClient{Id: "test-client", Address: ":2600", RequestNum: 1}

	res, err := c.Request("NoOp")
	assert.Nil(t, err)
	assert.Equal(t, "Response: NoOp", res.Response, )
	assert.Equal(t, uint32(1), res.RequestNum)
}

func TestCachedClientRequest(t *testing.T) {
	c1 := &client.ReptileClient{Id: "any-client", Address: ":2600", RequestNum: 1}

	res1, err := c1.Request("NoOp")
	assert.Nil(t, err)
	assert.Equal(t, "Response: NoOp", res1.Response)
	assert.Equal(t, uint32(1), res1.RequestNum)

	//same client id and request num, but different operations
	c2 := &client.ReptileClient{Id: "any-client", Address: ":2600", RequestNum: 1}
	res2, _ := c2.Request("Different")
	assert.Equal(t, "Response: NoOp", res2.Response)
	assert.Equal(t, uint32(1), res2.RequestNum)
}

func TestReplicaPrepare(t *testing.T) {
	c1 := &client.ReptileClient{Id: "test-client-2", Address: ":2600", RequestNum: 1}

	res1, err := c1.Request("1+1")
	assert.Nil(t, err)
	assert.Equal(t, "Response: 1+1", res1.Response)
	assert.Equal(t, uint32(1), res1.RequestNum)

	logs, err := c1.Log()
	assert.Nil(t, err)
	assert.Contains(t, logs, &pb.ClientLog{ClientId: "test-client-2", Log: "1+1"})
}
