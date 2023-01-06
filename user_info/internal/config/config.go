package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		UserTable struct {
			DataSourceName string
		}
	}
	RedisCluster struct {
		Host []string
		Pass string
	}
}
