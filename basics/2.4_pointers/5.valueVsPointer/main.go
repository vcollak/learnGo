package main

import "fmt"

//taking the pointer (reference of number)
func SquarePtr(number *int) {

	//make sure we show these are pointers of number that we are
	//multiplying
	*number = *number * *number
}

func Square(number int) {
	number = number * number
}

func main() {

	number := 2

	//sending a pointer to the number (passing by reference
	SquarePtr(&number)
	fmt.Println(number)

	//now we just call by value so the number 4 does not
	//actually get multipled in the main only in Square
	Square(number)
	fmt.Println(number)

}
