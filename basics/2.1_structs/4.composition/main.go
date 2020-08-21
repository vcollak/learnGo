package main

import (
	"fmt"
)

//person struct
type Human struct {
	name  string
	age   int
	phone string
}

//basically inherits from Human (using composition)
type Student struct {
	Human
	major string
	phone string
}

func main() {

	fmt.Println("### Structs 3 ###")

	//declare Bobby with names
	Bobby := Student{Human: Human{name: "Bobby", age: 40, phone: "555-555-5555"}, major: "CS", phone: "555-555-5555"}

	//or just values in right order
	//Bobby := Student{Human{"Bobby", 40}, "CS"}
	fmt.Printf("%s is %d years old and majorss in %s\n", Bobby.name, Bobby.age, Bobby.major)
	fmt.Printf("Bobby phone is %s\n", Bobby.phone)
	fmt.Printf("Human phone is %s\n", Bobby.Human.phone)

}
