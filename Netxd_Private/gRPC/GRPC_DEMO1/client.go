package main

import (
	hw "GRPC_DEMO1/helloworld"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main(){
	conn,err:=grpc.Dial("localhost:4040",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err)
	}
	defer conn.Close()
	client := hw.NewTaskServiceClient(conn)
	task := &hw.Task{
        Title: "Buy groceries",
    }
    addResp, err := client.AddTask(context.Background(), task)
    if err != nil {
        log.Fatal("failed to add task: %v", err)
    }
    fmt.Println(addResp)
    //retrive tasks
    tasksResp, err := client.GetTasks(context.Background(), &hw.Empty{})
    if err != nil {
        log.Fatal("failed to retrive tasks:%v", err)
    }
    fmt.Println("Tasks:")
    for _, task := range tasksResp.Tasks {
        fmt.Printf("ID:%s,Title:%s,Completed: %v\n", task.Id, task.Title, task)
    }
}