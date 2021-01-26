package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/benc-uk/grpc-experiment/api"
	"google.golang.org/grpc"
)

const timeout = 2000

func main() {
	var host = flag.String("host", "localhost:50051", "Hostname and port of server")
	var name = flag.String("name", "Foobar", "Name to send")
	flag.Parse()

	log.Printf("### Connecting to server: %s\n", *host)

	// Use timeout context for grpc dial
	dialCtx, cancel := context.WithTimeout(context.Background(), timeout*time.Millisecond)
	defer cancel()

	// Set up a connection to the server.
	conn, err := grpc.DialContext(dialCtx, *host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("### Failed to connect to server: %v", err)
	}
	defer conn.Close()
	client := api.NewHelloClient(conn)

	// Context for timeout for messages
	msgCtx, cancel := context.WithTimeout(context.Background(), timeout*time.Millisecond)
	defer cancel()

	// Finally call the API
	r, err := client.SayHello(msgCtx, &api.HelloRequest{
		Name: *name,
	})
	if err != nil {
		log.Fatalf("### Error with API: %v", err)
	}

	log.Printf("### Reply from server: %s", r.GetMessage())
}
