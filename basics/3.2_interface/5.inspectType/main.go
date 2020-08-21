package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element
type Person7 struct {
	name string
	age  int
}

func (p Person7) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func main() {

	fmt.Println("### Interfaces11 ###")

	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person7{"Dennis", 70}

	//this essentially allows us to inspect the interface type
	for index, element := range list {

		if value, ok := element.(int); ok {

			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)

		} else if value, ok := element.(string); ok {

			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)

		} else if value, ok := element.(Person7); ok {

			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)

		} else {

			fmt.Printf("list[%d] is of a different type", index)
		}
	}
}
