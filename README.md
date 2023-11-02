# rest service that hosts a rest api that in turn talks to another service using gRPC

There are two services, a client and a server and a shared rpc_defs project containing the shared proto definition and generated classes

## docker and docker-compose
The two services - grpc_server and rest_client have their own docker files for generating docker images.
Each image was generated using 'docker build -t <service_name> .' 
   - substitute 'rest_client' and 'grpc_server' for each service

The docker-compose.yaml file can be used to easily bring up and down each container by running
   'docker-compose up' and 'docker-compose down' (in this folder where docker-compose.yaml is located)

## rpc_defs
Contains the proto file and generated .go files

## grpc_server
grpc server that takes in keyvalue pairs and stores then in an in-memory dictionary

## rest_client
rest client that has rest endpoints to accept POST, PATCH, DELETE and GET requests that get passed along to the grpc_server for processing

## testing
Postman was used for manual testing of each service as they were developed
   - Two collections were created: one to probe the rest API and one to probe the gRPC service
     each collection contained http/gRPC calls for each exposed endpoint
     http
        GET: localhost:8080/getallidvals
        GET: localhost:8080/getidval/foo
        PUT: localhost:8080/addidval
        PATCH: localhost:8080/modifyidval
        DELETE: localhost:8080/deleteidval/foo

    gRPC
       localhost:9000/GetAll
       localhost:9000/AddVal
       localhost:9000/GetIDVal
       localhost:9000/DeleteVal
       localhost:9000/ModifyVal 

go test ./... will execute all unit tests within each service