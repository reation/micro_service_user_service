package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro_service_user_service/model"
	"micro_service_user_service/update_user/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSourceUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", c.Mysql.UserTable.User, c.Mysql.UserTable.Passwd, c.Mysql.UserTable.Addr, c.Mysql.UserTable.Port, c.Mysql.UserTable.DBName)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(dataSourceUrl)),
	}
}
