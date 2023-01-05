// Code generated by goctl. DO NOT EDIT.
// Source: user_info.proto

package server

import (
	"context"

	"micro_service_user_service/protoc"
	"micro_service_user_service/user_info/internal/logic"
	"micro_service_user_service/user_info/internal/svc"
)

type UserInfoServer struct {
	svcCtx *svc.ServiceContext
	protoc.UnimplementedUserInfoServer
}

func NewUserInfoServer(svcCtx *svc.ServiceContext) *UserInfoServer {
	return &UserInfoServer{
		svcCtx: svcCtx,
	}
}

func (s *UserInfoServer) GetUserInfoByID(ctx context.Context, in *protoc.UserInfoRequest) (*protoc.UserInfoResponse, error) {
	l := logic.NewGetUserInfoByIDLogic(ctx, s.svcCtx)
	return l.GetUserInfoByID(in)
}

func (s *UserInfoServer) GetNormalUserInfoByID(ctx context.Context, in *protoc.NormalUserInfoRequest) (*protoc.NormalUserInfoResponse, error) {
	l := logic.NewGetNormalUserInfoByIDLogic(ctx, s.svcCtx)
	return l.GetNormalUserInfoByID(in)
}