package main

import (
	"context"
	"log"
	"net"

	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/adapters/grpcserver"
	"github.com/codelikesuraj/learn-go-with-tests/scaling-acceptance-tests/domain/interactions"
	"google.golang.org/grpc"
)

type GreetServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{Message: interactions.Greet(request.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreetServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
