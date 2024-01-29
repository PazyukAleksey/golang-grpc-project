package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpcTest4/proto"
	"grpcTest4/service"
)

func main() {
	lis, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	myService := &service.UserServiceServer{}

	pb.RegisterUserServiceServer(grpcServer, myService)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}

}
