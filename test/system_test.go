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

	res, err := con.Write([]byte("hello"))
	if err != nil {
		t.Error(err)
	}

	if "Response: hello" != string(res) {
		t.Error("Expected: Response: hello, got : " + string(res))
	}
}
