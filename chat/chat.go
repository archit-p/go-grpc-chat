package chat

import (
	"context"
	"log"
)

// extends the ChatServiceClient interface
type Server struct {
}

// caller for sending messages
func (s *Server) SayHello(cxt context.Context, in *Message) (*Message, error) {
	log.Printf("Received message from client: %v", in.Body)
	return &Message{Body: "Hello from server!"}, nil
}
