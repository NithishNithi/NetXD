package main

import (
	"fmt"
	"time"
)

type Item struct {
	Id            int
	Name          string
	PackageEffort time.Duration
}

func PrepareItems(done <-chan bool) <-chan Item {
	items := make(chan Item)
	itemsToShip := []Item{
		{0, "Shirt", 1 * time.Second},
		{1, "legos", 5 * time.Second},
		{2, "t-shirt", 1 * time.Second},
		{3, "hat", 5 * time.Second},
		{4, "shorts", 3 * time.Second},
		{5, "jacket", 2 * time.Second},
		{6, "watch", 1 * time.Second},
		{7, "bracelet", 5 * time.Second},
		{8, "tie", 1 * time.Second},
		{9, "bananas", 3 * time.Second},
	}
	go func() {
		for _, item := range itemsToShip {
			select {
			case <-done:
				return
			case items <- item:
			}
		}
		close(items)
	}()
	return items
}

func PackItems(done <-chan bool, items <-chan Item) <-chan int {
	packages := make(chan int)
	go func() {
		for item := range items {
			select {
			case <-done:
				return
			case packages <- item.Id:
				time.Sleep(item.PackageEffort)
			}
		}
		close(packages)
	}()
	return packages
}

func main() {
	done := make(chan bool)
	defer close(done)
	start := time.Now()
	packages := PackItems(done, PrepareItems(done))
	numPackages := 0
	for p := range packages {
		numPackages++
		fmt.Printf("Shipping Package no. %d \n", p)
	}

	fmt.Printf("Took %fs to ship %d packages \n", time.Since(start).Seconds(), numPackages)
}
 