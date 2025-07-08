package main

import (
	"context"
	"log"
	"net"
	pb "test_user_pb/gen"

	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	log.Printf("Got request to user with ID: %d", req.GetId())

	var user = &pb.User{
		Id:    req.Id,
		Name:  "Alex",
		Email: "alex111@yandex.ru",
	}

	return &pb.UserResponse{User: user}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("net.Listen error: %v", err)
	}

	grpcServ := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServ, &userServiceServer{})

	log.Printf("gRPC serv is starting to: 50051")

	if err := grpcServ.Serve(listener); err != nil {
		log.Fatalf("Err run gRPC serv: %v", err)
	}
}
