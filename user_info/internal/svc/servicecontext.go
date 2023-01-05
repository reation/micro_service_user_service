package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro_service_user_service/model"
	"micro_service_user_service/user_info/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.UserTable.DataSourceName)),
	}
}