package grpc_client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rpc_defs "github.com/jeffOSRocko/rest_grpc/rpc_defs/src/keyvalue"
)

// Setup the gRPC client and return the connection object back
// to the caller.
// The caller is responsible for closing the connection as well as
// any error handling required by possible returned errors
func SetupGRPCClient() (*grpc.ClientConn, *rpc_defs.KeyValueServiceClient, error) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}

	client := rpc_defs.NewKeyValueServiceClient(conn)

	return conn, &client, nil
}
