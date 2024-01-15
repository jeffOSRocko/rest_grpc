package main

import (
	"log"
	"os"

	"github.com/jeffOSRocko/rest_service/grpc_client"
	"github.com/jeffOSRocko/rest_service/rest_server"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

type Options struct {
	GrpcServiceAddress string `long:"grpc_service_address" description:"Address for the grpc service" env:"GRPC_SERVICE_ADDRESS" default:"host.docker.internal:9000"`
	RestServiceAddress string `long:"rest_service_address" description:"Address for rest service" env:"REST_SERVICE_ADDRESS" default:"0.0.0.0:8080"`
}

func parseArguments() Options {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			logrus.Fatal(err)
		}
	}

	return opts
}

func main() {
	opts := parseArguments()

	// Setup the gRPC client and handle
	// any error, closing the connection, etc
	conn, client, err := grpc_client.SetupGRPCClient(opts.GrpcServiceAddress)

	if err != nil {
		log.Fatalf("failed to connect to gRPC: %v", err)
	}
	defer conn.Close()

	// Setup the REST server and pass on the gRPC client
	// so that gRPC calls can be made from the REST handlers
	router := rest_server.Start(client)

	// Start the rest server!
	router.Run(opts.RestServiceAddress)
}
