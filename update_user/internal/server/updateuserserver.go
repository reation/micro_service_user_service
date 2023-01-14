// Code generated by goctl. DO NOT EDIT.
// Source: update_user.proto

package server

import (
	"context"

	"github.com/reation/micro_service_user_service/protoc"
	"github.com/reation/micro_service_user_service/update_user/internal/logic"
	"github.com/reation/micro_service_user_service/update_user/internal/svc"
)

type UpdateUserServer struct {
	svcCtx *svc.ServiceContext
	protoc.UnimplementedUpdateUserServer
}

func NewUpdateUserServer(svcCtx *svc.ServiceContext) *UpdateUserServer {
	return &UpdateUserServer{
		svcCtx: svcCtx,
	}
}

func (s *UpdateUserServer) UpdateUserInfo(ctx context.Context, in *protoc.UpdateUserInfoRequest) (*protoc.UpdateUserInfoResponse, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

func (s *UpdateUserServer) UpdateUserMember(ctx context.Context, in *protoc.UpdateUserMemberRequest) (*protoc.UpdateUserMemberResponse, error) {
	l := logic.NewUpdateUserMemberLogic(ctx, s.svcCtx)
	return l.UpdateUserMember(in)
}

func (s *UpdateUserServer) UpdateUserPassWord(ctx context.Context, in *protoc.UpdateUserPassWordRequest) (*protoc.UpdateUserPassWordResponse, error) {
	l := logic.NewUpdateUserPassWordLogic(ctx, s.svcCtx)
	return l.UpdateUserPassWord(in)
}
