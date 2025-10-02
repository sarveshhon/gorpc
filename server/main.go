package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gorpc/proto"

	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &greeterServer{})

	fmt.Println("Server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
