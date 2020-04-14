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