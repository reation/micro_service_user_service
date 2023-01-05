package logic

import (
	"context"
	"micro_service_user_service/model"

	"micro_service_user_service/protoc"
	"micro_service_user_service/user_rc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCancelLogic {
	return &UserCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCancelLogic) UserCancel(in *protoc.CancelRequest) (*protoc.CancelResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UserModel.UpdateUserStates(l.ctx, in.GetId(), model.STATES_CANCEL)
	if err != nil {
		return &protoc.CancelResponse{States: 0}, nil
	}
	return &protoc.CancelResponse{States: 1}, nil
}
