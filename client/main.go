package client

import (
	pb "github.com/torczuk/reptile/protocol/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type ReptileClient struct {
	Id         string
	Address    string
	RequestNum uint32
}

func (s *ReptileClient) Request(string) (*pb.ClientResponse, error) {
	conn, err := grpc.Dial(s.Address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("can't connect: %v", err)
		return nil, err
	}
	client := pb.NewReptileClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Request(ctx, &pb.ClientRequest{RequestNum: s.RequestNum, ClientId: s.Id, Operation: "NoOp"})
	s.RequestNum += 1
	return res, err
}
