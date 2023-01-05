package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_service_user_service/protoc"
	"time"
)

const (
	address = "192.168.1.21:8081"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
