package main

import (
	"log"

	"github.com/jeffOSRocko/rest_grpc/rpc_defs/src/keyvalue"
	"github.com/jeffOSRocko/rest_service/grpc_client"
	"github.com/jeffOSRocko/rest_service/rest_server"

	"github.com/gin-gonic/gin"
)

// Inject our gRPC client connection into the context so all
// rest handlers can have access
func GRPCClientMiddleware(grpc_client *keyvalue.KeyValueServiceClient) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("grpcClient", *grpc_client)
		context.Next()
	}
}

func setupRestRouter(grpc_client *keyvalue.KeyValueServiceClient) *gin.Engine {
	router := gin.Default()

	// Inject the gRPC client into the context
	router.Use(GRPCClientMiddleware(grpc_client))

	// Wire in all the rest endpoints
	router.GET("/getidval/:key", rest_server.GetIDVal)
	router.PATCH("/modifyidval", rest_server.ModifyIDVal)
	router.GET("/getallidvals", rest_server.GetAllIDVals)
	router.DELETE("/deleteidval/:key", rest_server.DeleteIDVal)
	router.PUT("/addidval", rest_server.AddIdVal)

	return router
}

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
	router := setupRestRouter(client)

	// Start the rest server!
	router.Run("localhost:8080")
}
