package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/archit-p/go-grpc-chat/chat"
)

func main() {
	var conn *grpc.ClientConn

	// create a new connection with server
	conn, err := grpc.Dial(":51343", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	// try to send message
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from client!"})
	if err != nil {
		log.Fatalf("Error in server response: %v", err)
	}
	log.Printf("Response from server: %v", response.Body)
}
