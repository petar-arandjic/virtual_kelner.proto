// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: identity/v1/identity.proto

package identitypb

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

// BusinessServiceClient is the client API for BusinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessServiceClient interface {
	InviteBusiness(ctx context.Context, in *InviteBusinessRequest, opts ...grpc.CallOption) (*InviteBusinessResponse, error)
	AcceptInvitationAsBusiness(ctx context.Context, in *AcceptInvitationAsBusinessRequest, opts ...grpc.CallOption) (*AcceptInvitationAsBusinessResponse, error)
	GetBusinessById(ctx context.Context, in *GetBusinessByIdRequest, opts ...grpc.CallOption) (*GetBusinessByIdResponse, error)
}

type businessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessServiceClient(cc grpc.ClientConnInterface) BusinessServiceClient {
	return &businessServiceClient{cc}
}

func (c *businessServiceClient) InviteBusiness(ctx context.Context, in *InviteBusinessRequest, opts ...grpc.CallOption) (*InviteBusinessResponse, error) {
	out := new(InviteBusinessResponse)
	err := c.cc.Invoke(ctx, "/identity.v1.BusinessService/InviteBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) AcceptInvitationAsBusiness(ctx context.Context, in *AcceptInvitationAsBusinessRequest, opts ...grpc.CallOption) (*AcceptInvitationAsBusinessResponse, error) {
	out := new(AcceptInvitationAsBusinessResponse)
	err := c.cc.Invoke(ctx, "/identity.v1.BusinessService/AcceptInvitationAsBusiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) GetBusinessById(ctx context.Context, in *GetBusinessByIdRequest, opts ...grpc.CallOption) (*GetBusinessByIdResponse, error) {
	out := new(GetBusinessByIdResponse)
	err := c.cc.Invoke(ctx, "/identity.v1.BusinessService/GetBusinessById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServiceServer is the server API for BusinessService service.
// All implementations must embed UnimplementedBusinessServiceServer
// for forward compatibility
type BusinessServiceServer interface {
	InviteBusiness(context.Context, *InviteBusinessRequest) (*InviteBusinessResponse, error)
	AcceptInvitationAsBusiness(context.Context, *AcceptInvitationAsBusinessRequest) (*AcceptInvitationAsBusinessResponse, error)
	GetBusinessById(context.Context, *GetBusinessByIdRequest) (*GetBusinessByIdResponse, error)
	mustEmbedUnimplementedBusinessServiceServer()
}

// UnimplementedBusinessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBusinessServiceServer struct {
}

func (UnimplementedBusinessServiceServer) InviteBusiness(context.Context, *InviteBusinessRequest) (*InviteBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) AcceptInvitationAsBusiness(context.Context, *AcceptInvitationAsBusinessRequest) (*AcceptInvitationAsBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvitationAsBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) GetBusinessById(context.Context, *GetBusinessByIdRequest) (*GetBusinessByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessById not implemented")
}
func (UnimplementedBusinessServiceServer) mustEmbedUnimplementedBusinessServiceServer() {}

// UnsafeBusinessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessServiceServer will
// result in compilation errors.
type UnsafeBusinessServiceServer interface {
	mustEmbedUnimplementedBusinessServiceServer()
}

func RegisterBusinessServiceServer(s grpc.ServiceRegistrar, srv BusinessServiceServer) {
	s.RegisterService(&BusinessService_ServiceDesc, srv)
}

func _BusinessService_InviteBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).InviteBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.v1.BusinessService/InviteBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).InviteBusiness(ctx, req.(*InviteBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_AcceptInvitationAsBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptInvitationAsBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).AcceptInvitationAsBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.v1.BusinessService/AcceptInvitationAsBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).AcceptInvitationAsBusiness(ctx, req.(*AcceptInvitationAsBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_GetBusinessById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).GetBusinessById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/identity.v1.BusinessService/GetBusinessById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).GetBusinessById(ctx, req.(*GetBusinessByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BusinessService_ServiceDesc is the grpc.ServiceDesc for BusinessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BusinessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.v1.BusinessService",
	HandlerType: (*BusinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InviteBusiness",
			Handler:    _BusinessService_InviteBusiness_Handler,
		},
		{
			MethodName: "AcceptInvitationAsBusiness",
			Handler:    _BusinessService_AcceptInvitationAsBusiness_Handler,
		},
		{
			MethodName: "GetBusinessById",
			Handler:    _BusinessService_GetBusinessById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/v1/identity.proto",
}
