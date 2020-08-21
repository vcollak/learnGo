package main

import (
	"fmt"
)

//changeNumber takes a pointer to a number and changes the value
func changeNumberPointer(n *int) {
	*n = 10
}

//changeNumber takes an n and chnages it's value
func changeNumber(n int) {
	n = 10
}

func main() {

	//crate a number
	number := 1

	//by value - will not change the number because it's passed by value
	changeNumber(number)
	fmt.Println(number)

	//by pointer - changes the value as it's passed by pointer
	n := &number
	changeNumberPointer(n)
	fmt.Println(number)

}
