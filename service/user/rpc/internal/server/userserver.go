// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"bwcxgdz/dousheng/service/user/rpc/internal/logic"
	"bwcxgdz/dousheng/service/user/rpc/internal/svc"
	"bwcxgdz/dousheng/service/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUserInfo(ctx context.Context, in *user.IdRequest) (*user.UserListResponse, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServer) CreateUser(ctx context.Context, in *user.CreateUserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
