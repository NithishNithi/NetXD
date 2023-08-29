package main

import (
    "fmt"
)

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (ll *LinkedList) insert(data int) {
    newNode := &Node{data: data, next: nil}

    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

func (ll *LinkedList) display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    fmt.Println("Hello World!")
    a := LinkedList{}

    a.insert(10)
    a.insert(20)
    a.insert(30)

    a.display()
}
