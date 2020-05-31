package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/torczuk/reptile/client"
	"log"
	"testing"
	"time"
)

const ADDRESS = ":2600"

func Test_ReplicationOnLongRunningExecution(t *testing.T) {
	clientCount := 100
	requestCount := 10
	clientStatus := make(chan bool, clientCount)

	before := time.Now().Unix()
	for _, client := range NewClients(clientCount, ADDRESS) {
		go ExecuteRequests(requestCount, client, clientStatus)
	}

	//wait for all responses
	for i := 0; i < clientCount; i++ {
		<-clientStatus
	}
	after := time.Now().Unix()
	total := after - before
	fmt.Printf("total time: %v", total)

	time.Sleep(10 * time.Second)

	log1, _ := NewClient("client-log1", ":2600").Log()
	log2, _ := NewClient("client-log2", ":2700").Log()
	log3, _ := NewClient("client-log3", ":2800").Log()

	assert.Equal(t, log1, log2)
	assert.Equal(t, log2, log3)
}

func Test_ReplicationLong(t *testing.T) {
	log1, _ := NewClient("client-log1", ":2600").Log()
	log2, _ := NewClient("client-log2", ":2700").Log()
	log3, _ := NewClient("client-log3", ":2800").Log()

	fmt.Printf("log1 size %v\n", len(log1))
	fmt.Printf("log2 size %v\n", len(log2))
	fmt.Printf("log3 size %v\n", len(log3))

	count := make(map[string]int)
	for i := 0; i < len(log1); i++ {
		clientId := log1[i].ClientId
		count[clientId] = count[clientId] + 1
	}
	fmt.Printf("%v", count)
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
		_, err := c.Request(request)
		if err != nil {
			log.Panic(err)
		}
		//fmt.Printf("%v: requestNum: %v, operationNum: %v \n", c.Id, res.RequestNum, res.OperationNum)
	}
	done <- true
}
