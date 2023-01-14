package logic

import (
	"context"
	"github.com/reation/micro_service_user_service/config"
	"github.com/reation/micro_service_user_service/model"

	"github.com/reation/micro_service_user_service/protoc"
	"github.com/reation/micro_service_user_service/user_info/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNormalUserInfoByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNormalUserInfoByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNormalUserInfoByIDLogic {
	return &GetNormalUserInfoByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNormalUserInfoByIDLogic) GetNormalUserInfoByID(in *protoc.NormalUserInfoRequest) (*protoc.NormalUserInfoResponse, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserModel.GetNormalUserInfoByID(l.ctx, in.GetId())
	if err != nil {
		switch err {
		case model.ErrNotFound:
			return &protoc.NormalUserInfoResponse{States: config.STATES_EMPTY, UserInfo: &protoc.UserDetail{}}, err
		default:
			return &protoc.NormalUserInfoResponse{States: config.STATES_ERROR, UserInfo: &protoc.UserDetail{}}, err
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

	return &protoc.NormalUserInfoResponse{States: config.STATES_NORMAL, UserInfo: &userDetail}, nil
}
