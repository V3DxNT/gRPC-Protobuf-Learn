package main

import (
	"log"
	"net"

	pb "github.com/V3DxNT/gRPC-Protobuf-Learn/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	// 1. Open the raw TCP pipe (Layer 4)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the Server: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the Server: %v", err)
	}
}
