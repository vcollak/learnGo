package main

import (
	"fmt"
)

//---- Interface ----
type Animal interface {
	Speak() string
}

//---- Dog ----
type Dog struct {
}

func (d Dog) Speak() string {
	return "Dog says Woof!"
}

//---- Cat ----
type Cat struct {
}

func (c Cat) Speak() string {
	return "Cat says Meow!"
}

func main() {

	fmt.Println("### Interfaces3 ###")

	//create animals slice
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}
