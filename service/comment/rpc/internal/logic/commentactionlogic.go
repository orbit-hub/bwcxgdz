package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *proto.ActionRequest) (*proto.ActionResponse, error) {
	// todo: add your logic here and delete this line

	return &proto.ActionResponse{}, nil
}
