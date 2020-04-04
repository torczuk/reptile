package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	VIEW_NUMBER   = 0
	OP_NUMBER     = 0
	COMMIT_NUMBER = 0
)

func main() {
	l, err := net.Listen("tcp", "localhost:2600")
	if err != nil {
		fmt.Println("Can't start server", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Print("Started listening on port 2600")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	bytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	request := string(bytes)
	fmt.Print("Receive: " + request)
	conn.Write([]byte("Response: " + request))
	conn.Close()
}
