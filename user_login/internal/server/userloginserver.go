// Code generated by goctl. DO NOT EDIT.
// Source: user_login.proto

package server

import (
	"context"

	"micro_service_user_service/protoc"
	"micro_service_user_service/user_login/internal/logic"
	"micro_service_user_service/user_login/internal/svc"
)

type UserLoginServer struct {
	svcCtx *svc.ServiceContext
	protoc.UnimplementedUserLoginServer
}

func NewUserLoginServer(svcCtx *svc.ServiceContext) *UserLoginServer {
	return &UserLoginServer{
		svcCtx: svcCtx,
	}
}

func (s *UserLoginServer) UserLogin(ctx context.Context, in *protoc.LoginRequest) (*protoc.LoginResponse, error) {
	l := logic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}
