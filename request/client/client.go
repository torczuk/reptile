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

type ClientReponse struct {
	requestNum int
	response   []byte
}

type ClientTable struct {
	mapping map[string]*ClientReponse
}

var cliTable = &ClientTable{
	mapping: make(map[string]*ClientReponse),
}

func Handle(conn net.Conn) {
	bytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		conn.Close()
	}

	cliReq, err := CreateRequest(string(bytes))
	cliRes, err := Execute(cliReq, cliTable)

	if err != nil {
		conn.Write([]byte("Response: " + err.Error()))
	}
	if cliRes != nil {
		conn.Write(cliRes.response)
	}
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

func Execute(request *ClientRequest, table *ClientTable) (req *ClientReponse, err error) {
	cliRes, cliErr := LastRequest(request, table)
	if cliErr != nil {
		return nil, cliErr
	}
	if cliRes != nil {
		return cliRes, nil
	}
	// echo response
	echo := fmt.Sprintf("Response: %s", request.operation)
	return &ClientReponse{requestNum: request.requestNum, response: []byte(echo)}, nil
}

func LastRequest(request *ClientRequest, table *ClientTable) (req *ClientReponse, err error) {
	clientId := request.clientId
	last := LastRequestNum(request, table)
	current := request.requestNum

	if last == current {
		return table.mapping[clientId], nil
	} else if last > current {
		return nil, fmt.Errorf("last req num: %v, current was: [%#v]", last, request)
	} else {
		return nil, nil
	}
}

func LastRequestNum(request *ClientRequest, table *ClientTable) int {
	clientId := request.clientId
	last := table.mapping[clientId]
	if last == nil {
		return 0
	}
	return last.requestNum
}
