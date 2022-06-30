package logic

import (
	"bwcxgdz/dousheng/common/cryptx"
	"bwcxgdz/dousheng/service/user/model"
	"context"

	"google.golang.org/grpc/status"

	"bwcxgdz/dousheng/service/user/rpc/internal/svc"
	"bwcxgdz/dousheng/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	res, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user.LoginResponse{
		UserId: res.Id,
	}, nil
}
