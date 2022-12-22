package main

import (
	"log"
	"net"

	pb "github.com/Christomesh/go-grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcserver := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcserver, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcserver.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err)
	}

}
