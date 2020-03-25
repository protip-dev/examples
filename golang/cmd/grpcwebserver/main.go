package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/protip-dev/examples/internal/helloworldserver"
	"github.com/protip-dev/examples/internal/proto/helloworld"
	"google.golang.org/grpc"
)

// NOTE: Not recommended for production use!
//
// Refer to CORS documentation on how to properly configure your server:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func allowAllOrigins(string) bool {
	return true
}

func main() {
	port := flag.Int("port", 8888, "server port")
	flag.Parse()

	grpcServer := grpc.NewServer()
	helloworld.RegisterHelloWorldServer(grpcServer, &helloworldserver.Server{})
	grpcWebServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(allowAllOrigins))

	log.Printf("listening on :%d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = http.Serve(lis, grpcWebServer)
	if err != nil {
		log.Fatalf("failed to start serving: %v", err)
	}
}
