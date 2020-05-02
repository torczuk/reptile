package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/torczuk/reptile/client"
	pb "github.com/torczuk/reptile/protocol/client"
	"testing"
)

func Test_FirstClientRequest(t *testing.T) {
	c := &client.ReptileClient{Id: "test-client", Address: ":2600", RequestNum: 1}

	res, err := c.Request("NoOp")
	assert.Nil(t, err)
	assert.Equal(t, "Response: NoOp", res.Response, )
	assert.Equal(t, uint32(1), res.RequestNum)
}

func Test_CachedClientRequest(t *testing.T) {
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

func Test_ReplicaPrepare(t *testing.T) {
	c := &client.ReptileClient{Id: "test-client-2", Address: ":2600", RequestNum: 1}

	res1, err := c.Request("1+1")
	assert.Nil(t, err)
	assert.Equal(t, "Response: 1+1", res1.Response)
	assert.Equal(t, uint32(1), res1.RequestNum)

	logs, err := c.Log()
	assert.Nil(t, err)
	assert.Contains(t, logs, &pb.ClientLog{ClientId: "test-client-2", Log: "1+1"})
}

func Test_RequestIsReplicatedAcrossAllReplicas(t *testing.T) {
	c1 := &client.ReptileClient{Id: "test-client-3", Address: ":2600", RequestNum: 1}

	_, err := c1.Request("2+2")
	assert.Nil(t, err)

	c2 := &client.ReptileClient{Id: "test-client-3", Address: ":2700", RequestNum: 1}
	logs2, _ := c2.Log()
	assert.Contains(t, logs2, &pb.ClientLog{ClientId: "test-client-3", Log: "2+2"})

	c3 := &client.ReptileClient{Id: "test-client-3", Address: ":2800", RequestNum: 1}
	logs3, err := c3.Log()
	assert.Contains(t, logs3, &pb.ClientLog{ClientId: "test-client-3", Log: "2+2"})
}