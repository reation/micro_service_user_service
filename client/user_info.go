package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_service_user_service/protoc"
	"time"
)

const (
	UserInfoAddress = "192.168.1.21:7080"
)

func main() {
	//type a struct {
	//	Id       int64  `db:"id"`
	//	UserName string `db:"user_name"` // 用户名
	//	Password string `db:"password"`  // 密码
	//	Nickname string `db:"nickname"`  // 昵称
	//	Sex      int64  `db:"sex"`       // 性别 0:男 1:女
	//}
	//
	//b := a{
	//	Id:       3,
	//	UserName: "rea",
	//	Password: "344444",
	//}
	//
	//data, _ := json.Marshal(b)
	//var resp a
	//json.Unmarshal([]byte(data), &resp)
	//fmt.Println(resp)
	getUserInfoByID()
	//getNormalUserInfoByID()
}

func getUserInfoByID() {
	conn, err := grpc.Dial(UserInfoAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUserInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserInfoByID(ctx, &protoc.UserInfoRequest{Id: 2})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Println(r.UserInfo)
}

func getNormalUserInfoByID() {
	conn, err := grpc.Dial(UserInfoAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewUserInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetNormalUserInfoByID(ctx, &protoc.NormalUserInfoRequest{Id: 1})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Println(r.UserInfo)
}
