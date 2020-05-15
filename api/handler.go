package api

import (
	"golang.org/x/net/context"
	"log"
)

// Server
type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "echo test"}, nil
}
