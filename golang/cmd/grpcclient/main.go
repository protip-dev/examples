package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/protip-dev/examples/internal/proto/helloworld"
	"google.golang.org/grpc"
)

func main() {
	target := flag.String("target", "localhost:8888", "grpc target")
	language := flag.String("language", "", "two-letter ISO 639-1 language code")
	flag.Parse()

	if len(*language) == 0 {
		log.Fatal("language not provided")
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, *target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := helloworld.NewHelloWorldClient(conn)
	timeReq := helloworld.CurrentTimeRequest{}
	timeResp, err := client.CurrentTime(ctx, &timeReq)
	if err != nil {
		log.Fatal(err)
	}

	t, err := ptypes.Timestamp(timeResp.CurrentTime)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the server responded with '%s' for the current time\n", t)

	helloReq := helloworld.SayHelloRequest{
		Language: *language,
	}
	helloResp, err := client.SayHello(ctx, &helloReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the server responded with '%s' for language '%s'\n", helloResp.Hello, helloReq.Language)
}
