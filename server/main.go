package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "simple_grpc/pkg/services/user_service"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (us *UserServer) AddUser(ctx context.Context, data *pb.User) (*pb.User, error) {
	log.Println("User with data: " + data.GetName())
	return &pb.User{
		Name: data.GetName(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	HandleError(err)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserServer{})
	log.Println("gRPC server is listening on localhost:5000")
	err = grpcServer.Serve(listener)
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		log.Fatalf("Error occur with error: %v", err)
	}
}
