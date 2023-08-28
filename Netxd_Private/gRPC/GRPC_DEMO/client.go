package main

import (
	hw "GRPC_DEMO/helloworld"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)


func main(){
	conn,err:=grpc.Dial("localhost:4040",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client:= hw.NewGreeterClient(conn)

	name:="Nithi"
	
	response,err:= client.SayHello(context.Background(),&hw.HelloRequest{Name:name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Message)
}