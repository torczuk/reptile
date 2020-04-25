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
	assert.Equal(t, res.Response, "Response: NoOp")
	assert.Equal(t, res.RequestNum, uint32(1))
}

func TestCachedClientRequest(t *testing.T) {
	c1 := &client.ReptileClient{Id: "any-client", Address: ":2600", RequestNum: 1}

	res1, err := c1.Request("NoOp")
	assert.Nil(t, err)
	assert.Equal(t, res1.Response, "Response: NoOp")
	assert.Equal(t, res1.RequestNum, uint32(1))

	//same client id and request num, but different operations
	c2 := &client.ReptileClient{Id: "any-client", Address: ":2600", RequestNum: 1}
	res2, _ := c2.Request("Different")
	assert.Equal(t, res2.Response, "Response: NoOp")
	assert.Equal(t, res2.RequestNum, uint32(1))
}

func TestReplicaPrepare(t *testing.T) {
	c1 := &client.ReptileClient{Id: "test-client-1", Address: ":2600", RequestNum: 1}

	res1, err := c1.Request("NoOp")
	assert.Nil(t, err)
	assert.Equal(t, res1.Response, "Response: NoOp")
	assert.Equal(t, res1.RequestNum, uint32(1))

	logs, err := c1.Log()
	assert.Nil(t, err)
	assert.Contains(t, logs, &pb.ClientLog{ClientId: "test-client-1", Log: "Response: NoOp"})
}
