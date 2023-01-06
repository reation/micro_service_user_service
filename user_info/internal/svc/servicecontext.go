package svc

import (
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro_service_user_service/model"
	"micro_service_user_service/user_info/internal/config"
	"runtime"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
	RedisModel goRedis.ClusterClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisModel := GetRedisClusterConn(c.RedisCluster.Host, c.RedisCluster.Pass)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.UserTable.DataSourceName), redisModel),
	}
}

func GetRedisClusterConn(host []string, pass string) goRedis.ClusterClient {
	redisCluster := goRedis.NewClusterClient(&goRedis.ClusterOptions{
		Addrs:           host,
		Password:        pass,
		ReadOnly:        false,
		PoolSize:        20 * runtime.NumCPU(),
		MinIdleConns:    10,
		MaxRetries:      0,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,
		DialTimeout:     5 * time.Second,
		ReadTimeout:     3 * time.Second,
		WriteTimeout:    3 * time.Second,
		PoolTimeout:     4 * time.Second,
		IdleTimeout:     5 * time.Minute,
		MaxConnAge:      0 * time.Second,
	})

	return *redisCluster
}
