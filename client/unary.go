package main

import (
	"context"
	"log"
	"time"

	pb "github.com/V3DxNT/gRPC-Protobuf-Learn/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("SayHello Unary err: %v", err)
	}
	log.Printf("%v", res.Message)
}
