package chat

import (
	"log"

	"golang.org/x/net/context"
)

func main() {

}

type Server struct {
}

// mustEmbedUnimplementedChatServiceServer implements ChatServiceServer.
func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	panic("unimplemented")
}

func (s *Server) SayHello(ctx context.Context, msg *Message) (*Message, error) {
	log.Printf("Recieved Message from client: %s", msg.Body)
	return &Message{Body: "hello from the server"}, nil
}
