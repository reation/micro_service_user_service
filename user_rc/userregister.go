package main

import (
	"flag"
	"fmt"

	"github.com/reation/micro_service_user_service/protoc"
	"github.com/reation/micro_service_user_service/user_rc/internal/config"
	"github.com/reation/micro_service_user_service/user_rc/internal/server"
	"github.com/reation/micro_service_user_service/user_rc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/userregister.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		protoc.RegisterUserRegisterOrCancelServer(grpcServer, server.NewUserRegisterOrCancelServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
