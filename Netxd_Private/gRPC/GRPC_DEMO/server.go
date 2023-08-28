package main

import (
	"context"
	"fmt"
	"net"

	hw "GRPC_DEMO/helloworld" 
	"google.golang.org/grpc"
)

type server struct {
	hw.UnimplementedGreeterServer 
}

func (s *server) SayHello(ctx context.Context, req *hw.HelloRequest) (*hw.HelloReply, error) {
	return &hw.HelloReply{
		Message: fmt.Sprintf("Hello, %s", req.Name), 
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := grpc.NewServer() // It creates a gRPC instance
	hw.RegisterGreeterServer(s, &server{})

	fmt.Println("Server listening on Port: 4040")
	err1 := s.Serve(lis)
	if err1 != nil {
		fmt.Println(err1)
	}
}
