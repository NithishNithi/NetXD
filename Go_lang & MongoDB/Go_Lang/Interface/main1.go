package main

import "fmt"

type Shape interface {
	Square() int
	Rectangle() int
}

type Parameter1 struct {
	length  int
	breadth int
}

type Parameter2 struct {
	length  int
	breadth int
	width   int
}

func (tmp Parameter1) Square() int {
	t1 := tmp.length * tmp.breadth
	return t1
}

func (tmp Parameter1) Rectangle() int {
	t1 := 2 * (tmp.length + tmp.breadth)
	return t1
}

func (tmp Parameter2) Square() int {
	t1 := tmp.length * tmp.breadth * tmp.width
	return t1
}

func main() {
	var s Shape = Parameter1{length: 10, breadth: 10}
	fmt.Println(s.Square())
	var s1 Shape = 
}
