package client

import (
	"bufio"
	"fmt"
	"github.com/torczuk/reptile/state"
	"net"
	"strconv"
	"strings"
)

var replConf = &state.ReplicaState{
	OpNum:       0,
	Log:         make([]int, 0),
	CommitNum:   0,
	ClientTable: &state.ClientTable{Mapping: make(map[string]*state.ClientResponse)},
}

func Handle(conn net.Conn) {
	bytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		conn.Close()
	}

	cliReq, err := CreateRequest(string(bytes))
	cliRes, err := Execute(cliReq, replConf.ClientTable)

	if err != nil {
		conn.Write([]byte("Response: " + err.Error()))
	}
	if cliRes != nil {
		conn.Write(cliRes.Response)
	}
	conn.Close()
}

func CreateRequest(request string) (req *state.ClientRequest, err error) {
	splited := strings.Split(request, " ")

	if len(splited) != 4 || splited[0] != "REQUEST" {
		return nil, fmt.Errorf("wrong req: [%s]", request)
	}

	requestNum, err := strconv.Atoi(splited[3])
	if err != nil {
		return nil, fmt.Errorf("wrong req num: [%s]", request)
	}
	return &state.ClientRequest{Operation: splited[1], ClientId: splited[2], RequestNum: requestNum}, nil
}

func Execute(request *state.ClientRequest, table *state.ClientTable) (req *state.ClientResponse, err error) {
	cliRes, cliErr := LastRequest(request, table)
	if cliErr != nil {
		return nil, cliErr
	}
	if cliRes != nil {
		return cliRes, nil
	}
	// echo response
	echo := fmt.Sprintf("Response: %s", request.Operation)
	return &state.ClientResponse{RequestNum: request.RequestNum, Response: []byte(echo)}, nil
}

func LastRequest(request *state.ClientRequest, table *state.ClientTable) (req *state.ClientResponse, err error) {
	clientId := request.ClientId
	last := LastRequestNum(request, table)
	current := request.RequestNum

	if last == current {
		return table.Mapping[clientId], nil
	} else if last > current {
		return nil, fmt.Errorf("last req num: %v, current was: [%#v]", last, request)
	} else {
		return nil, nil
	}
}

func LastRequestNum(request *state.ClientRequest, table *state.ClientTable) int {
	clientId := request.ClientId
	last := table.Mapping[clientId]
	if last == nil {
		return 0
	}
	return last.RequestNum
}
