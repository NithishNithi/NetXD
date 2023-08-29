package main

import (
	"fmt"
	"time"
	// "runtime"
	// 	"sync"
)

func main() {

	counter := 0

	const num = 15
	

	for i := 0; i < num; i++ {
		go func() {
			temp := counter
			temp++
			counter = temp
			
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("count:", counter)
}