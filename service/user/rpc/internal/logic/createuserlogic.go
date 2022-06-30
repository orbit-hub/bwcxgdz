package logic

import (
	"bwcxgdz/dousheng/common/cryptx"
	"bwcxgdz/dousheng/service/user/model"
	"bwcxgdz/dousheng/service/user/rpc/internal/svc"
	"bwcxgdz/dousheng/service/user/rpc/user"
	"context"

	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserInfoRequest) (*user.UserInfoResponse, error) {

	_, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Name)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}
	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &user.UserInfoResponse{
			Id:            newUser.Id,
			Name:          newUser.Name,
			FollowCount:   newUser.FollowCount,
			FollowerCount: newUser.FollowerCount,
		}, nil
	}

	return &user.UserInfoResponse{}, nil
}
