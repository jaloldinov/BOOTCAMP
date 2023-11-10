// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: staff.proto

package staff_service

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

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	Create(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error)
	Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetStaffResponse, error)
	List(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffResponse, error)
	Update(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) Create(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error) {
	out := new(CreateStaffResponse)
	err := c.cc.Invoke(ctx, "/staff_service.StaffService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetStaffResponse, error) {
	out := new(GetStaffResponse)
	err := c.cc.Invoke(ctx, "/staff_service.StaffService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) List(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffResponse, error) {
	out := new(ListStaffResponse)
	err := c.cc.Invoke(ctx, "/staff_service.StaffService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Update(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/staff_service.StaffService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/staff_service.StaffService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations must embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	Create(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error)
	Get(context.Context, *IdRequest) (*GetStaffResponse, error)
	List(context.Context, *ListStaffRequest) (*ListStaffResponse, error)
	Update(context.Context, *UpdateStaffRequest) (*Response, error)
	Delete(context.Context, *IdRequest) (*Response, error)
	mustEmbedUnimplementedStaffServiceServer()
}

// UnimplementedStaffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) Create(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStaffServiceServer) Get(context.Context, *IdRequest) (*GetStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStaffServiceServer) List(context.Context, *ListStaffRequest) (*ListStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStaffServiceServer) Update(context.Context, *UpdateStaffRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStaffServiceServer) Delete(context.Context, *IdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff_service.StaffService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Create(ctx, req.(*CreateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff_service.StaffService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Get(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff_service.StaffService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).List(ctx, req.(*ListStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff_service.StaffService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Update(ctx, req.(*UpdateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff_service.StaffService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff_service.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StaffService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StaffService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StaffService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff.proto",
}
