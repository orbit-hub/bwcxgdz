package logic

import (
	"bwcxgdz/dousheng/service/video/api/pack"
	"bwcxgdz/dousheng/service/video/rpc/video"
	"context"

	"bwcxgdz/dousheng/service/video/api/internal/svc"
	"bwcxgdz/dousheng/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (resp *types.VideoListResponse, err error) {
	res, err := l.svcCtx.VideoRpc.Feed(l.ctx, &video.FeedRequest{
		LatestTime: req.LatestTime,
	})
	if err != nil {
		return nil, err
	}
	infos := pack.VideoInfos(res.VideoList)
	return &types.VideoListResponse{
		StatusCode:    0,
		StatusMsg:     "success",
		NextTime:      1000000000,
		VideoInfoList: infos,
	}, nil
}
