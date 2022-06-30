// Code generated by goctl. DO NOT EDIT!
// Source: video.proto

package video

import (
	"context"

	"bwcxgdz/dousheng/service/video/rpc/proto"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FeedListResponse  = proto.FeedListResponse
	FeedRequest       = proto.FeedRequest
	IdRequest         = proto.IdRequest
	PublishResponse   = proto.PublishResponse
	VideoInfoRequest  = proto.VideoInfoRequest
	VideoInfoResponse = proto.VideoInfoResponse
	VideoListResponse = proto.VideoListResponse

	Video interface {
		Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedListResponse, error)
		Publish(ctx context.Context, in *VideoInfoRequest, opts ...grpc.CallOption) (*PublishResponse, error)
		PublishVideoList(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoListResponse, error)
		GetVideoInfoById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoInfoResponse, error)
	}

	defaultVideo struct {
		cli zrpc.Client
	}
)

func NewVideo(cli zrpc.Client) Video {
	return &defaultVideo{
		cli: cli,
	}
}

func (m *defaultVideo) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedListResponse, error) {
	client := proto.NewVideoClient(m.cli.Conn())
	return client.Feed(ctx, in, opts...)
}

func (m *defaultVideo) Publish(ctx context.Context, in *VideoInfoRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	client := proto.NewVideoClient(m.cli.Conn())
	return client.Publish(ctx, in, opts...)
}

func (m *defaultVideo) PublishVideoList(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoListResponse, error) {
	client := proto.NewVideoClient(m.cli.Conn())
	return client.PublishVideoList(ctx, in, opts...)
}

func (m *defaultVideo) GetVideoInfoById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*VideoInfoResponse, error) {
	client := proto.NewVideoClient(m.cli.Conn())
	return client.GetVideoInfoById(ctx, in, opts...)
}
