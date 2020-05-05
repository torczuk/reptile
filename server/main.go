package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	client "github.com/torczuk/reptile/protocol/client"
	server "github.com/torczuk/reptile/protocol/server"
	"github.com/torczuk/reptile/server/config"
	"github.com/torczuk/reptile/server/executor"
	"github.com/torczuk/reptile/server/network"
	"github.com/torczuk/reptile/server/reptile"
	"github.com/torczuk/reptile/server/request/backup"
	"github.com/torczuk/reptile/server/request/primary"
	"github.com/torczuk/reptile/server/state"
	"google.golang.org/grpc"
	logger "log"
	"net"
	"time"
)

var replConf = state.NewReplicaState()

type reptileServer struct {
	client.UnimplementedReptileServer
	server.UnimplementedReplicaServer
}

func (s *reptileServer) Request(ctx context.Context, in *client.ClientRequest) (*client.ClientResponse, error) {
	logger.Printf("new request: %v", in)
	return primary.ExecuteRequest(in, replConf)
}

func (s *reptileServer) Log(req *empty.Empty, stream client.Reptile_LogServer) error {
	logger.Printf("streamn log")
	return primary.Log(replConf, stream)
}

func (s *reptileServer) Prepare(ct context.Context, in *server.PrepareReplica) (*server.PrepareOk, error) {
	logger.Printf("prepare %v", in)
	return backup.Prepare(in, replConf)
}

func (s *reptileServer) SendHeartBeat(ct context.Context, in *server.HeartBeat) (*server.HeartBeat, error) {
	return backup.HeartBean(in, replConf)
}

func main() {
	configReplicaState(config.Servers(), replConf)
	if replConf.AmIPrimary() {
		scheduleHeartBeat(replConf)
	}
	registerGRPC()
}

func configReplicaState(servers []string, replState *state.ReplicaState) {
	network.SortIPAddresses(servers)
	myAddress, err := network.MyAddress(servers)
	if err != nil {
		logger.Fatal(err)
	}
	replState.Configuration = servers
	replState.MyAddress = myAddress
	logger.Print(replState)
}

func registerGRPC() {
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

func scheduleHeartBeat(replState *state.ReplicaState) {
	logger.Printf("scheduling sending heart beat")
	for _, ip := range replState.OthersIp() {
		reptileCli := reptile.NewReptileClient(ip)
		task := func() {
			reptileCli.SendHeartBeat(&server.HeartBeat{CommitNum: replState.CommitNum})
		}
		go executor.NewExecutor(task, 500*time.Millisecond).Start()
	}
}
