package main

import "fmt"

func main() {

	//simple var declaration with type
	var yearBorn int
	yearBorn = 2000
	fmt.Println(yearBorn)

	//several all at once (without repeating var)
	var (
		name     string
		age      int
		location string
	)
	name = "Bobby Peters"
	age = 35
	location = "Earth"
	fmt.Println(name, age, location)

	//same types together
	var (
		nameA, locationA string
		ageA             int
	)
	nameA = "Bobby Peters"
	ageA = 34
	locationA = "Milky Way"
	fmt.Println(nameA, ageA, locationA)

	//declare and initialize
	var (
		nameC string = "Bobby Peters"
		ageC  int    = 21
	)
	fmt.Println(nameC, ageC)

	//can omit type if initializing
	//type gets inferred
	var (
		nameD = "Bobbyimir Peters"
	)
	fmt.Println(nameD)

	//initialize all in one line
	var (
		nameE, locationE, ageE = "Bobby Peters", "Houston", 39
	)
	fmt.Println(nameE, locationE, ageE)

	//no need for var if using a shorthand
	//set and initialize at the same time
	nameF := "Bobbyimir Peters"
	ageF := 39
	fmt.Println(nameF, ageF)

	//function as a variable
	action := func() {
		fmt.Println("test function as a var")
	}
	//execute the function
	action()

	//print the var the type
	someStringVar := "this is a string"
	fmt.Printf("%T\n", someStringVar)

	//use in fmt print statement
	fName, lName := "Bobby", "Peters"
	fmt.Printf("%s %s", fName, lName)

}
