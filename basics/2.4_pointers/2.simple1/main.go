package main

import "fmt"

func main() {

	name := "Bobby"
	fName := &name

	//print name - shows Bobby
	fmt.Println(name)

	//print the pointer - shows the memory address
	fmt.Println(fName)

	//prints the value behind the pointer
	fmt.Println(*fName)

	//change the value of the pointer
	*fName = "Liz"

	//print the name and fName
	fmt.Println(name)   //printing by value
	fmt.Println(*fName) //printing by value of the pointer

}
