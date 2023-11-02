package main

import (
	"log"

	"github.com/jeffOSRocko/rest_service/grpc_client"
	"github.com/jeffOSRocko/rest_service/rest_server"

	"github.com/gin-gonic/gin"
)

func setupRestRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getidval/:key", rest_server.GetIDVal)
	router.PATCH("/modifyidval", rest_server.ModifyIDVal)
	router.GET("/getallidvals", rest_server.GetAllIDVals)
	router.DELETE("/deleteidval/:key", rest_server.DeleteIDVal)
	router.PUT("/addidval", rest_server.AddIdVal)

	return router
}

func main() {
	// Setup the gRPC client
	conn, client, err := grpc_client.SetupGRPCClient()

	if err != nil {
		log.Fatalf("failed to connect to gRPC: %v", err)
	}
	defer conn.Close()

	// Setup the REST server and pass on the gRPC client
	rest_server.RPCClient = client

	router := setupRestRouter()
	router.Run("localhost:8080")
}
