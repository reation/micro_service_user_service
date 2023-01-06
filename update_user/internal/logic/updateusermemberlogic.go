package logic

import (
	"context"
	"micro_service_user_service/config"
	"micro_service_user_service/model"

	"micro_service_user_service/protoc"
	"micro_service_user_service/update_user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserMemberLogic {
	return &UpdateUserMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserMemberLogic) UpdateUserMember(in *protoc.UpdateUserMemberRequest) (*protoc.UpdateUserMemberResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.GetUserInfoByID(l.ctx, in.GetId())
	if err == model.ErrNotFound {
		return &protoc.UpdateUserMemberResponse{States: config.STATES_EMPTY, Id: 0}, nil
	}
	if err != nil {
		return &protoc.UpdateUserMemberResponse{States: config.STATES_ERROR, Id: 0}, nil
	}
	id, err := l.svcCtx.UserModel.UpdateUserMemberByID(l.ctx, in.GetId(), in.GetIsMember())
	if err != nil {
		return &protoc.UpdateUserMemberResponse{States: config.STATES_ERROR, Id: 0}, nil
	}

	return &protoc.UpdateUserMemberResponse{States: config.STATES_NORMAL, Id: id}, nil
}
