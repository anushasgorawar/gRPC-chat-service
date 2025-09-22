package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
	defer listener.Close()
	log.Println("Server started on localhost:8080")

	gRPCserver := grpc.NewServer()
	if err := gRPCserver.Serve(listener); err != nil {
		log.Fatal("Could not start grpc server: ", err)
	}
}
