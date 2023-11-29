package main

import (
	"log"

	"github.com/jeffOSRocko/rest_service/grpc_client"
	"github.com/jeffOSRocko/rest_service/rest_server"
)

func main() {
	// Setup the gRPC client and handle
	// any error, closing the connection, etc
	conn, client, err := grpc_client.SetupGRPCClient()

	if err != nil {
		log.Fatalf("failed to connect to gRPC: %v", err)
	}
	defer conn.Close()

	// Setup the REST server and pass on the gRPC client
	// so that gRPC calls can be made from the REST handlers
	router := rest_server.Start(client)

	// Start the rest server!
	router.Run("localhost:8080")
}
