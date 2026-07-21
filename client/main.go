package main

import (
	"log"

	pb "github.com/V3DxNT/gRPC-Protobuf-Learn/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Vedant", "Asthana", "Ved"},
	}

	//callSayHelloServerStreaming(client, names)
	log.Println("Starting the Client Calling here")
	callSayHelloClientStream(client, names)

	//callSayHelloBidrectionalStream(client,names)

	//callSayHello(client)
}
