package logic

import (
	"bwcxgdz/dousheng/service/user/model"
	"bwcxgdz/dousheng/service/user/rpc/internal/svc"
	"bwcxgdz/dousheng/service/user/rpc/user"
	"context"

	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.IdRequest) (*user.UserListResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	userInfoRsp := ModelToRsponse(res)
	rsp := &user.UserListResponse{}
	rsp.Data = append(rsp.Data, &userInfoRsp)
	return rsp, nil
}

func ModelToRsponse(userInfo *model.User) user.UserInfoResponse {
	userInfoRsp := user.UserInfoResponse{
		Id:            userInfo.Id,
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      true,
	}
	return userInfoRsp
}
