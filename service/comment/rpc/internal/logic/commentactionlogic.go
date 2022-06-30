package logic

import (
	"bwcxgdz/dousheng/service/comment/model"
	"bwcxgdz/dousheng/service/comment/rpc/internal/svc"
	"bwcxgdz/dousheng/service/comment/rpc/proto"
	"context"
	"time"

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
	comment := model.Comment{
		UserId:    in.UserId,
		VideoId:   in.VideoID,
		Content:   in.CommentText,
		CreatedAt: time.Now(),
	}
	var id int64 = 0
	if in.ActionType == 1 {
		rs, err := l.svcCtx.CommentModel.Insert(l.ctx, &comment)
		if err != nil {
			return nil, err
		}
		id, err = rs.LastInsertId()
	} else {
		err := l.svcCtx.CommentModel.Delete(l.ctx, in.CommentId)
		if err != nil {
			return nil, err
		}
	}
	return &proto.ActionResponse{
		CommentId: id,
	}, nil
}
