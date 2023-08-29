package main

import (
	hw "GRPC_DEMO2/helloworld"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := hw.NewGreeting_CustomerClient(conn)
	det:=&hw.AddCustomerDetails{
		Name: "Nithi",
		Age: "21",
	}
	addResp,err:=client.AddCustomer(context.Background(),det)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addResp)
	tasksResp, err := client.GetCustomer(context.Background(), &hw.Empty{})
    if err != nil {
        log.Fatal("failed to retrive tasks:%v", err)
    }
    fmt.Println("Tasks:")
    for _, task := range tasksResp.GetTasks1(){
        fmt.Printf("ID:%s,Title:%s,Completed: %v\n", task.Id, task.Title, task)
    }

}
