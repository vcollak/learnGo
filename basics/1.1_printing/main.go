package main

import "fmt"

func main() {

	fmt.Println("### Printing ###")

	cylonModel := 6
	fmt.Println(cylonModel)

	/*
		fmt.Println prints the pass variables values and appends a newline.
		fmt.Printf is used when you want to print one or multiple values using a defined format specifier.
	*/

	name := "Bobbyimir"
	aka := "Bobby"
	fmt.Printf("%s is also known as %s", name, aka)
	fmt.Println("")

	myString := "Bobby"

	//just print
	fmt.Printf("No New Line: Print this stuff: %s", myString)

	//print, but return what was printed
	returnVal := fmt.Sprintf("Print this stuff: %s", myString)
	fmt.Println(returnVal)

	num := 1
	fmt.Printf("Number: %d", num)

}
