package logic

import (
	"context"
	"micro_service_user_service/config"
	"micro_service_user_service/model"
	"micro_service_user_service/tool"

	"micro_service_user_service/protoc"
	"micro_service_user_service/update_user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPassWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPassWordLogic {
	return &UpdateUserPassWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserPassWordLogic) UpdateUserPassWord(in *protoc.UpdateUserPassWordRequest) (*protoc.UpdateUserPassWordResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.GetUserInfoByID(l.ctx, in.GetId())
	if err != model.ErrNotFound {
		return &protoc.UpdateUserPassWordResponse{States: config.STATES_EMPTY, Id: 0}, nil
	}
	if err != nil {
		return &protoc.UpdateUserPassWordResponse{States: config.STATES_ERROR, Id: 0}, nil
	}
	password := tool.MD5(in.GetPassword())
	id, err := l.svcCtx.UserModel.UpdateUserPassWordByID(l.ctx, in.GetId(), password)
	if err != nil {
		return &protoc.UpdateUserPassWordResponse{States: config.STATES_ERROR, Id: 0}, nil
	}
	return &protoc.UpdateUserPassWordResponse{States: config.STATES_NORMAL, Id: id}, nil
}
