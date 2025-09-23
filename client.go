package main

import (
	"log"

	"github.com/anushasgorawar/gRPC-chat-service/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect: ", err)
	}
	defer conn.Close()

	chatclient := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Hello from client",
	}

	res, err := chatclient.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatal("Could not connect: ", err)
	}
	log.Printf("Response from server: %v", res.Body)
}
