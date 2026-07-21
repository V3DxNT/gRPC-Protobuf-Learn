package main

import (
	"context"
	"io"
	"log"

	pb "github.com/V3DxNT/gRPC-Protobuf-Learn/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming Started ")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("SayHelloServerStreaming failed: %v", err)
		return
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while Streamign SayHelloServerStreaming failed: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming End ")

}
