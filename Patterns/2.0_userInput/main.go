package main

import (
	"fmt"
)

func main() {

	fmt.Println("### UserInput ###")

	fmt.Print("Enter a number: ")

	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Printf("Input %g times 2 is: %g", input, output)
	fmt.Print("\n")

}
