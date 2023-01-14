package logic

import (
	"context"
	"github.com/reation/micro_service_user_service/config"
	"github.com/reation/micro_service_user_service/find_user/internal/svc"
	"github.com/reation/micro_service_user_service/model"
	"github.com/reation/micro_service_user_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserMemberByIdLogic {
	return &CheckUserMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserMemberByIdLogic) CheckUserMemberById(in *protoc.CheckUserMemberRequest) (*protoc.CheckUserMemberResponse, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UserModel.GetUserInfoByID(l.ctx, in.GetId())
	switch err {
	case nil:
		return &protoc.CheckUserMemberResponse{States: config.STATES_NORMAL, IsMember: result.IsMember}, nil
	case model.ErrNotFound:
		return &protoc.CheckUserMemberResponse{States: config.STATES_EMPTY, IsMember: model.MEMBER_NO}, nil
	default:
		return &protoc.CheckUserMemberResponse{States: config.STATES_ERROR, IsMember: model.MEMBER_NO}, nil
	}
}
