package helloworldserver

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/protip-dev/examples/internal/proto/helloworld"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

var _ helloworld.HelloWorldServer = (*Server)(nil)

var iso639ToHello = map[string]string{
	"en": "Hello",
	"es": "Hola",
	"de": "Guten Tag",
	"fr": "Bonjour",
}

func (s *Server) SayHello(ctx context.Context, req *helloworld.SayHelloRequest) (*helloworld.SayHelloResponse, error) {
	var iso639 string
	if len(req.Language) == 0 {
		iso639 = "en"
	} else if len(req.Language) == 2 {
		iso639 = req.Language
	} else {
		return nil, status.Error(codes.InvalidArgument, "language should be two letter ISO 639-1 language code")
	}

	hello, ok := iso639ToHello[iso639]
	if !ok {
		return nil, status.Error(codes.NotFound, "language not found")
	}

	resp := helloworld.SayHelloResponse{
		Hello: hello,
	}
	return &resp, nil
}

func (s *Server) CurrentTime(ctx context.Context, req *helloworld.CurrentTimeRequest) (*helloworld.CurrentTimeResponse, error) {
	resp := helloworld.CurrentTimeResponse{
		CurrentTime: ptypes.TimestampNow(),
	}
	return &resp, nil
}
