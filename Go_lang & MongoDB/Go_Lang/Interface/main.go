package main

import (
	"fmt"
	"math"
)


type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	Radius float64
}
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}


func main(){
	var s Shape = Circle{10.4}
	fmt.Println(s.Area())

}