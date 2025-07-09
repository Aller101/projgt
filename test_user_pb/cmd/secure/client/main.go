package main

import (
	"context"
	"log"
	"time"

	pb "test_user_pb/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	tlsCreds, err := generateTLSCreds()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(
		tlsCreds,
	))
	if err != nil {
		log.Fatalf("dont connect to serv: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	defer cancel()

	req := &pb.UserRequest{Id: 1}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("err compl req: %v", err)
	}
	log.Printf("Got user: ID=%d, Name=%s, Email=%s\n", res.User.Id, res.User.Name, res.User.Email)

}

func generateTLSCreds() (credentials.TransportCredentials, error) {

	certFile := "ca.crt"

	return credentials.NewClientTLSFromFile(certFile, "")
}
