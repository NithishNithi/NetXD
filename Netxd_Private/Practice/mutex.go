package main

import (
	"fmt"
	"sync"
	"time"
)
var m sync.Mutex
var count int = 0

func test(){
	for i:=0;i<2;i++{
	m.Lock()
	count++
	fmt.Println("1",count)
	m.Unlock()
	}

}

func test1(){
	
	for i:=0;i<2;i++{
		m.Lock()
		count++
		fmt.Println("2",count)
		m.Unlock()
		}
}
func main(){
	
	go test()
	go test1()

	time.Sleep(1 * time.Second)


}