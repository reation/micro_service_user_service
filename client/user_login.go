package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_service_user_service/protoc"
	"time"
)

const (
	address = "192.168.1.21:8080"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUserLoginClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserLogin(ctx, &protoc.LoginRequest{UserName: "reation", Password: "198610290"})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("ID: %d", r.GetUserInfo().GetId())
	log.Printf("user_name: %s", r.GetUserInfo().GetUserName())
	log.Printf("is_member: %d", r.GetUserInfo().GetIsMember())
}
