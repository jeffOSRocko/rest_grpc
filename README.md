# rest api that talks server using gRPC

There are two services, a client and a server and a shared rpc_defs project containing the shared proto definition and generated classes

## rpc_defs
Contains the proto file and generated .go files

## grpc_server
grpc server that takes in keyvalue pairs and stores then in an in-memory dictionary

## rest_client
rest client that has rest endpoints to accept POST and GET requests that get passed along to the grpc_server for processing