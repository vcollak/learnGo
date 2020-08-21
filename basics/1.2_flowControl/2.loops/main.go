package main

import (
	"fmt"
)

func main() {

	//typical for loop
	fmt.Println("### A typical 'for' loop ###")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)

	}

	//like a while loop
	fmt.Println("### While loop ###")
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
		fmt.Println(sum1)
	}

}
