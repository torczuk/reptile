package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

func TestReadWrite(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:2600")
	defer conn.Close()
	assert.Nil(t, err)

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	assert.Nil(t, err)
	res := make(chan string)
	go func(rw *bufio.ReadWriter) {
		response, err := rw.ReadString('\n')
		assert.Nil(t, err)
		res <- response
	}(rw)
	time.Sleep(1 * time.Second)

	go func(rw *bufio.ReadWriter) {
		_, err = rw.WriteString("REQUEST NoOp test-client-1 1\n")
		assert.Nil(t, err)
		rw.Flush()
	}(rw)

	response := <-res
	assert.Equal(t, "Response: NoOp\n", response)
}
