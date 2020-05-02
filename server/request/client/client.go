package client

import pb "github.com/torczuk/reptile/protocol/client"

func NewClientRequest(operation string, clientId string, requestNum uint32) *pb.ClientRequest {
	return &pb.ClientRequest{Operation: operation, ClientId: clientId, RequestNum: requestNum}
}

func NewClientResponse(requestNum uint32, response string, operationNum uint32) *pb.ClientResponse {
	return &pb.ClientResponse{RequestNum: requestNum, Response: response, OperationNum: operationNum}
}
