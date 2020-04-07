package client

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type ClientRequest struct {
	operation  string
	clientId   string
	requestNum int
}

func Handle(conn net.Conn) {
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

func CreateRequest(request string) (req *ClientRequest, err error) {
	splited := strings.Split(request, " ")

	if len(splited) != 4 || splited[0] != "REQUEST" {
		return nil, fmt.Errorf("wrong req: [%s]", request)
	}

	requestNum, err := strconv.Atoi(splited[3])
	if err != nil {
		return nil, fmt.Errorf("wrong req num: [%s]", request)
	}
	return &ClientRequest{operation: splited[1], clientId: splited[2], requestNum: requestNum}, nil
}
