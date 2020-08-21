package main

import (
	"fmt"
)

type PersonA struct {
	FirstName string
	LastName  string
}

//method using a pointer to Bobby
func (p *PersonA) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {

	fmt.Println("### Pointers ###")

	/*
		   Go has pointers, but no pointer arithmetic.
			Struct fields can be accessed through a struct pointer.
		   The indirection through the pointer is transparent
			(you can directly call fields and methods on a pointer).

		   Note that by default Go passes arguments by value (copying the arguments),
		   if you want to pass the arguments by reference, you need to pass pointers
			(or use a structure using reference values like slices (Section 4.2)
			and maps (Section 4.4).

		   To get the pointer of a value, use the & symbol in front of the value,
			to dereference a pointer, use the * symbol.

		   Methods are often defined on pointers and not values (although they can be defined on both),
		   so you will often store a pointer in a variable as in the example below:
	*/

	//declare Bobby, set it's params and print full name
	Bobby := new(PersonA)
	Bobby.FirstName = "Bobby"
	Bobby.LastName = "Peters"
	fmt.Println(Bobby.FullName())

}
