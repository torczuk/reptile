package main

import (
	"context"
	client "github.com/torczuk/reptile/protocol/client"
	server "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/config"
	"github.com/torczuk/reptile/server/network"
	"github.com/torczuk/reptile/server/request/primary"
	"github.com/torczuk/reptile/server/state"
	"google.golang.org/grpc"
	"log"
	"net"
)

var replConf = &state.ReplicaState{
	OpNum:       0,
	Log:         make([]uint32, 0),
	CommitNum:   0,
	ClientTable: &state.ClientTable{Mapping: make(map[string]*client.ClientResponse)},
}

type reptileServer struct {
	client.UnimplementedReptileServer
	server.UnimplementedReplicaServer
}

func (s *reptileServer) Request(ctx context.Context, in *client.ClientRequest) (*client.ClientResponse, error) {
	log.Printf("Received: %v", in)
return primary.Execute(in, replConf.ClientTable)
}

func (s *reptileServer) Prepare(ct context.Context, in *server.PrepareReplica) (*server.PrepareOk, error) {
	return &server.PrepareOk{View: 1, OperationNum: 1, ReplicaNum: 1}, nil
}

func main() {
	servers := config.Servers()
	network.SortIPAddresses(servers)
	myAddress, err := network.MyAddress(servers)
	if err != nil {
		log.Fatal(err)
	}
	replConf.Configuration = servers
	replConf.MyAddress = myAddress
	log.Print(replConf)
	listener, err := net.Listen("tcp", "0.0.0.0:2600")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	client.RegisterReptileServer(s, &reptileServer{})
	server.RegisterReplicaServer(s, &reptileServer{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
