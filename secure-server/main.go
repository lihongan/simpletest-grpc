package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lihongan/simpletest-grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// gRPC server main
func main() {
	// create a server
	s := api.Server{}

	// create the TLS credential
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("Could not load TLS: %s", err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// create gRPC server object with creds option
	grpcServer := grpc.NewServer(opts...)

	// register Server
	api.RegisterPingServer(grpcServer, &s)

	// create a listener on port 50443
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 50443))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// start the server
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
