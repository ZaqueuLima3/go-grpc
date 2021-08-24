package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ZaqueuLima3/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jhon",
		Email: "j2@j2.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jhon",
		Email: "j2@j2.com",
	}
	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, "-", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "0",
			Name:  "Jhon",
			Email: "j2@j2.com",
		},
		{
			Id:    "1",
			Name:  "Jose",
			Email: "j3@j3.com",
		},
		{
			Id:    "2",
			Name:  "Joa",
			Email: "j4@j4.com",
		},
	}
	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error reciving response: %v", err)
	}
	fmt.Println("users: ", res)
}
