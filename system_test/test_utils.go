package main

import (
	"github.com/torczuk/reptile/client"
	logger "log"
	"time"
)

type fn func() bool

func await(condition fn, pool time.Duration, maxTime time.Duration) {
	start := time.Now()
	for !condition() {
		time.Sleep(pool)
		elapsed := time.Since(start)
		if elapsed > maxTime {
			logger.Fatalf("condition didn't finished within %v", maxTime)
		}
	}
}

func NewClient(id string, address string) *client.ReptileClient {
	return &client.ReptileClient{Id: id, Address: address, RequestNum: 1}
}
