package main

import "fmt"

func enqueue(queue []int, num int) []int {
	queue = append(queue, num)
	return queue
}

func dequeue(queue []int) []int {
	return queue[1:]
}

func main() {
	var queue = make([]int, 0)
	queue = enqueue(queue, 10)
	fmt.Println(queue)
	queue = enqueue(queue, 20)
	fmt.Println(queue)
	queue = dequeue(queue)
	fmt.Println(queue)
}
