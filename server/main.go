package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/benc-uk/go-starter/pkg/envhelper"
	"github.com/benc-uk/grpc-experiment/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement the service
type server struct {
	api.UnimplementedHelloServer
}

func main() {
	// Port to listen on, change the default as you see fit
	serverPort := envhelper.GetEnvInt("PORT", 50051)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register new server
	srv := grpc.NewServer()
	api.RegisterHelloServer(srv, &server{})
	reflection.Register(srv)

	for srvName := range srv.GetServiceInfo() {
		log.Printf("### '%s' server registered", srvName)
	}

	log.Printf("### Starting server on %d\n", serverPort)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("### Failed to serve: %v", err)
	}
}

// SayHello implementation
func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &api.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}
