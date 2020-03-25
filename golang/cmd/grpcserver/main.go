package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/protip-dev/examples/internal/helloworldserver"
	"github.com/protip-dev/examples/internal/proto/helloworld"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 8888, "server port")
	flag.Parse()

	grpcServer := grpc.NewServer()
	helloworld.RegisterHelloWorldServer(grpcServer, &helloworldserver.Server{})

	log.Printf("listening on :%d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start serving: %v", err)
	}
}
