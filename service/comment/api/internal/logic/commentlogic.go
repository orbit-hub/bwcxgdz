package logic

import (
	"bwcxgdz/dousheng/service/comment/rpc/proto"
	"context"
	"encoding/json"

	"bwcxgdz/dousheng/service/comment/api/internal/svc"
	"bwcxgdz/dousheng/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.CommentRequest) (resp *types.CommentResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CommentRpc.CommentAction(l.ctx, &proto.ActionRequest{
		UserId:      uid,
		VideoID:     req.VideoId,
		CommentText: req.CommentText,
		ActionType:  req.ActionType,
		CommentId:   req.CommentId,
	})
	if err != nil {
		return nil, err
	}
	info := types.CommentInfo{
		Id:      res.CommentId,
		Content: req.CommentText,
	}
	return &types.CommentResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Comment:    []*types.CommentInfo{&info},
	}, nil
}
