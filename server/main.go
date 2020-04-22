package main

import (
	"context"
	client "github.com/torczuk/reptile/protocol/client"
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

type server struct {
	client.UnimplementedReptileServer
}

func (s *server) Request(ctx context.Context, in *client.ClientRequest) (*client.ClientResponse, error) {
	log.Printf("Received: %v", in)
	return primary.Execute(in, replConf.ClientTable)
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
	client.RegisterReptileServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
