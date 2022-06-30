package logic

import (
	"context"

	"bwcxgdz/dousheng/service/video/rpc/internal/svc"
	"bwcxgdz/dousheng/service/video/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoListLogic {
	return &PublishVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoListLogic) PublishVideoList(in *proto.IdRequest) (*proto.VideoListResponse, error) {
	list, err := l.svcCtx.VideoModel.GetVideoListByUserId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	responses := make([]*proto.VideoInfoResponse, 0)
	for _, video := range list {
		resp := proto.VideoInfoResponse{
			Id:            video.Id,
			AuthorId:      video.AuthorId,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			Title:         video.Title,
		}
		responses = append(responses, &resp)
	}
	return &proto.VideoListResponse{VideoList: responses}, nil
}
