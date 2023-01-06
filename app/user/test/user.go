package main

import (
	"context"
	"fmt"
	v1 "galileo/api/user/v1"
	"google.golang.org/grpc"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func init() {
	var err error
	conn, err = grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		panic("grpc link error" + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}
func main() {
	defer conn.Close()
	testCreateUser()
}

func testCreateUser() {
	rsp, err := userClient.CreateUser(context.Background(), &v1.CreateUserRequest{
		Phone:    fmt.Sprintf("1234567890%d", 1),
		Password: "test1234",
		Username: "wentao3.chen",
		Email:    "wentao3.chen@gmail.com",
	})
	if err != nil {
		panic("grpc create user fail: " + err.Error())
	}
	fmt.Println(rsp)
}

func testGetUser() {
	rep, err := userClient.GetUser(context.Background(), &v1.GetUserRequest{Id: 7})
	if err != nil {
		panic("grpc get user fail: " + err.Error())
	}
	fmt.Println(rep)
}
