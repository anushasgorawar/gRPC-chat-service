package main

import (
	"log"
	"net"

	"github.com/anushasgorawar/gRPC-chat-service/chat"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
	defer listener.Close()
	log.Println("Server started on localhost:8080")

	s := chat.Server{}

	gRPCserver := grpc.NewServer()

	chat.RegisterChatServiceServer(gRPCserver, &s)
	if err := gRPCserver.Serve(listener); err != nil {
		log.Fatal("Could not start grpc server: ", err)
	}
}
