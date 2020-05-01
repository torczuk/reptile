package reptile

import (
	"context"
	"fmt"
	server "github.com/torczuk/reptile/protocol/server"
	"google.golang.org/grpc"
	"log"
	"time"
)

const PORT = "2600"

type ReptileClient struct {
	Address string
}

func NewReptileClient(address string) *ReptileClient {
	return &ReptileClient{address}
}

func (r *ReptileClient) Prepare(prepareReplica *server.PrepareReplica) (*server.PrepareOk, error) {
	addres := fmt.Sprint("%v:%v", r.Address, PORT)
	conn, err := grpc.Dial(addres, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("can't connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := server.NewReplicaClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return client.Prepare(ctx, prepareReplica)
}
