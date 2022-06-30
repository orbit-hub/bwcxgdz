package logic

import (
	"bwcxgdz/dousheng/service/favorite/rpc/internal/svc"
	"bwcxgdz/dousheng/service/favorite/rpc/proto"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteVideoListByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoListByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoListByIdLogic {
	return &GetFavoriteVideoListByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteVideoListByIdLogic) GetFavoriteVideoListById(in *proto.IdRequest) (*proto.VideoIdListResponse, error) {
	var resp proto.VideoIdListResponse
	ids, err := l.svcCtx.FavoriteModel.FindListByUserId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	resp.VideoIdList = append(resp.VideoIdList, ids...)
	return &resp, nil
}
