package logic

import (
	"bwcxgdz/dousheng/service/comment/rpc/comment"
	"context"
	"time"

	"bwcxgdz/dousheng/service/comment/api/internal/svc"
	"bwcxgdz/dousheng/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	res, err := l.svcCtx.CommentRpc.GetCommentListById(l.ctx, &comment.VideoIdRequest{
		Id: req.VideoId,
	})
	infos := make([]*types.CommentInfo, 0)
	for _, com := range res.GetData() {
		infos = append(infos, &types.CommentInfo{
			Id:         com.Id,
			Content:    com.Content,
			CreateDate: time.Now().String(),
		})
	}

	return &types.CommentListResponse{
		Comment: infos,
	}, nil
}
