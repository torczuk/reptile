package main

import (
	"net"
	"testing"
)

func TestResponseFromServer(t *testing.T) {
	con, err := net.Dial("localhost", "2600")
	if err != nil {
		t.Error(err)
	}
	defer con.Close()

	res, err := con.Write([]byte("REQUEST NoOp test-client-1 1"))
	if err != nil {
		t.Error(err)
	}

	if "Response: NoOp" != string(res) {
		t.Error("Expected: Response: hello, got : " + string(res))
	}
}
