package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lihongan/simpletest-grpc/api"
	"google.golang.org/grpc"
)

// gRPC server main
func main() {
	// create a server
	s := api.Server{}

	// create gRPC object
	grpcServer := grpc.NewServer()

	// register Server
	api.RegisterPingServer(grpcServer, &s)

	// create a listener on port 50080
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 50080))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// start the server
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
