// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: sale.proto

package sale_service

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

// SaleServiceClient is the client API for SaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SaleServiceClient interface {
	Create(ctx context.Context, in *CreateSaleRequest, opts ...grpc.CallOption) (*CreateSaleResponse, error)
	Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetSaleResponse, error)
	List(ctx context.Context, in *ListSaleRequest, opts ...grpc.CallOption) (*ListSaleResponse, error)
	Update(ctx context.Context, in *UpdateSaleRequest, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error)
}

type saleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSaleServiceClient(cc grpc.ClientConnInterface) SaleServiceClient {
	return &saleServiceClient{cc}
}

func (c *saleServiceClient) Create(ctx context.Context, in *CreateSaleRequest, opts ...grpc.CallOption) (*CreateSaleResponse, error) {
	out := new(CreateSaleResponse)
	err := c.cc.Invoke(ctx, "/sale_service.SaleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saleServiceClient) Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*GetSaleResponse, error) {
	out := new(GetSaleResponse)
	err := c.cc.Invoke(ctx, "/sale_service.SaleService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saleServiceClient) List(ctx context.Context, in *ListSaleRequest, opts ...grpc.CallOption) (*ListSaleResponse, error) {
	out := new(ListSaleResponse)
	err := c.cc.Invoke(ctx, "/sale_service.SaleService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saleServiceClient) Update(ctx context.Context, in *UpdateSaleRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sale_service.SaleService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saleServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sale_service.SaleService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaleServiceServer is the server API for SaleService service.
// All implementations must embed UnimplementedSaleServiceServer
// for forward compatibility
type SaleServiceServer interface {
	Create(context.Context, *CreateSaleRequest) (*CreateSaleResponse, error)
	Get(context.Context, *IdRequest) (*GetSaleResponse, error)
	List(context.Context, *ListSaleRequest) (*ListSaleResponse, error)
	Update(context.Context, *UpdateSaleRequest) (*Response, error)
	Delete(context.Context, *IdRequest) (*Response, error)
	mustEmbedUnimplementedSaleServiceServer()
}

// UnimplementedSaleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSaleServiceServer struct {
}

func (UnimplementedSaleServiceServer) Create(context.Context, *CreateSaleRequest) (*CreateSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSaleServiceServer) Get(context.Context, *IdRequest) (*GetSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSaleServiceServer) List(context.Context, *ListSaleRequest) (*ListSaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSaleServiceServer) Update(context.Context, *UpdateSaleRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSaleServiceServer) Delete(context.Context, *IdRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSaleServiceServer) mustEmbedUnimplementedSaleServiceServer() {}

// UnsafeSaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SaleServiceServer will
// result in compilation errors.
type UnsafeSaleServiceServer interface {
	mustEmbedUnimplementedSaleServiceServer()
}

func RegisterSaleServiceServer(s grpc.ServiceRegistrar, srv SaleServiceServer) {
	s.RegisterService(&SaleService_ServiceDesc, srv)
}

func _SaleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.SaleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaleServiceServer).Create(ctx, req.(*CreateSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.SaleService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaleServiceServer).Get(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.SaleService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaleServiceServer).List(ctx, req.(*ListSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.SaleService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaleServiceServer).Update(ctx, req.(*UpdateSaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sale_service.SaleService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaleServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SaleService_ServiceDesc is the grpc.ServiceDesc for SaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sale_service.SaleService",
	HandlerType: (*SaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SaleService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SaleService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SaleService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SaleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SaleService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sale.proto",
}