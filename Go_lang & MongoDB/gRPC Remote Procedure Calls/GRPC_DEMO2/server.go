package main

import (
	hw "GRPC_DEMO2/helloworld"
	"context"
	"fmt"
	"net"

	"sync"

	"github.com/google/uuid"

	"google.golang.org/grpc"
)

type server struct {
	hw.UnimplementedGreeting_CustomerServer
	details map[string]*hw.AddCustomerDetails
	mutex   sync.Mutex
}

func (s *server) AddCustomerDetails(ctx context.Context, req *hw.AddCustomerDetails) (*hw.CustomerResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cusID := uuid.New().String() // Generate a new UUID
	req.Id = cusID
	s.details[cusID] = req
	return &hw.CustomerResponse{Id: cusID}, nil
}

func (s *server) GetCustomer(ctx context.Context, req *hw.Empty) (*hw.Customerlist, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	temp := make([]*hw.AddCustomerDetails, 0, len(s.details))

	for _, i := range s.details {
		temp = append(temp, i)
	}
	return &hw.Customerlist{Tasks1: temp}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	hw.RegisterGreeting_CustomerServer(s, &server{details: make(map[string]*hw.AddCustomerDetails)})

	fmt.Println("Server listening on port 5050...")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}

}
