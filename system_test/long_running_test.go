package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/torczuk/reptile/client"
	"testing"
)

const ADDRESS = ":2600"

func Test_ReplicationOnLongRunningExecution(t *testing.T) {
	clientCount := 10
	requestCount := 100
	clientStatus := make(chan bool, clientCount)

	for _, client := range NewClients(clientCount, ADDRESS) {
		go ExecuteRequests(requestCount, client, clientStatus)
	}

	//wait for all responses
	for i := 0; i < clientCount; i++ {
		<-clientStatus
	}

	log1, _ := NewClient("client-log1", ":2600").Log()
	log2, _ := NewClient("client-log2", ":2700").Log()
	log3, _ := NewClient("client-log3", ":2800").Log()

	assert.Equal(t, log1, log2)
	assert.Equal(t, log2, log3)
}

func NewClients(count int, address string) []*client.ReptileClient {
	var clients []*client.ReptileClient
	for i := 0; i < count; i++ {
		clientId := fmt.Sprintf("client-%v", i)
		clients = append(clients, NewClient(clientId, address))
	}
	return clients
}

func ExecuteRequests(count int, c *client.ReptileClient, done chan bool) {
	for i := 0; i < count; i++ {
		request := fmt.Sprintf("NoOp %v", i)
		c.Request(request)
	}
	done <- true
}
