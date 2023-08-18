package main

import "fmt"

type node struct{
	data int
	next *node
	prev *node
}
func createNode(val int)*node{
	return &node{
	data:val,
	prev:nil,
	next:nil,
	}
}

func insertFirst(val int,head **node){
	var temp = createNode(val)
	temp.next= *head
	(*head).prev = temp
	*head = temp
}
func insertLast(val int, tail **node) {
	var temp *node = createNode((val))
	(*tail).next = temp
	temp.prev = *tail
	*tail = temp
}

func main(){
	var head,tail *node
	head = createNode(0)
	tail = head
	insertFirst(10,&head)
	insertFirst(80,&head)
	insertFirst(40,&head)
	insertLast(20,&tail)

	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
