package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/benc-uk/grpc-experiment/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var host = flag.String("host", "localhost:50051", "Hostname and port of server")

	var name = flag.String("name", "Foobar", "Name to send")

	var msgCount = flag.Int("count", 1, "Number of calls to make")

	var useLoadBalancer = flag.Bool("loadbalance", false, "Enable client side loadbalancing")

	var timeout = flag.Int("timeout", 2000, "Network timeout in millisec")

	var secure = flag.Bool("secure", false, "Use a secure TLS connection")

	flag.Parse()

	log.Printf("### Connecting to server: %s\n", *host)

	// Use timeout context for grpc dial
	dialCtx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Millisecond)
	defer cancel()

	opts := []grpc.DialOption{}

	// If secure enabled, set up creds and TLS
	var creds credentials.TransportCredentials

	if *secure {
		log.Printf("### TLS is enabled")

		var tlsConf tls.Config

		tlsConf.MaxVersion = tls.VersionTLS12
		// Important! This allows for TLS *without* client cert validation
		tlsConf.ClientAuth = tls.NoClientCert
		// Important! This allows for self signed certificate
		tlsConf.InsecureSkipVerify = true
		creds = credentials.NewTLS(&tlsConf)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// Enable round robin if client side loadbalancing is enabled
	if *useLoadBalancer {
		log.Printf("### Client side loadbalancing enabled")

		opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	}

	// Set up a connection to the server.
	conn, err := grpc.DialContext(dialCtx, *host, opts...)

	if err != nil {
		log.Fatalf("### Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := api.NewHelloClient(conn)

	// Finally call the API, in a loop to test load balancing
	for i := 0; i < *msgCount; i++ {
		// Context for timeout for messages
		msgCtx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Millisecond)
		defer cancel()

		r, err := client.SayHello(msgCtx, &api.HelloRequest{
			Name: fmt.Sprintf("%s %d", *name, i+1),
		})

		if err != nil {
			log.Fatalf("### Error with API: %v", err)
		}

		log.Printf("### Reply from server: %s", r.GetMessage())
	}
}
