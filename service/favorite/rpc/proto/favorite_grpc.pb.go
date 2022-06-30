// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: favorite.proto

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

// FavoriteClient is the client API for Favorite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteClient interface {
	GetFavoriteVideoListById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoIdListResponse, error)
	FavoriteAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
}

type favoriteClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteClient(cc grpc.ClientConnInterface) FavoriteClient {
	return &favoriteClient{cc}
}

func (c *favoriteClient) GetFavoriteVideoListById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoIdListResponse, error) {
	out := new(VideoIdListResponse)
	err := c.cc.Invoke(ctx, "/favorite.favorite/GetFavoriteVideoListById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) FavoriteAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/favorite.favorite/FavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServer is the server API for Favorite service.
// All implementations must embed UnimplementedFavoriteServer
// for forward compatibility
type FavoriteServer interface {
	GetFavoriteVideoListById(context.Context, *IdRequest) (*VideoIdListResponse, error)
	FavoriteAction(context.Context, *ActionRequest) (*ActionResponse, error)
	mustEmbedUnimplementedFavoriteServer()
}

// UnimplementedFavoriteServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServer struct {
}

func (UnimplementedFavoriteServer) GetFavoriteVideoListById(context.Context, *IdRequest) (*VideoIdListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteVideoListById not implemented")
}
func (UnimplementedFavoriteServer) FavoriteAction(context.Context, *ActionRequest) (*ActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteServer) mustEmbedUnimplementedFavoriteServer() {}

// UnsafeFavoriteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServer will
// result in compilation errors.
type UnsafeFavoriteServer interface {
	mustEmbedUnimplementedFavoriteServer()
}

func RegisterFavoriteServer(s grpc.ServiceRegistrar, srv FavoriteServer) {
	s.RegisterService(&Favorite_ServiceDesc, srv)
}

func _Favorite_GetFavoriteVideoListById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).GetFavoriteVideoListById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.favorite/GetFavoriteVideoListById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).GetFavoriteVideoListById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.favorite/FavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).FavoriteAction(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Favorite_ServiceDesc is the grpc.ServiceDesc for Favorite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Favorite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "favorite.favorite",
	HandlerType: (*FavoriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFavoriteVideoListById",
			Handler:    _Favorite_GetFavoriteVideoListById_Handler,
		},
		{
			MethodName: "FavoriteAction",
			Handler:    _Favorite_FavoriteAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
