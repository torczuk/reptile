package primary

import (
	"fmt"
	client "github.com/torczuk/reptile/protocol/client"
	"github.com/torczuk/reptile/server/state"
)

func Execute(request *client.ClientRequest, table *state.ClientTable) (req *client.ClientResponse, err error) {
	cliRes, cliErr := table.LastRequest(request)
	if cliErr != nil {
		return nil, cliErr
	}
	if cliRes == nil {
		echo := fmt.Sprintf("Response: %s", request.Operation)
		cliRes = &client.ClientResponse{RequestNum: request.RequestNum, Response: echo}
		table.SaveRequest(request, cliRes)
	}
	return cliRes, nil
}
