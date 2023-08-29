package main

import (
	hw "GRPC_DEMO1/helloworld"
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	mu    sync.Mutex
	tasks map[string]*hw.Task
	hw.UnimplementedTaskServiceServer
}

func generateId() string {
	return "hi"
}

func (s *server) AddTask(ctx context.Context, req *hw.Task) (*hw.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	taskID := generateId()
	req.Id = taskID 
	s.tasks[taskID] = req
	return &hw.TaskResponse{Id: taskID}, nil
}

func (s *server) GetTasks(ctx context.Context, req *hw.Empty) (*hw.TaskList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	tasks := make([]*hw.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return &hw.TaskList{Tasks: tasks}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	taskServer := &server{
		tasks: make(map[string]*hw.Task),
	}
	hw.RegisterTaskServiceServer(s, taskServer)

	fmt.Println("Server listening on port 4040...")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
