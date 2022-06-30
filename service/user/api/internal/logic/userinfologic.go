package logic

import (
	"bwcxgdz/dousheng/service/user/api/internal/svc"
	"bwcxgdz/dousheng/service/user/api/internal/types"
	"bwcxgdz/dousheng/service/user/rpc/user"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.IdRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	userInfo := types.UserInfo{
		Id:            res.Data[0].Id,
		Name:          res.Data[0].Name,
		FollowCount:   res.Data[0].FollowCount,
		FollowerCount: res.Data[0].FollowerCount,
		IsFollow:      res.Data[0].IsFollow,
	}
	userInfos := []*types.UserInfo{&userInfo}
	return &types.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserInfo:   userInfos,
	}, nil
}
