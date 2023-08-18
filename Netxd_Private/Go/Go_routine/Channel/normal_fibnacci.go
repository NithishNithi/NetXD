package main

import (
	"fmt"
	"time"
)

func main() {
	t1:=time.Now()
	a:=0
	b:=1
	var c int
	fmt.Print(a,b," ")
	for i:=0;i<10-2;i++{
		c = a+b
		a = b
		b = c
		fmt.Print(c," ")
	}
	t2:=time.Now().Sub(t1)
	fmt.Println(t2)
}