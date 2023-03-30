package main

import (
	"context"
	"fmt"
	v1 "galileo/api/project/v1"
	"google.golang.org/grpc"
)

var (
	projctClient v1.ProjectClient
	conn         *grpc.ClientConn
)

func init() {
	var err error
	conn, err = grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic("grpc link errResponse" + err.Error())
	}
	projctClient = v1.NewProjectClient(conn)
}

func main() {
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("grpc close errResponse" + err.Error())
		}
	}(conn)
	testCreateProject()
}

func testCreateProject() {
	rsp, err := projctClient.CreateProject(context.Background(), &v1.CreateProjectRequest{
		Name:       "project",
		Identifier: "Myriad",
	})
	if err != nil {
		panic("grpc create project fail: " + err.Error())
	}
	fmt.Println("test result: ", rsp)
}
