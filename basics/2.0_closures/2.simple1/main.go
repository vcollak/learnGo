package main

import "fmt"

//define function called anonFunction that returns
//an anonynmous function that takes a string
func anonFunction() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}

func main() {

	//place function in a variable
	printFunc := anonFunction()

	//call the function
	printFunc("hello return from anon function")

}
