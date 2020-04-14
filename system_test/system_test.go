package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/torczuk/reptile/client"
	"testing"
)

func TestFirstClientRequest(t *testing.T) {
	c := &client.ReptileClient{"test-client", ":2600", 1}

	res, err := c.Request("NoOp")
	assert.Nil(t, err)
	assert.Equal(t, res.Response, "Response: NoOp")
	assert.Equal(t, res.RequestNum, uint32(1))
}

func TestCachedClientRequest(t *testing.T) {
	c1 := &client.ReptileClient{"any-client", ":2600", 1}

	res1, err := c1.Request("NoOp")
	assert.Nil(t, err)
	assert.Equal(t, res1.Response, "Response: NoOp")
	assert.Equal(t, res1.RequestNum, uint32(1))

	//same client id and request num, but different operations
	c2 := &client.ReptileClient{"any-client", ":2600", 1}
	res2, _ := c2.Request("Different")
	assert.Equal(t, res2.Response, "Response: NoOp")
	assert.Equal(t, res2.RequestNum, uint32(1))
}