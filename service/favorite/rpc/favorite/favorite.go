// Code generated by goctl. DO NOT EDIT!
// Source: favorite.proto

package favorite

import (
	"context"

	"bwcxgdz/dousheng/service/favorite/rpc/proto"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRequest       = proto.ActionRequest
	ActionResponse      = proto.ActionResponse
	IdRequest           = proto.IdRequest
	VideoIdListResponse = proto.VideoIdListResponse

	Favorite interface {
		GetFavoriteVideoListById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoIdListResponse, error)
		FavoriteAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	}

	defaultFavorite struct {
		cli zrpc.Client
	}
)

func NewFavorite(cli zrpc.Client) Favorite {
	return &defaultFavorite{
		cli: cli,
	}
}

func (m *defaultFavorite) GetFavoriteVideoListById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoIdListResponse, error) {
	client := proto.NewFavoriteClient(m.cli.Conn())
	return client.GetFavoriteVideoListById(ctx, in, opts...)
}

func (m *defaultFavorite) FavoriteAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	client := proto.NewFavoriteClient(m.cli.Conn())
	return client.FavoriteAction(ctx, in, opts...)
}
