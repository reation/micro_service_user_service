package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		UserTable struct {
			User   string
			Passwd string
			Addr   string
			Port   string
			DBName string
		}
	}
	RedisCluster struct {
		Host []string
		Pass string
	}
}
