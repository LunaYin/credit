// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package entity

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

// CreditServiceClient is the client API for CreditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreditServiceClient interface {
	AddCredit(ctx context.Context, in *AddUserCredit, opts ...grpc.CallOption) (*UserCredit, error)
	GetCredit(ctx context.Context, in *GetUserCredit, opts ...grpc.CallOption) (*UserCredit, error)
}

type creditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreditServiceClient(cc grpc.ClientConnInterface) CreditServiceClient {
	return &creditServiceClient{cc}
}

func (c *creditServiceClient) AddCredit(ctx context.Context, in *AddUserCredit, opts ...grpc.CallOption) (*UserCredit, error) {
	out := new(UserCredit)
	err := c.cc.Invoke(ctx, "/entity.CreditService/AddCredit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditServiceClient) GetCredit(ctx context.Context, in *GetUserCredit, opts ...grpc.CallOption) (*UserCredit, error) {
	out := new(UserCredit)
	err := c.cc.Invoke(ctx, "/entity.CreditService/GetCredit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreditServiceServer is the server API for CreditService service.
// All implementations must embed UnimplementedCreditServiceServer
// for forward compatibility
type CreditServiceServer interface {
	AddCredit(context.Context, *AddUserCredit) (*UserCredit, error)
	GetCredit(context.Context, *GetUserCredit) (*UserCredit, error)
	mustEmbedUnimplementedCreditServiceServer()
}

// UnimplementedCreditServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCreditServiceServer struct {
}

func (UnimplementedCreditServiceServer) AddCredit(context.Context, *AddUserCredit) (*UserCredit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCredit not implemented")
}
func (UnimplementedCreditServiceServer) GetCredit(context.Context, *GetUserCredit) (*UserCredit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredit not implemented")
}
func (UnimplementedCreditServiceServer) mustEmbedUnimplementedCreditServiceServer() {}

// UnsafeCreditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreditServiceServer will
// result in compilation errors.
type UnsafeCreditServiceServer interface {
	mustEmbedUnimplementedCreditServiceServer()
}

func RegisterCreditServiceServer(s grpc.ServiceRegistrar, srv CreditServiceServer) {
	s.RegisterService(&CreditService_ServiceDesc, srv)
}

func _CreditService_AddCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserCredit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditServiceServer).AddCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.CreditService/AddCredit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditServiceServer).AddCredit(ctx, req.(*AddUserCredit))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditService_GetCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCredit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditServiceServer).GetCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.CreditService/GetCredit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditServiceServer).GetCredit(ctx, req.(*GetUserCredit))
	}
	return interceptor(ctx, in, info, handler)
}

// CreditService_ServiceDesc is the grpc.ServiceDesc for CreditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entity.CreditService",
	HandlerType: (*CreditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCredit",
			Handler:    _CreditService_AddCredit_Handler,
		},
		{
			MethodName: "GetCredit",
			Handler:    _CreditService_GetCredit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credit.proto",
}
