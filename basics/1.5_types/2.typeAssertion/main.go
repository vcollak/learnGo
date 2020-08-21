package main

import "fmt"

//Stringer is an interface
type Stringer interface {
	String() string
}

//fakeString isa struct
type fakeString struct {
	content string
}

// function used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}

//prints the types
func printString(value interface{}) {

	switch value.(type) {
	case string:
		fmt.Println("string")
	case Stringer:
		fmt.Println("stringer interface")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("not sure what this is")
	}

}

func main() {

	fmt.Println("### Type Assertion ###")

	var myInt int = 10
	var myFloat float64 = 4.13
	var myStr string = "This is a string"
	var myBoolean bool = true
	var myStruct Stringer = &fakeString{"Ahoj Svet!"}

	//uses the printString function to print
	//the types
	printString(myInt)
	printString(myFloat)
	printString(myStr)
	printString(myBoolean)
	printString(myStruct)

}
