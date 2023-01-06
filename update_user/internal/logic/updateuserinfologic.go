package logic

import (
	"context"
	"micro_service_user_service/config"
	"micro_service_user_service/model"
	"time"

	"micro_service_user_service/protoc"
	"micro_service_user_service/update_user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *protoc.UpdateUserInfoRequest) (*protoc.UpdateUserInfoResponse, error) {
	// todo: add your logic here and delete this line]
	_, err := l.svcCtx.UserModel.GetUserInfoByID(l.ctx, in.GetId())
	if err != model.ErrNotFound {
		return &protoc.UpdateUserInfoResponse{States: config.STATES_EMPTY, Id: 0}, nil
	}
	if err != nil {
		return &protoc.UpdateUserInfoResponse{States: config.STATES_ERROR, Id: 0}, nil
	}

	birthday, _ := time.Parse("2006-01-02", in.GetBirthday())
	id, err := l.svcCtx.UserModel.UpdateUserInfoByID(l.ctx, in.GetId(), in.GetNickname(), in.GetSex(), birthday)
	if err != nil {
		return &protoc.UpdateUserInfoResponse{States: config.STATES_ERROR, Id: 0}, nil
	}

	return &protoc.UpdateUserInfoResponse{States: config.STATES_NORMAL, Id: id}, nil
}
