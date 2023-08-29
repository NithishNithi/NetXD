package main

import (
	"fmt"
	"time"
)

func helloworld(){
	fmt.Println("hi world")
}

func main(){
	fmt.Println("start")
	helloworld()
	go helloworld()
	time.Sleep(2 * time.Second)
	fmt.Println("end")
}