package logic

import (
	"bwcxgdz/dousheng/service/video/api/pack"
	"bwcxgdz/dousheng/service/video/rpc/video"
	"context"

	"bwcxgdz/dousheng/service/video/api/internal/svc"
	"bwcxgdz/dousheng/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListRequest) (resp *types.VideoListResponse, err error) {
	res, err := l.svcCtx.VideoRpc.PublishVideoList(l.ctx, &video.IdRequest{Id: req.UserId})
	if err != nil {
		return nil, err
	}
	infos := pack.VideoInfos(res.VideoList)
	return &types.VideoListResponse{
		StatusCode:    0,
		StatusMsg:     "success",
		VideoInfoList: infos,
	}, nil
}
