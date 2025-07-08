package main

import (
	"context"
	"log"
	"time"

	pb "test_user_pb/gen"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dont connect to serv: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UserRequest{Id: 1}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("err compl req: %v", err)
	}
	log.Printf("Got user: ID=%d, Name=%s, Email=%s\n", res.User.Id, res.User.Name, res.User.Email)

}
