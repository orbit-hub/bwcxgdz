package logic

import (
	"context"
	"time"

	"bwcxgdz/dousheng/service/video/rpc/internal/svc"
	"bwcxgdz/dousheng/service/video/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *proto.FeedRequest) (*proto.FeedListResponse, error) {
	var now int64
	if in.LatestTime != 0 {
		now = in.LatestTime
	} else {
		now = time.Now().Unix()
	}
	list, err := l.svcCtx.VideoModel.GetVideoList(l.ctx, now)
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
	return &proto.FeedListResponse{
		NextTime:  now,
		VideoList: responses,
	}, nil
}
