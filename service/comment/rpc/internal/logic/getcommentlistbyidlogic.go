package logic

import (
	"bwcxgdz/dousheng/service/comment/rpc/internal/svc"
	"bwcxgdz/dousheng/service/comment/rpc/proto"
	"context"

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
	res, err := l.svcCtx.CommentModel.FindByVideoId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	infos := make([]*proto.CommentInfoResponse, 0)
	for _, r := range res {
		infos = append(infos, &proto.CommentInfoResponse{
			Id:      r.Id,
			UserId:  r.UserId,
			Content: r.Content,
		})
	}
	return &proto.CommentListResponse{
		Data: infos,
	}, nil
}
