package main

import "fmt"

//calc returns names values. this allows us just to add
//return without any variables
func calc(a, b int) (sum int, subtraction int) {

	sum = a + b
	subtraction = a - b

	return

}

func main() {

	//calc returns a tuple of sum and subtraction
	sum, subtraction := calc(10, 5)
	fmt.Printf("10 + 5 = %d | 10 - 5 = %d", sum, subtraction)

}
