package main

import (
	"log"
	"os"

	"github.com/lihongan/simpletest-grpc/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	addr := "127.0.0.1:50080"
	if os.Args != nil && len(os.Args) == 2 {
		addr = os.Args[1]
	} else {
		log.Printf("Use default server address: 127.0.0.1:50080")
	}

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "echo requst"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
