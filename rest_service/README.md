# rest_service

* rest service that hosts the endpoints on localhost:8080
* gRPC client that connects to the gRPC server as defined in the ../rpc_defs shared project on localhost:9000

Passes rest requests onto the grpc_client service for processing by making gRPC calls

## docker image
builddocker.sh shell script will build the docker image