package primary

import (
	"fmt"
	pb "github.com/torczuk/reptile/protocol"
	"github.com/torczuk/reptile/server/state"
)

func Execute(request *pb.ClientRequest, table *state.ClientTable) (req *pb.ClientResponse, err error) {
	cliRes, cliErr := table.LastRequest(request)
	if cliErr != nil {
		return nil, cliErr
	}
	if cliRes == nil {
		echo := fmt.Sprintf("Response: %s", request.Operation)
		cliRes = &pb.ClientResponse{RequestNum: request.RequestNum, Response: echo}
		table.SaveRequest(request, cliRes)
	}
	return cliRes, nil
}
