package logic

import (
	"context"
	"micro_service_user_service/config"
	"micro_service_user_service/model"

	"micro_service_user_service/protoc"
	"micro_service_user_service/user_info/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIDLogic {
	return &GetUserInfoByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIDLogic) GetUserInfoByID(in *protoc.UserInfoRequest) (*protoc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserModel.GetUserInfoByID(l.ctx, in.GetId())
	if err != nil {
		switch err {
		case model.ErrNotFound:
			return &protoc.UserInfoResponse{States: config.STATES_EMPTY, UserInfo: &protoc.UserDetail{}}, nil
		default:
			return &protoc.UserInfoResponse{States: config.STATES_ERROR, UserInfo: &protoc.UserDetail{}}, nil
		}
	}
	userDetail := protoc.UserDetail{
		Id:       userInfo.Id,
		UserName: userInfo.UserName,
		Nickname: userInfo.Nickname,
		Sex:      userInfo.Sex,
		Birthday: userInfo.Birthday.Format("2006-01-02"),
		States:   userInfo.State,
		IsMember: userInfo.IsMember,
	}

	return &protoc.UserInfoResponse{States: config.STATES_NORMAL, UserInfo: &userDetail}, nil
}
