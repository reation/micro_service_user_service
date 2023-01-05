package logic

import (
	"context"
	"micro_service_user_service/config"
	"micro_service_user_service/model"

	"micro_service_user_service/find_user/internal/svc"
	"micro_service_user_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserInfoByUserNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserInfoByUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserInfoByUserNameLogic {
	return &FindUserInfoByUserNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserInfoByUserNameLogic) FindUserInfoByUserName(in *protoc.ByUserNameRequest) (*protoc.ByUserNameResponse, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UserModel.FindUserInfoByUserName(l.ctx, in.UserName)
	if err == model.ErrNotFound {
		return &protoc.ByUserNameResponse{States: config.STATES_EMPTY, UserData: &protoc.UserData{}}, nil
	}

	if err != nil {
		return &protoc.ByUserNameResponse{States: config.STATES_ERROR, UserData: &protoc.UserData{}}, nil
	}
	userInfo := protoc.UserData{
		Id:       result.Id,
		UserName: result.UserName,
		Nickname: result.Nickname,
		Sex:      result.Sex,
		Birthday: result.Birthday.Format("2006-01-02"),
		IsMember: result.IsMember,
		States:   result.State,
	}

	return &protoc.ByUserNameResponse{States: config.STATES_NORMAL, UserData: &userInfo}, nil
}
