package rest_server

import (
	"log"
	"net/http"
	"time"

	"github.com/jeffOSRocko/rest_grpc/rpc_defs/src/keyvalue"
	rpc_defs "github.com/jeffOSRocko/rest_grpc/rpc_defs/src/keyvalue"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type idVal struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func GetIDVal(context *gin.Context) {
	log.Printf("REST_SERVICE: GetIdVal")

	id := context.Param("key")

	grpcClient, exists := context.Value("grpcClient").(rpc_defs.KeyValueServiceClient)
	if !exists {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "gRPC client not found"})
		return
	}
	resp, err := grpcClient.GetVal(context, &rpc_defs.Key{Key: id})

	if err != nil {
		context.String(http.StatusOK, "GetIdVal Failed. Error: "+err.Error())

		return
	}
	// context.IndentedJSON(http.StatusOK, resp)
	context.String(http.StatusOK, resp.Val)
}

func GetAllIDVals(context *gin.Context) {
	log.Printf("REST_SERVICE: Received GetAllIDVals")

	rawVal := context.Value("grpcClient")
	if rawVal == nil {
		log.Printf("REST_SERVICE GetAllIDVals: failed: context.Value('grpcClient')")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "gRPC rawval Client not found"})
		return
	}

	grpcClient, exists := rawVal.(rpc_defs.KeyValueServiceClient)
	if !exists {
		log.Printf("REST_SERVICE GetAllIDVals: gRPC client not found")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "gRPC client not found"})
		return
	}
	resp, err := grpcClient.GetAll(context, &rpc_defs.GetAllRequest{})

	if err != nil {
		log.Printf("REST_SERVICE GetAllIDVals: Error: " + err.Error())
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

	grpcClient, exists := context.Value("grpcClient").(rpc_defs.KeyValueServiceClient)
	if !exists {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "gRPC client not found"})
		return
	}
	resp, err := grpcClient.AddVal(context, &rpc_defs.KeyValue{Key: newIdVal.Key, Val: newIdVal.Val})

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

	grpcClient, exists := context.Value("grpcClient").(rpc_defs.KeyValueServiceClient)
	if !exists {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "gRPC client not found"})
		return
	}
	resp, err := grpcClient.ModifyVal(context, &rpc_defs.KeyValue{Key: newIdVal.Key, Val: newIdVal.Val})

	if err != nil {
		context.IndentedJSON(http.StatusOK, "REST_SERVICE: ModifyIdVal Failed. Error: "+err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, resp)
}

func DeleteIDVal(context *gin.Context) {
	log.Printf("REST_SERVICE: DeleteIdVal")

	id := context.Param("key")

	grpcClient, exists := context.Value("grpcClient").(rpc_defs.KeyValueServiceClient)
	if !exists {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "gRPC client not found"})
		return
	}
	resp, err := grpcClient.DeleteVal(context, &rpc_defs.Key{Key: id})

	if err != nil {
		context.IndentedJSON(http.StatusOK, "DeleteIdVal Failed. Error: "+err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, resp)
}

// Inject our gRPC client connection into the context so all
// rest handlers can have access
func GRPCClientMiddleware(grpc_client *keyvalue.KeyValueServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("grpcClient", *grpc_client)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Start(grpc_client *keyvalue.KeyValueServiceClient) *gin.Engine {
	router := gin.Default()

	// CORS middleware configuration
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // or use AllowOrigins for specific origins
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))

	// Inject the gRPC client into the context
	router.Use(GRPCClientMiddleware(grpc_client))

	// Wire in all the rest endpoints
	router.GET("/getidval/:key", GetIDVal)
	router.PATCH("/modifyidval", ModifyIDVal)
	router.GET("/getallidvals", GetAllIDVals)
	router.DELETE("/deleteidval/:key", DeleteIDVal)
	router.PUT("/addidval", AddIdVal)

	return router
}
