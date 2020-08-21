package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"LearnGo/Patterns/3.4_gRPC/chat"
)

func main() {

	//Listen using TCP on the port 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	//use package chat to create a new server
	s := chat.Server{}

	//create a new grpc server
	grpcServer := grpc.NewServer()

	//register the grpcserver with the chat server
	chat.RegisterChatServiceServer(grpcServer, &s)

	//let grpc listen on the listener we declared before
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}

}
