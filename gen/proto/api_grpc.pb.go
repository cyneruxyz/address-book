// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIServiceClient interface {
	CreateAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (*AddressField, error)
	ReadAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (APIService_ReadAddressFieldClient, error)
	UpdateAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (*Response, error)
	DeleteAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (*Response, error)
}

type aPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServiceClient(cc grpc.ClientConnInterface) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) CreateAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (*AddressField, error) {
	out := new(AddressField)
	err := c.cc.Invoke(ctx, "/cyneruxyz.api.v1.APIService/CreateAddressField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) ReadAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (APIService_ReadAddressFieldClient, error) {
	stream, err := c.cc.NewStream(ctx, &APIService_ServiceDesc.Streams[0], "/cyneruxyz.api.v1.APIService/ReadAddressField", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIServiceReadAddressFieldClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type APIService_ReadAddressFieldClient interface {
	Recv() (*AddressField, error)
	grpc.ClientStream
}

type aPIServiceReadAddressFieldClient struct {
	grpc.ClientStream
}

func (x *aPIServiceReadAddressFieldClient) Recv() (*AddressField, error) {
	m := new(AddressField)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIServiceClient) UpdateAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cyneruxyz.api.v1.APIService/UpdateAddressField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) DeleteAddressField(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cyneruxyz.api.v1.APIService/DeleteAddressField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
// All implementations must embed UnimplementedAPIServiceServer
// for forward compatibility
type APIServiceServer interface {
	CreateAddressField(context.Context, *AddressFieldQuery) (*AddressField, error)
	ReadAddressField(*AddressFieldQuery, APIService_ReadAddressFieldServer) error
	UpdateAddressField(context.Context, *AddressFieldQuery) (*Response, error)
	DeleteAddressField(context.Context, *AddressFieldQuery) (*Response, error)
	mustEmbedUnimplementedAPIServiceServer()
}

// UnimplementedAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (UnimplementedAPIServiceServer) CreateAddressField(context.Context, *AddressFieldQuery) (*AddressField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddressField not implemented")
}
func (UnimplementedAPIServiceServer) ReadAddressField(*AddressFieldQuery, APIService_ReadAddressFieldServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadAddressField not implemented")
}
func (UnimplementedAPIServiceServer) UpdateAddressField(context.Context, *AddressFieldQuery) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddressField not implemented")
}
func (UnimplementedAPIServiceServer) DeleteAddressField(context.Context, *AddressFieldQuery) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddressField not implemented")
}
func (UnimplementedAPIServiceServer) mustEmbedUnimplementedAPIServiceServer() {}

// UnsafeAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServiceServer will
// result in compilation errors.
type UnsafeAPIServiceServer interface {
	mustEmbedUnimplementedAPIServiceServer()
}

func RegisterAPIServiceServer(s grpc.ServiceRegistrar, srv APIServiceServer) {
	s.RegisterService(&APIService_ServiceDesc, srv)
}

func _APIService_CreateAddressField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressFieldQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).CreateAddressField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyneruxyz.api.v1.APIService/CreateAddressField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).CreateAddressField(ctx, req.(*AddressFieldQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_ReadAddressField_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddressFieldQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServiceServer).ReadAddressField(m, &aPIServiceReadAddressFieldServer{stream})
}

type APIService_ReadAddressFieldServer interface {
	Send(*AddressField) error
	grpc.ServerStream
}

type aPIServiceReadAddressFieldServer struct {
	grpc.ServerStream
}

func (x *aPIServiceReadAddressFieldServer) Send(m *AddressField) error {
	return x.ServerStream.SendMsg(m)
}

func _APIService_UpdateAddressField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressFieldQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).UpdateAddressField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyneruxyz.api.v1.APIService/UpdateAddressField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).UpdateAddressField(ctx, req.(*AddressFieldQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_DeleteAddressField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressFieldQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).DeleteAddressField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyneruxyz.api.v1.APIService/DeleteAddressField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).DeleteAddressField(ctx, req.(*AddressFieldQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// APIService_ServiceDesc is the grpc.ServiceDesc for APIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cyneruxyz.api.v1.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAddressField",
			Handler:    _APIService_CreateAddressField_Handler,
		},
		{
			MethodName: "UpdateAddressField",
			Handler:    _APIService_UpdateAddressField_Handler,
		},
		{
			MethodName: "DeleteAddressField",
			Handler:    _APIService_DeleteAddressField_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadAddressField",
			Handler:       _APIService_ReadAddressField_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api.proto",
}
