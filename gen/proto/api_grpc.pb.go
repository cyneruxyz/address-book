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

// AddressBookServiceClient is the client API for AddressBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressBookServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *AddressFieldRequest, opts ...grpc.CallOption) (*Response, error)
	Read(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (AddressBookService_ReadClient, error)
	Update(ctx context.Context, in *AddressFieldUpdateRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Response, error)
}

type addressBookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressBookServiceClient(cc grpc.ClientConnInterface) AddressBookServiceClient {
	return &addressBookServiceClient{cc}
}

func (c *addressBookServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cyneruxyz.api.v1.AddressBookService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) Create(ctx context.Context, in *AddressFieldRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cyneruxyz.api.v1.AddressBookService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) Read(ctx context.Context, in *AddressFieldQuery, opts ...grpc.CallOption) (AddressBookService_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &AddressBookService_ServiceDesc.Streams[0], "/cyneruxyz.api.v1.AddressBookService/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &addressBookServiceReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AddressBookService_ReadClient interface {
	Recv() (*AddressField, error)
	grpc.ClientStream
}

type addressBookServiceReadClient struct {
	grpc.ClientStream
}

func (x *addressBookServiceReadClient) Recv() (*AddressField, error) {
	m := new(AddressField)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *addressBookServiceClient) Update(ctx context.Context, in *AddressFieldUpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cyneruxyz.api.v1.AddressBookService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) Delete(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/cyneruxyz.api.v1.AddressBookService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressBookServiceServer is the server API for AddressBookService service.
// All implementations must embed UnimplementedAddressBookServiceServer
// for forward compatibility
type AddressBookServiceServer interface {
	Echo(context.Context, *EchoRequest) (*Response, error)
	Create(context.Context, *AddressFieldRequest) (*Response, error)
	Read(*AddressFieldQuery, AddressBookService_ReadServer) error
	Update(context.Context, *AddressFieldUpdateRequest) (*Response, error)
	Delete(context.Context, *Phone) (*Response, error)
	mustEmbedUnimplementedAddressBookServiceServer()
}

// UnimplementedAddressBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressBookServiceServer struct {
}

func (UnimplementedAddressBookServiceServer) Echo(context.Context, *EchoRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedAddressBookServiceServer) Create(context.Context, *AddressFieldRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAddressBookServiceServer) Read(*AddressFieldQuery, AddressBookService_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedAddressBookServiceServer) Update(context.Context, *AddressFieldUpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAddressBookServiceServer) Delete(context.Context, *Phone) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAddressBookServiceServer) mustEmbedUnimplementedAddressBookServiceServer() {}

// UnsafeAddressBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressBookServiceServer will
// result in compilation errors.
type UnsafeAddressBookServiceServer interface {
	mustEmbedUnimplementedAddressBookServiceServer()
}

func RegisterAddressBookServiceServer(s grpc.ServiceRegistrar, srv AddressBookServiceServer) {
	s.RegisterService(&AddressBookService_ServiceDesc, srv)
}

func _AddressBookService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyneruxyz.api.v1.AddressBookService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyneruxyz.api.v1.AddressBookService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).Create(ctx, req.(*AddressFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddressFieldQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AddressBookServiceServer).Read(m, &addressBookServiceReadServer{stream})
}

type AddressBookService_ReadServer interface {
	Send(*AddressField) error
	grpc.ServerStream
}

type addressBookServiceReadServer struct {
	grpc.ServerStream
}

func (x *addressBookServiceReadServer) Send(m *AddressField) error {
	return x.ServerStream.SendMsg(m)
}

func _AddressBookService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressFieldUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyneruxyz.api.v1.AddressBookService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).Update(ctx, req.(*AddressFieldUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cyneruxyz.api.v1.AddressBookService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).Delete(ctx, req.(*Phone))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressBookService_ServiceDesc is the grpc.ServiceDesc for AddressBookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressBookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cyneruxyz.api.v1.AddressBookService",
	HandlerType: (*AddressBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _AddressBookService_Echo_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AddressBookService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AddressBookService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AddressBookService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _AddressBookService_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api.proto",
}
