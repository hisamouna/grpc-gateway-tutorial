package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	helloworldpb "github.com/hisamouna/grgw-tutorial/internal/helloworld"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.DialContext(
		context.Background(),
		"server:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
