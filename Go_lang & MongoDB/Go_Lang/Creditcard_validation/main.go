package main

import (
	"creditcard/card"
	"fmt"
)

func main() {

	var cardno string
	fmt.Scanln(&cardno)

	ans := card.Validation(cardno)
	fmt.Println(ans)
}
