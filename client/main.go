package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "gorpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Sarvesh"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("Greeting:", resp.GetMessage())
}
