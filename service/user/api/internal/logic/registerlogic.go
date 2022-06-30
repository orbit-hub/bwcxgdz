package logic

import (
	"bwcxgdz/dousheng/common/jwtx"
	"bwcxgdz/dousheng/service/user/rpc/user"
	"context"
	"time"

	"bwcxgdz/dousheng/service/user/api/internal/svc"
	"bwcxgdz/dousheng/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &user.CreateUserInfoRequest{
		Name:     req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     res.Id,
		Token:      accessToken,
	}, nil
}
