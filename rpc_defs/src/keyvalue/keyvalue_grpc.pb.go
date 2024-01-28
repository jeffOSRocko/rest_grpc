// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: keyvalue.proto

package keyvalue

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KeyValueService_GetVal_FullMethodName    = "/keyvalue.KeyValueService/GetVal"
	KeyValueService_AddVal_FullMethodName    = "/keyvalue.KeyValueService/AddVal"
	KeyValueService_DeleteVal_FullMethodName = "/keyvalue.KeyValueService/DeleteVal"
	KeyValueService_ModifyVal_FullMethodName = "/keyvalue.KeyValueService/ModifyVal"
	KeyValueService_GetAll_FullMethodName    = "/keyvalue.KeyValueService/GetAll"
)

// KeyValueServiceClient is the client API for KeyValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueServiceClient interface {
	GetVal(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValue, error)
	AddVal(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error)
	DeleteVal(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValue, error)
	ModifyVal(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type keyValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueServiceClient(cc grpc.ClientConnInterface) KeyValueServiceClient {
	return &keyValueServiceClient{cc}
}

func (c *keyValueServiceClient) GetVal(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, KeyValueService_GetVal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) AddVal(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, KeyValueService_AddVal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) DeleteVal(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, KeyValueService_DeleteVal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) ModifyVal(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*KeyValue, error) {
	out := new(KeyValue)
	err := c.cc.Invoke(ctx, KeyValueService_ModifyVal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, KeyValueService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueServiceServer is the server API for KeyValueService service.
// All implementations should embed UnimplementedKeyValueServiceServer
// for forward compatibility
type KeyValueServiceServer interface {
	GetVal(context.Context, *Key) (*KeyValue, error)
	AddVal(context.Context, *KeyValue) (*KeyValue, error)
	DeleteVal(context.Context, *Key) (*KeyValue, error)
	ModifyVal(context.Context, *KeyValue) (*KeyValue, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
}

// UnimplementedKeyValueServiceServer should be embedded to have forward compatible implementations.
type UnimplementedKeyValueServiceServer struct {
}

func (UnimplementedKeyValueServiceServer) GetVal(context.Context, *Key) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVal not implemented")
}
func (UnimplementedKeyValueServiceServer) AddVal(context.Context, *KeyValue) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVal not implemented")
}
func (UnimplementedKeyValueServiceServer) DeleteVal(context.Context, *Key) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVal not implemented")
}
func (UnimplementedKeyValueServiceServer) ModifyVal(context.Context, *KeyValue) (*KeyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyVal not implemented")
}
func (UnimplementedKeyValueServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

// UnsafeKeyValueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueServiceServer will
// result in compilation errors.
type UnsafeKeyValueServiceServer interface {
	mustEmbedUnimplementedKeyValueServiceServer()
}

func RegisterKeyValueServiceServer(s grpc.ServiceRegistrar, srv KeyValueServiceServer) {
	s.RegisterService(&KeyValueService_ServiceDesc, srv)
}

func _KeyValueService_GetVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).GetVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueService_GetVal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).GetVal(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_AddVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).AddVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueService_AddVal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).AddVal(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_DeleteVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).DeleteVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueService_DeleteVal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).DeleteVal(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_ModifyVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).ModifyVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueService_ModifyVal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).ModifyVal(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValueService_ServiceDesc is the grpc.ServiceDesc for KeyValueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyvalue.KeyValueService",
	HandlerType: (*KeyValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVal",
			Handler:    _KeyValueService_GetVal_Handler,
		},
		{
			MethodName: "AddVal",
			Handler:    _KeyValueService_AddVal_Handler,
		},
		{
			MethodName: "DeleteVal",
			Handler:    _KeyValueService_DeleteVal_Handler,
		},
		{
			MethodName: "ModifyVal",
			Handler:    _KeyValueService_ModifyVal_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _KeyValueService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keyvalue.proto",
}
