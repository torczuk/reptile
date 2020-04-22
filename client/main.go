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

func (r *ReptileClient) Request(string) (*pb.ClientResponse, error) {
	conn, err := grpc.Dial(r.Address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("can't connect: %v", err)
		return nil, err
	}
	client := pb.NewReptileClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Request(ctx, &pb.ClientRequest{RequestNum: r.RequestNum, ClientId: r.Id, Operation: "NoOp"})
	r.RequestNum += 1
	return res, err
}
