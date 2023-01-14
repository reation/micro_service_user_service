package main

import (
	"context"
	"github.com/reation/micro_service_user_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	UserRcAddress = "192.168.1.21:8081"
)

func main() {

}

func userCancel() {
	conn, err := grpc.Dial(UserRcAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUserRegisterOrCancelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserCancel(ctx, &protoc.CancelRequest{Id: 1})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
}

func userRegister() {
	conn, err := grpc.Dial(UserRcAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUserRegisterOrCancelClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserRegister(ctx, &protoc.RegisterRequest{UserName: "levi", Password: "6676211", Nickname: "table", Sex: 0, Birthday: "1986-10-29"})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("id: %d", r.GetId())
}
