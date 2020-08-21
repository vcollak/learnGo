package main

import "fmt"

type anyValue interface{}

//function takes an empty interface and thus takes any value
func takeAnyValue(val anyValue) {
	fmt.Println("Took value:", val)
}

func main() {

	fmt.Println("### Interfaces10 ###")

	takeAnyValue("1")
	takeAnyValue(1)

}
