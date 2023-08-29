package main
import (
	"fmt"
)

func myfun(channel chan string){
	for i:=0;i<100;i++{
		channel <- "hi how are u"
	}
	close(channel)
}

func main() {
	c:=make(chan string,8)

	go myfun(c) 
	counter:=0
	for{
		res,ok := <- c
		counter++

		if !ok {
			fmt.Println("Channel is closed",ok)
			break
		}

		fmt.Println(counter)
		fmt.Println(res,ok)
	}
}
