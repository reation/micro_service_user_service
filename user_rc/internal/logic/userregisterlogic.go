package logic

import (
	"context"
	"github.com/reation/micro_service_user_service/config"
	"github.com/reation/micro_service_user_service/model"
	"github.com/reation/micro_service_user_service/tool"
	"time"

	"github.com/reation/micro_service_user_service/protoc"
	"github.com/reation/micro_service_user_service/user_rc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *protoc.RegisterRequest) (*protoc.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.FindUserInfoByUserName(l.ctx, in.GetUserName())
	if err == nil {
		return &protoc.RegisterResponse{States: config.STATES_CHECK_REPEAT, Id: 0}, nil
	}

	if err != model.ErrNotFound {
		return &protoc.RegisterResponse{States: config.STATES_ERROR, Id: 0}, nil
	}

	birthday, _ := time.Parse("2006-01-02", in.GetBirthday())
	nowTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	userInfo := model.User{
		UserName:   in.GetUserName(),
		Password:   tool.MD5(in.GetPassword()),
		Nickname:   in.GetNickname(),
		Sex:        in.GetSex(),
		Birthday:   birthday,
		State:      1,
		IsMember:   0,
		UpdateTime: nowTime,
		CreateTime: nowTime,
	}

	result, err := l.svcCtx.UserModel.Insert(l.ctx, &userInfo)
	if err != nil {
		return &protoc.RegisterResponse{States: config.STATES_ERROR, Id: 0}, nil
	}
	lastId, _ := result.LastInsertId()
	return &protoc.RegisterResponse{States: config.STATES_NORMAL, Id: lastId}, nil
}
