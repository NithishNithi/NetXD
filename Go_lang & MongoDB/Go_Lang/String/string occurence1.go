package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string 
	fmt.Scanln(&str)
	fmt.Println("Hi")
	var str1 string
	fmt.Scanln(&str1)

	fmt.Println(strings.Count(str,str1))
}