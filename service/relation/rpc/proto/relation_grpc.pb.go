// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: relation.proto

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

// RelationClient is the client API for Relation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationClient interface {
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	GetFollowList(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserIdListResponse, error)
	GetFollowerList(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserIdListResponse, error)
}

type relationClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationClient(cc grpc.ClientConnInterface) RelationClient {
	return &relationClient{cc}
}

func (c *relationClient) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/relation.relation/RelationAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GetFollowList(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserIdListResponse, error) {
	out := new(UserIdListResponse)
	err := c.cc.Invoke(ctx, "/relation.relation/GetFollowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationClient) GetFollowerList(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserIdListResponse, error) {
	out := new(UserIdListResponse)
	err := c.cc.Invoke(ctx, "/relation.relation/GetFollowerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServer is the server API for Relation service.
// All implementations must embed UnimplementedRelationServer
// for forward compatibility
type RelationServer interface {
	RelationAction(context.Context, *RelationActionRequest) (*ActionResponse, error)
	GetFollowList(context.Context, *IdRequest) (*UserIdListResponse, error)
	GetFollowerList(context.Context, *IdRequest) (*UserIdListResponse, error)
	mustEmbedUnimplementedRelationServer()
}

// UnimplementedRelationServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServer struct {
}

func (UnimplementedRelationServer) RelationAction(context.Context, *RelationActionRequest) (*ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedRelationServer) GetFollowList(context.Context, *IdRequest) (*UserIdListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowList not implemented")
}
func (UnimplementedRelationServer) GetFollowerList(context.Context, *IdRequest) (*UserIdListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerList not implemented")
}
func (UnimplementedRelationServer) mustEmbedUnimplementedRelationServer() {}

// UnsafeRelationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServer will
// result in compilation errors.
type UnsafeRelationServer interface {
	mustEmbedUnimplementedRelationServer()
}

func RegisterRelationServer(s grpc.ServiceRegistrar, srv RelationServer) {
	s.RegisterService(&Relation_ServiceDesc, srv)
}

func _Relation_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relation/RelationAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).RelationAction(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GetFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GetFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relation/GetFollowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GetFollowList(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relation_GetFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServer).GetFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.relation/GetFollowerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServer).GetFollowerList(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Relation_ServiceDesc is the grpc.ServiceDesc for Relation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.relation",
	HandlerType: (*RelationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RelationAction",
			Handler:    _Relation_RelationAction_Handler,
		},
		{
			MethodName: "GetFollowList",
			Handler:    _Relation_GetFollowList_Handler,
		},
		{
			MethodName: "GetFollowerList",
			Handler:    _Relation_GetFollowerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}