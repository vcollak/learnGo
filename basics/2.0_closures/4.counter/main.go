package main

import "fmt"

//intSequence is a closure. When invoked, returns an
//anonymous function that returns int
func intSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//put intSequence in counter variables
	nextCounter1 := intSequence()
	nextCounter2 := intSequence()

	//the key property of this is that the
	//function will remember it's scope even ourside
	//of the scope
	fmt.Println(nextCounter1()) //returns 1
	fmt.Println(nextCounter1()) //returns 2
	fmt.Println(nextCounter2()) //returns 1
	fmt.Println(nextCounter1()) //returns 3

}
