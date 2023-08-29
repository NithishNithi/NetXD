package main

import (
	"fmt"
)

//A sender without a reciver leads to Deadlock


func routine(a chan int){
	//reciver
	fmt.Println(100 + <-a)
}

func main() {
	fmt.Println("Start !")

	a:=make(chan int)
	go routine(a)
	//sender
	a<-214
	fmt.Printf("%T",a)
	
}