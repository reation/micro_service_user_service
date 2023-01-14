package main

import (
	"context"
	"github.com/reation/micro_service_user_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	FindUserAddress = "192.168.1.21:7081"
)

func main() {
	FindUserInfoByUserName()
	FindUserInfoByNickName()
	CheckUserMemberById()
}

func FindUserInfoByUserName() {
	conn, err := grpc.Dial(FindUserAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewFindUserInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FindUserInfoByUserName(ctx, &protoc.ByUserNameRequest{UserName: "reation"})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Println(r.UserData)
}

func FindUserInfoByNickName() {
	conn, err := grpc.Dial(FindUserAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewFindUserInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FindUserInfoByNickName(ctx, &protoc.ByNickNameRequest{Nickname: "枫"})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Println(r.UserData)
}

func CheckUserMemberById() {
	conn, err := grpc.Dial(FindUserAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewFindUserInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CheckUserMemberById(ctx, &protoc.CheckUserMemberRequest{Id: 1})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("is_member: %d", r.GetIsMember())
}
