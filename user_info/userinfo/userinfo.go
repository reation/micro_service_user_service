// Code generated by goctl. DO NOT EDIT.
// Source: user_info.proto

package userinfo

import (
	"context"

	"github.com/reation/micro_service_user_service/protoc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	NormalUserInfoRequest  = protoc.NormalUserInfoRequest
	NormalUserInfoResponse = protoc.NormalUserInfoResponse
	UserDetail             = protoc.UserDetail
	UserInfoRequest        = protoc.UserInfoRequest
	UserInfoResponse       = protoc.UserInfoResponse

	UserInfo interface {
		GetUserInfoByID(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		GetNormalUserInfoByID(ctx context.Context, in *NormalUserInfoRequest, opts ...grpc.CallOption) (*NormalUserInfoResponse, error)
	}

	defaultUserInfo struct {
		cli zrpc.Client
	}
)

func NewUserInfo(cli zrpc.Client) UserInfo {
	return &defaultUserInfo{
		cli: cli,
	}
}

func (m *defaultUserInfo) GetUserInfoByID(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := protoc.NewUserInfoClient(m.cli.Conn())
	return client.GetUserInfoByID(ctx, in, opts...)
}

func (m *defaultUserInfo) GetNormalUserInfoByID(ctx context.Context, in *NormalUserInfoRequest, opts ...grpc.CallOption) (*NormalUserInfoResponse, error) {
	client := protoc.NewUserInfoClient(m.cli.Conn())
	return client.GetNormalUserInfoByID(ctx, in, opts...)
}
