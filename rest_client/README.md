# rest_client

* rest client that hosts the endpoints on localhost:9090
* gRPC client that connects to the gRPC server as defined in the ../rest_client shared project on localhost:9000

Passes rest POST and GET requests onto the grpc_client service for processing by making gRPC calls