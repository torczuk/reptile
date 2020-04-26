package client

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/torczuk/reptile/protocol/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

type ReptileClient struct {
	Id         string
	Address    string
	RequestNum uint32
}

// send request to the replica to execute an operation
func (r *ReptileClient) Request(operation string) (*pb.ClientResponse, error) {
	conn, err := grpc.Dial(r.Address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("can't connect: %v", err)
		return nil, err
	}
	client := pb.NewReptileClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Request(ctx, &pb.ClientRequest{RequestNum: r.RequestNum, ClientId: r.Id, Operation: operation})
	r.RequestNum += 1
	return res, err
}

// for testing
// blocking method that streams all replica logs
func (r *ReptileClient) Log() ([]*pb.ClientLog, error) {
	conn, err := grpc.Dial(r.Address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("can't connect: %v", err)
		return nil, nil
	}
	client := pb.NewReptileClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.Log(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("cant stream logs: %v", err)
	}

	var response []*pb.ClientLog
	for {
		clientLog, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading logs for client: %v %v", client, err)
		}
		response = append(response, clientLog)
	}
	return response, nil
}
