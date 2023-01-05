package logic

import (
	"context"
	"micro_service_user_service/model"
	"micro_service_user_service/tool"

	"micro_service_user_service/protoc"
	"micro_service_user_service/user_login/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *protoc.LoginRequest) (*protoc.LoginResponse, error) {
	// todo: add your logic here and delete this line
	var returnInfo protoc.UserInfo
	userInfo, err := l.svcCtx.UserModel.FindUserInfoByUserName(l.ctx, in.GetUserName())
	if err != nil {
		switch err {
		case model.ErrNotFound:
			return &protoc.LoginResponse{States: 2, UserInfo: &returnInfo}, nil
		default:
			return &protoc.LoginResponse{States: 11, UserInfo: &returnInfo}, nil
		}
	}
	if userInfo.Password != tool.MD5(in.GetPassword()) {
		return &protoc.LoginResponse{States: 3, UserInfo: &returnInfo}, err
	}
	returnInfo.Id = userInfo.Id
	returnInfo.UserName = userInfo.UserName
	returnInfo.IsMember = userInfo.IsMember

	return &protoc.LoginResponse{States: 1, UserInfo: &returnInfo}, nil
}
