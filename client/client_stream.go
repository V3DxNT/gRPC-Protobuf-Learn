package main

import (
	"context"
	"log"
	"time"

	pb "github.com/V3DxNT/gRPC-Protobuf-Learn/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming Started")

	stream, err := client.SayHelloClientStreaming(context.TODO())
	if err != nil {
		log.Fatalf("SayHelloClientStreaming: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("SayHelloClientStreaming: %v", err)
		}
		log.Printf("Sent the request: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf(" Error : Client Streaming finished %v", err)
	}
	log.Printf("Client Streaming Finished")
	log.Printf(" Hi %v", res.Messages)

}
