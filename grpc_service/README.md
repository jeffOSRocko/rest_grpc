# grpc_service

gRPC server that implements the messages defined in the ../rpc_defs proto file on localhost:9000

## Basic operation
The gRPC service is hosted at port 9000 and has the following endpoints
   localhost:9000/GetAll
   localhost:9000/AddVal
   localhost:9000/GetIDVal
   localhost:9000/DeleteVal
   localhost:9000/ModifyVal   

## docker image
builddocker.sh shell script will build the docker image


