package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
	defer listen.Close()
	log.Println("Server started on localhost:8085")

	for {
		log.Println("inside for loop")
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Could not start server: ", err)
			continue
		}

		go handleconn(conn)
		log.Println("post go handleconn")

	}
}

func handleconn(conn net.Conn) {
	defer conn.Close()
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 11\r\n" +
		"\r\n" +
		"Hello There"
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Fatal("error writing: ", err)
		return
	}
}
