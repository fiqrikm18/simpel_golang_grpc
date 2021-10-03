package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "simple_grpc/pkg/services/user_service"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	HandleError(err)
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)
	ctx, cancle := context.WithTimeout(context.Background(), time.Second * 30)
	defer cancle()

	_, err = c.AddUser(ctx, &pb.User{
		Name: "fiqri2",
		Address: "asd",
		Id: 1,
		Email: "asd@email.com",
	})
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		log.Fatalf("Error occur with error: %v", err)
	}
}

