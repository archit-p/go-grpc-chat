package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/archit-p/go-grpc-chat/chat"
)

func main() {
	// start listening on port 51343
	lis, err := net.Listen("tcp", ":51343")
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	// register it for calls to SayHello
	chat.RegisterChatServiceServer(grpcServer, &s)

	// start listening for clients
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}
