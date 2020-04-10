package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/torczuk/reptile/request/client"
	"log"
	"net"
	"os"
)

const (
	REQUEST = "REQUEST"
)

func main() {
	l, err := net.Listen("tcp", "localhost:2600")
	if err != nil {
		log.Println("Can't start server", err.Error())
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
	request, _, err := bufio.NewReader(conn).ReadLine()
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		conn.Close()
	}
	if bytes.HasPrefix(request, []byte(REQUEST)) {
		client.Handle(request, conn)
	} else {
		fmt.Println("unknown method in: " + string(request))
		conn.Close()
	}
}
