package main

import (
	"context"
	"log"
	"net"

	"github.com/jeffOSRocko/rest_grpc/rpc_defs/src/keyvalue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KeyValue struct {
	Key   string
	Value string
}

type Server struct {
	stored_vals map[string]string
}

func (s *Server) Init(addInitialValue bool) {
	s.stored_vals = make(map[string]string)

	if addInitialValue {
		// Add an intial value to the store
		s.stored_vals["foo"] = "bar"
	}
}

func (s *Server) AddVal(ctx context.Context, keyVal *keyvalue.KeyValue) (*keyvalue.KeyValue, error) {
	if keyVal == nil || keyVal.Key == "" {
		log.Printf("GRPC_SERVICE: AddVal received invalid keyval from client")
		err := status.Error(codes.InvalidArgument, "missing key.Val")
		return nil, err
	}

	log.Printf("GRPC_SERVICE: Received AddVal message from client: key:%s, val:%s", keyVal.Key, keyVal.Val)

	existingVal, found := s.stored_vals[keyVal.Key]
	if found {
		log.Printf("GRPC_SERVICE: AddVal attempted to add val with existing key entry. key:%s, val:%s", keyVal.Key, existingVal)
		err := status.Error(codes.InvalidArgument, "item with key: '"+keyVal.Key+"' already exists")
		return nil, err
	}
	s.stored_vals[keyVal.Key] = keyVal.Val

	retval := keyvalue.KeyValue{
		Key: keyVal.Key,
		Val: s.stored_vals[keyVal.Key],
	}

	return &retval, nil
}

func (s *Server) GetVal(ctx context.Context, key *keyvalue.Key) (*keyvalue.KeyValue, error) {
	log.Printf("GRPC_SERVICE: Received GetVal message from client: %s", key.Key)

	found_val, found := s.stored_vals[key.Key]
	if !found {
		log.Printf("GRPC_SERVICE: Failed to locate kvp with key=%s", key.Key)

		return nil, status.Error(codes.NotFound, "key: '"+key.Key+"' not found")
	}

	retval := keyvalue.KeyValue{
		Key: key.Key,
		Val: found_val,
	}

	return &retval, nil
}

func (s *Server) GetAll(ctx context.Context, req *keyvalue.GetAllRequest) (*keyvalue.GetAllResponse, error) {
	log.Printf("GRPC_SERVICE: Received GetAll message from client")

	var retval keyvalue.GetAllResponse

	for k, v := range s.stored_vals {
		retval.Keyvals = append(retval.Keyvals, &keyvalue.KeyValue{
			Key: k,
			Val: v,
		})
	}

	return &retval, nil
}

func (s *Server) DeleteVal(ctx context.Context, key *keyvalue.Key) (*keyvalue.KeyValue, error) {
	log.Printf("GRPC_SERVICE: Received DeleteVal message from client: %s", key.Key)

	found_val, found := s.stored_vals[key.Key]
	if !found {
		log.Printf("GRPC_SERVICE: Failed to locate kvp with key=%s", key.Key)

		return nil, status.Error(codes.NotFound, "key: '"+key.Key+"' not found")
	}

	delete(s.stored_vals, key.Key)

	retval := keyvalue.KeyValue{
		Key: key.Key,
		Val: found_val,
	}

	return &retval, nil
}

func (s *Server) ModifyVal(ctx context.Context, keyval *keyvalue.KeyValue) (*keyvalue.KeyValue, error) {
	log.Printf("GRPC_SERVICE: Received ModifyVal message from client: %s", keyval.Key)

	_, found := s.stored_vals[keyval.Key]
	if !found {
		log.Printf("GRPC_SERVICE: Failed to locate kvp with key=%s", keyval.Key)

		return nil, status.Error(codes.NotFound, "key: '"+keyval.Key+"' not found")
	}

	s.stored_vals[keyval.Key] = keyval.Val

	retval := keyvalue.KeyValue{
		Key: keyval.Key,
		Val: s.stored_vals[keyval.Key],
	}

	return &retval, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("GRPC_SERVICE: Failed to listen on port 9000: %v", err)
	}

	s := Server{}
	s.Init(true)

	grpcServer := grpc.NewServer()

	keyvalue.RegisterKeyValueServiceServer(grpcServer, &s)

	log.Printf("Starting gRPC server on port 9000...")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
