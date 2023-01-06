package model

import (
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, redisModel goRedis.ClusterClient) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, redisModel),
	}
}
