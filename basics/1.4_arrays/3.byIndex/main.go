package main

import "fmt"

func main() {

	//slice instead of the array
	fruitsB := []string{"raspberries", "strawberries", "blueberries"}
	fmt.Println(fruitsB)

	//print count
	fmt.Println(len(fruitsB))

	//just like in python we can use this notation to get just some elements
	fmt.Println(fruitsB[1:2])
	fmt.Println(fruitsB[1:])
	fmt.Println(fruitsB[:2])

}
