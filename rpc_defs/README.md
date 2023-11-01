# rpc_defs

This folder contains the keyvalue.proto defintion file and generated *.pb.go files for consumption by other services

## Generating go files from proto definition
To generate the *.pb.go files from the proto file run the following protoc command from Terminal and at the parent ./cesys folder
The generated files will be located in the ./rpc_defs/src/keyvalue folder
`protoc --proto_path=rpc_defs --go_out=rpc_defs/src --go-grpc_out=require_unimplemented_servers=false:./rpc_defs/src keyvalue.proto`

## Consuming the generated go files
To consume the generated files after running protoc use the following import statement:
`import "github.com/jeffOSRocko/rpc_defs/src/keyvalue"`