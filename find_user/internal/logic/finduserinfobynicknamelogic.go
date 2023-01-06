package logic

import (
	"context"
	"micro_service_user_service/config"
	"micro_service_user_service/find_user/internal/svc"
	"micro_service_user_service/model"
	"micro_service_user_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserInfoByNickNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserInfoByNickNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserInfoByNickNameLogic {
	return &FindUserInfoByNickNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserInfoByNickNameLogic) FindUserInfoByNickName(in *protoc.ByNickNameRequest) (*protoc.ByNickNameResponse, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UserModel.FindUserInfoByNickName(l.ctx, in.Nickname)
	switch err {
	case nil:
		resp := protoc.UserData{
			Id:       result.Id,
			UserName: result.UserName,
			Nickname: result.Nickname,
			Sex:      result.Sex,
			Birthday: result.Birthday.Format("2006-01-02"),
			States:   result.State,
			IsMember: result.IsMember,
		}
		return &protoc.ByNickNameResponse{States: config.STATES_NORMAL, UserData: &resp}, nil
	case model.ErrNotFound:
		return &protoc.ByNickNameResponse{States: config.STATES_EMPTY, UserData: &protoc.UserData{}}, nil
	default:
		return &protoc.ByNickNameResponse{States: config.STATES_ERROR, UserData: &protoc.UserData{}}, nil
	}
}
