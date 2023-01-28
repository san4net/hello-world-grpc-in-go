package main

import (
	"context"
	"fmt"
	"hello-world-grpc-in-go/gen/go/hello"
	"log"

	"google.golang.org/grpc"
)

func CreateClient() hello.HelloServiceClient {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Printf("error client connection %v", err)
	}
	return hello.NewHelloServiceClient(conn)
}

// client
// 1. first create hello.HelloServiceClient with url as "localhost:8080"
// 2. call hello.HelloServiceClient.SayHello and get reply from server
func main() {
	reply, err := CreateClient().SayHello(context.Background(), &hello.HelloRequest{Name: "this is santosh"})

	if err != nil {
		log.Printf("error requesting server %v", err)
	}

	fmt.Printf("got reply from server %v", reply)
}
