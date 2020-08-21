package main

import (
	"fmt"
)

//person struct
type person struct {
	name string
	age  int
}

//return the older person
func Older(p1, p2 person) person {

	if p1.age > p2.age {
		return p1
	} else {
		return p2
	}
}

func main() {

	fmt.Println("### Structs 2 ###")

	p1 := person{name: "Bobby", age: 40}
	p2 := person{name: "Liz", age: 41}

	var olderPerson person

	olderPerson = Older(p1, p2)

	fmt.Printf("The older person is %s with age %d\n", olderPerson.name, olderPerson.age)

}
