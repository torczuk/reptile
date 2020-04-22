package reptile

import (
	"context"
	server "github.com/torczuk/reptile/protocol/server"
	"google.golang.org/grpc"
	"log"
	"time"
)

type ReptileClient struct {
	Address string
}

func (r *ReptileClient) Prepare(prepareReplica *server.PrepareReplica) (*server.PrepareOk, error) {
	conn, err := grpc.Dial(r.Address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Printf("can't connect: %v", err)
		return nil, err
	}
	client := server.NewReplicaClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return client.Prepare(ctx, prepareReplica)
}
