# rest_service

* rest service that hosts the endpoints on localhost:8080
* gRPC client that connects to the gRPC server as defined in the ../rpc_defs shared project on localhost:9000

Passes rest requests onto the grpc_client service for processing by making gRPC calls

## docker image
builddocker.sh shell script will build the docker image

## testing
Postman was used to test the rest service. 
All rest endpoints were tested and responses verified

Next step: Unit tests need to be added to the rest service that mock gRPC calls so that only the rest endpoints are tested
           gRPC content is tested from within the grpc_service itself