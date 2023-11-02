package rest_server

import (
	"log"
	"net/http"

	rpc_defs "github.com/jeffOSRocko/rest_grpc/rpc_defs/src/keyvalue"

	"github.com/gin-gonic/gin"
)

var RPCClient *rpc_defs.KeyValueServiceClient

type idVal struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func GetAllIDVals(context *gin.Context) {
	log.Printf("REST_SERVICE: Received GetAllIDVals")

	foo := *RPCClient
	resp, err := foo.GetAll(context, &rpc_defs.GetAllRequest{})

	if err != nil {
		context.IndentedJSON(http.StatusOK, "Failed to get all IDVals. Error: "+err.Error())

		return
	}
	context.IndentedJSON(http.StatusOK, resp.Keyvals)
}

func AddIdVal(context *gin.Context) {
	log.Printf("REST_SERVICE: AddIdVal")

	var newIdVal idVal

	if err := context.BindJSON(&newIdVal); err != nil {
		return
	}

	foo := *RPCClient
	resp, err := foo.AddVal(context, &rpc_defs.KeyValue{Key: newIdVal.Key, Val: newIdVal.Val})

	if err != nil {
		context.IndentedJSON(http.StatusOK, "REST_SERVICE: AddIdVal Failed. Error: "+err.Error())

		return
	}
	context.IndentedJSON(http.StatusOK, resp)
}

func ModifyIDVal(context *gin.Context) {
	log.Printf("REST_SERVICE: ModifyIdVal")

	var newIdVal idVal

	if err := context.BindJSON(&newIdVal); err != nil {
		return
	}

	foo := *RPCClient
	resp, err := foo.ModifyVal(context, &rpc_defs.KeyValue{Key: newIdVal.Key, Val: newIdVal.Val})

	if err != nil {
		context.IndentedJSON(http.StatusOK, "REST_SERVICE: ModifyIdVal Failed. Error: "+err.Error())

		return
	}
	context.IndentedJSON(http.StatusOK, resp)
}

func GetIDVal(context *gin.Context) {
	log.Printf("REST_SERVICE: GetIdVal")

	id := context.Param("key")

	foo := *RPCClient
	resp, err := foo.GetVal(context, &rpc_defs.Key{Key: id})

	if err != nil {
		context.IndentedJSON(http.StatusOK, "GetIdVal Failed. Error: "+err.Error())

		return
	}
	context.IndentedJSON(http.StatusOK, resp)
}

func DeleteIDVal(context *gin.Context) {
	log.Printf("REST_SERVICE: DeleteIdVal")

	id := context.Param("key")

	foo := *RPCClient
	resp, err := foo.DeleteVal(context, &rpc_defs.Key{Key: id})

	if err != nil {
		context.IndentedJSON(http.StatusOK, "DeleteIdVal Failed. Error: "+err.Error())

		return
	}
	context.IndentedJSON(http.StatusOK, resp)
}

func Start(gRPClient *rpc_defs.KeyValueServiceClient) {
	RPCClient = gRPClient

	router := gin.Default()

	router.GET("/getallidvals", GetAllIDVals)
	router.POST("/addidval", AddIdVal)
	router.Run("localhost:9090")
}
