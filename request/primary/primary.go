package primary

import (
	"fmt"
	"github.com/torczuk/reptile/request/replica"
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

func Handle(request []byte, conn net.Conn) {
	cliReq, err := CreateRequest(string(request))
	cliRes, err := Execute(cliReq, replConf.ClientTable)

	if err != nil {
		conn.Write([]byte("Response: " + err.Error()))
	}
	if cliRes != nil {
		conn.Write([]byte(string(cliRes.Response) + "\n"))
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
	cliRes, cliErr := table.LastRequest(request)
	if cliErr != nil {
		return nil, cliErr
	}
	if cliRes == nil {
		echo := fmt.Sprintf("Response: %s", request.Operation)
		cliRes = &state.ClientResponse{RequestNum: request.RequestNum, Response: []byte(echo)}
		table.SaveRequest(request, cliRes)
	}
	replica.Prepare()
	return cliRes, nil
}