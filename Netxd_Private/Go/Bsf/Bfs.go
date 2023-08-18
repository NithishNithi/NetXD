package main

import "fmt"

type node struct {
	left *node
	data int
	ring *node
}

func creatnode(data int) *node {
	newnode := new(node)
	newnode.data = data
	newnode.left = nil
	newnode.ring = nil
	return newnode
}

func insertnode(data int, root *node) *node {
	if root == nil {
		return creatnode(data)
	}
	if data < root.data {
		root.left = insertnode(data, root.left)
	} else if data > root.data {
		root.ring = insertnode(data, root.ring)
	}
	return root
}

func display(head *node) {
	if head == nil {
		return
	}
	fmt.Printf("%d ", head.data)
	display(head.left)
	display(head.ring)
}

func search(data int, head *node) {
	if head == nil {
		return
	}
	if data == head.data {
		fmt.Printf("Found %d\n", head.data)
	}
	if data < head.data {
		search(data, head.left)
	}
	if data > head.data {
		search(data, head.ring)
	}
}

func main() {

	root := new(node)
	head := root
	a := []int{
		10, 20, 30, 40, 50, 60, 70, 80, 90, 100,
	}

	for i := 0; i < len(a); i++ {
		head = insertnode(a[i], head)
	}

	display(head)

}
