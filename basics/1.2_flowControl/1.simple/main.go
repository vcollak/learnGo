package main

import (
	"fmt"
)

func main() {

	//define a var
	answer := 42

	//if
	fmt.Println("### Simple if ###")
	if answer != 42 {
		fmt.Println("number not found")
	} else {
		fmt.Println("number found")
	}

	//define the var directly in the if statement
	fmt.Println("### if - assign and evaluate ###")
	if newAnswer := 42; newAnswer != 42 {
		fmt.Println("number not found")
	} else {
		fmt.Println("number found")
	}

	//will not compile. the newAnswer is not visible outside 
	//of IF scope
	//fmt.Println(newAnswer)

	//else if
	fmt.Println("### else if ###")
	a := 10
	if a > 10 {
		fmt.Println("a > 10")
	} else if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a < 10")
	}

}
