package main

import (
	"fmt"
)

//Person is a struct that defines the person
type Person struct {
	FirstName string
	LastName  string
}

//FullName says the full name of the person
func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

//NewPerson creates and returns a new person object
//created a function that creates the object so we can profive some logic
//on creation
func NewPerson(fname, lname string) *Person {

	if len(fname) != 0 || len(lname) != 0 {
		p := new(Person)
		p.FirstName = fname
		p.LastName = lname
		return p
	}

	return nil

}

func main() {

	//creates a new person
	p := NewPerson("Bobbyimir", "Peters")

	//bad name
	//p := NewPerson("", "")

	if p != nil {
		//prints person's name
		fmt.Printf("My name is %s", p.FullName())
	} else {
		fmt.Println("Unable to get name")
	}

}
