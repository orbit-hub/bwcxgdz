package logic

import (
	"bwcxgdz/dousheng/service/video/model"
	"context"
	"time"

	"bwcxgdz/dousheng/service/video/rpc/internal/svc"
	"bwcxgdz/dousheng/service/video/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *proto.VideoInfoRequest) (*proto.PublishResponse, error) {
	_, err := l.svcCtx.VideoModel.Insert(l.ctx, &model.Video{
		AuthorId:    in.UserId,
		PlayUrl:     in.PlayUrl,
		CoverUrl:    in.CoverUrl,
		Title:       in.Title,
		PublishTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &proto.PublishResponse{
		Success: true,
	}, nil
}
