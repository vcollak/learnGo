package main

import (
	"LearnGo/basics/2.8_packages/3.moduleNumbers/numbers"
	"fmt"
)

func main() {

	//adds bunch of numbers
	fmt.Println(numbers.Sum(1, 2, 3, 4, 5))

	//subtracts 2 numbers
	fmt.Println(numbers.Subtract(10, 5))

}
