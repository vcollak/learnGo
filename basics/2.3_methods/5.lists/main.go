package main

import (
	"fmt"
)

//struct that holds a person
type Person4 struct {
	FirstName string
	LastName  string
}

//slice that holds list of Person structs
type Persons struct {
	Persons []Person4
}

//method that allows to add a person
func (p *Persons) Add(person Person4) {

	//add a person to persons
	p.Persons = append(p.Persons, person)
}

func main() {

	fmt.Println("### Methods4 ###")

	//new persons slice
	persons := new(Persons)

	//create 2 Person structs and add them to Persons
	persons.Add(Person4{"Bobby", "Peters"})
	persons.Add(Person4{"Liz", "Peters"})

	//loop through persons and print the structs
	for _, p := range persons.Persons {
		fmt.Println(p)
	}

}
