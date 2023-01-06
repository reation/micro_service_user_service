package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		UserTable struct {
			DataSourceName string
			User           string
			Passwd         string
			Addr           string
			Port           string
			DBName         string
		}
	}
}
