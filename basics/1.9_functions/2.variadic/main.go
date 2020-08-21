package main

import (
	"fmt"
)

//sum is a function that can take any number of int arguments
//and it will return a sum
func sum(numbers ...int) int {

	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}

	return sum

}

func main() {

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))

}
