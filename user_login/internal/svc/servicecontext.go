package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro_service_user_service/model"
	"micro_service_user_service/user_login/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.UserTable.DataSourceName)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}