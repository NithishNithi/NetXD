package main

import (
	hw "GRPC_DEMO2/helloworld"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := hw.NewGreeting_CustomerClient(conn)

}
