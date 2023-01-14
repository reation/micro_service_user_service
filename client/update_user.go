package main

import (
	"context"
	"github.com/reation/micro_service_user_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	UpdateUserAddress = "192.168.1.21:8082"
)

func main() {
	UpdateUserInfo()
	//UpdateUserMember()
	//UpdateUserPassWord()
}

func UpdateUserInfo() {
	conn, err := grpc.Dial(UpdateUserAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUpdateUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateUserInfo(ctx, &protoc.UpdateUserInfoRequest{Id: 2, Nickname: "levi1", Sex: 0, Birthday: "1986-10-28"})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("last_id: %d", r.Id)
}

func UpdateUserMember() {
	conn, err := grpc.Dial(UpdateUserAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUpdateUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateUserMember(ctx, &protoc.UpdateUserMemberRequest{Id: 2, IsMember: 1})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("last_id: %d", r.Id)
}

func UpdateUserPassWord() {
	conn, err := grpc.Dial(UpdateUserAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUpdateUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateUserPassWord(ctx, &protoc.UpdateUserPassWordRequest{Id: 1, Password: "123456"})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("last_id: %d", r.Id)
}
