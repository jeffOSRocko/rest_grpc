package grpc_client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rpc_defs "rpc_defs/src/keyvalue"
)

func SetupGRPCClient() (*grpc.ClientConn, *rpc_defs.KeyValueServiceClient, error) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}

	client := rpc_defs.NewKeyValueServiceClient(conn)

	return conn, &client, nil
}
