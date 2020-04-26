package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	client "github.com/torczuk/reptile/protocol/client"
	server "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/config"
	"github.com/torczuk/reptile/server/network"
	"github.com/torczuk/reptile/server/request/primary"
	"github.com/torczuk/reptile/server/state"
	"google.golang.org/grpc"
	logger "log"
	"net"
)

var replicaLog = &state.Log{Sequence: make([]*state.Operation, 0)}

var replConf = &state.ReplicaState{
	OpNum:       0,
	Log:         replicaLog,
	CommitNum:   0,
	ClientTable: &state.ClientTable{Mapping: make(map[string]*client.ClientResponse)},
}

type reptileServer struct {
	client.UnimplementedReptileServer
	server.UnimplementedReplicaServer
}

func (s *reptileServer) Request(ctx context.Context, in *client.ClientRequest) (*client.ClientResponse, error) {
	logger.Printf("new request: %v", in)
	return primary.Execute(in, replConf)
}

func (s *reptileServer) Log(req *empty.Empty, stream client.Reptile_LogServer) error {
	logger.Printf("streamn log")
	return primary.Log(replConf, stream)
}

func (s *reptileServer) Prepare(ct context.Context, in *server.PrepareReplica) (*server.PrepareOk, error) {
	return &server.PrepareOk{View: 1, OperationNum: 1, ReplicaNum: 1}, nil
}

func main() {
	servers := config.Servers()
	network.SortIPAddresses(servers)
	myAddress, err := network.MyAddress(servers)
	if err != nil {
		logger.Fatal(err)
	}
	replConf.Configuration = servers
	replConf.MyAddress = myAddress
	logger.Print(replConf)
	listener, err := net.Listen("tcp", "0.0.0.0:2600")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	client.RegisterReptileServer(s, &reptileServer{})
	server.RegisterReplicaServer(s, &reptileServer{})

	if err := s.Serve(listener); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
