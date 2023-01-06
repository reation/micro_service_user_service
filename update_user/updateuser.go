package main

import (
	"flag"
	"fmt"

	"micro_service_user_service/protoc"
	"micro_service_user_service/update_user/internal/config"
	"micro_service_user_service/update_user/internal/server"
	"micro_service_user_service/update_user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/updateuser.yaml", "the config file")

//var configFile2 = flag.String("f", "../etc/updateuser.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//conf.MustLoad(*configFile2, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		protoc.RegisterUpdateUserServer(grpcServer, server.NewUpdateUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
