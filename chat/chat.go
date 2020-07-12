package chat

import (
	"log"

	"context"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Msgs, error) {
	// type Reply []Message

	var reply []*Message

	for i := 0; i < 5; i++ {
		reply = append(reply, &Message{Body: "Hi Client"})
	}
	log.Printf("Recieved client: %s", message.Body)
	return &Msgs{Messages: reply}, nil
}
