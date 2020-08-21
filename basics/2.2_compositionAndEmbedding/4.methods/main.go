package main

import (
	"fmt"
)

//square struct
type Square struct {
	x int
}

//struct method
func (s Square) Area() int {
	return s.x * s.x
}

func main() {

	fmt.Println("### Composition 3 ###")

	//define square
	square := Square{x: 10}

	//print its area
	fmt.Printf("Area of the square is %d\n", square.Area())

}
