package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

//SayHello provides the logic for the SayHello ChatService
//defined in chat.proto
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Hello from Server!"}, nil

}
