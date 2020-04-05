package client

import (
	"bufio"
	"fmt"
	"net"
)

func HandleRequest(conn net.Conn) {
	bytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		conn.Close()
	}
	request := string(bytes)
	fmt.Print("Receive: " + request)
	conn.Write([]byte("Response: " + request))
	conn.Close()
}
