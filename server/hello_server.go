package main

import (
	"context"
	"fmt"
	"hello-world-grpc-in-go/gen/go/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server struct must embed hello.UnimplementedHelloServiceServer
type Server struct {
	hello.UnimplementedHelloServiceServer
	grpcServer *grpc.Server
}

func (s *Server) SayHello(_ context.Context, request *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("server: replying")
	return &hello.HelloReply{Message: "hello mr"}, nil
}

var port = "8080"

func NewServer() *Server {
	instance := &Server{}
	instance.grpcServer = grpc.NewServer()
	return instance
}

// create server
// register the grpc server with implementation
// listen on tcp port 8080
func main() {
	helloGrpcServer := NewServer()
	hello.RegisterHelloServiceServer(helloGrpcServer.grpcServer, helloGrpcServer)

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("error in listening on port %v | %v", port, err)
	} else {
		log.Printf("consumer svc listening on port %v", port)
	}

	err = helloGrpcServer.startServer(l)
	if err != nil {
		log.Fatalf("error starting server %v", err)
	}
}

func (s *Server) startServer(listener net.Listener) error {
	if err := s.grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("unable to start server | %w", err)
	}
	return nil
}

func (s *Server) stopServer() {
	s.grpcServer.GracefulStop()
}
