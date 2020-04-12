package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/torczuk/reptile/request/primary"
	"log"
	"net"
	"os"
)

const (
	REQUEST = "REQUEST"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:2600")
	if err != nil {
		log.Println("Can't start server", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	log.Println("started listening on port 2600")
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
	log.Printf("handle %v\n", string(request))

	if err != nil {
		fmt.Println("Error reading:", err.Error())
		conn.Close()
	}
	if bytes.HasPrefix(request, []byte(REQUEST)) {
		primary.Handle(request, conn)
	} else {
		fmt.Println("unknown method in: " + string(request))
		conn.Close()
	}
}
