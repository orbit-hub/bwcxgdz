package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListByIdLogic {
	return &GetCommentListByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListByIdLogic) GetCommentListById(in *proto.VideoIdRequest) (*proto.CommentListResponse, error) {
	// todo: add your logic here and delete this line

	return &proto.CommentListResponse{}, nil
}
