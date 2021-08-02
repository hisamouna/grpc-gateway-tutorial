package main

import (
	"context"
	"log"
	"net"

	helloworldpb "github.com/hisamouna/grgw-tutorial/internal/helloworld"
	"google.golang.org/grpc"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloResponse, error) {
	return &helloworldpb.HelloResponse{
		Message: in.Name + " world",
	}, nil
}

func (s *server) HiHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloResponse, error) {
	return &helloworldpb.HelloResponse{
		Message: "Hi! " + in.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	s := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(s, &server{})
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
