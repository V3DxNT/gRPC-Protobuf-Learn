package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/V3DxNT/gRPC-Protobuf-Learn/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("BiDirectional Calling")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Failed to SayHelloBidirectional: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatalf("SayHelloBidirectional error: %v", err)
			}
			log.Printf("%v", message)
		}
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil && err != io.EOF {
			log.Fatalf("SayHelloBidirectional error: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Println("BiDirectional Calling")

}
