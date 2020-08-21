package main

import (
	"fmt"
)

func main() {

	fmt.Println("### Loop over array - show value ###")
	myArray := [...]string{"hello", "world"}
	for _, element := range myArray {
		fmt.Println(element)
	}

	fmt.Println("### Loop over array - show index and value ###")
	myArrayA := [...]string{"hello", "world"}
	for indexA, elementA := range myArrayA {
		fmt.Printf("%d %s\n", indexA, elementA)
	}

}
