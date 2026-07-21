package main

import (
	"io"
	"log"

	pb "github.com/V3DxNT/gRPC-Protobuf-Learn/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		str, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})

		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println("Got Request with names : %v", str.Name)
		messages = append(messages, str.Name)
	}

	return nil
}
