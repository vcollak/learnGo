package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"LearnGo/Patterns/3.4_gRPC/chat"
)

func main() {

	//delare a grpc client and connect to the server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	//register a new client with chat server
	c := chat.NewChatServiceClient(conn)

	//message to be sent
	message := chat.Message{
		Body: "Hello from the client",
	}

	//call the remote method on te grpc server
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	//get the server repsonse
	log.Printf("Response from Server: %s", response.Body)

}
