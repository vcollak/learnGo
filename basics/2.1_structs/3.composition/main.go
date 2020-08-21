package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	DOB       time.Time
}

type Employee struct {

	//this would have created an object under Employee called person
	//it's embedding
	//Person Person

	//this, without a type is composition like inheritance
	Person
	EmployeeNumber int
}

func (p *Person) FullName() string {
	return p.LastName + ", " + p.FirstName
}

func main() {

	fmt.Println("### Structs 1 ###")

	myPerson := Person{
		FirstName: "Bobby",
		LastName:  "Peters",
		DOB:       time.Now(),
	}

	fmt.Println(myPerson.FirstName)
	fmt.Println(myPerson.DOB)
	fmt.Println(myPerson.FullName())

	//------------------------
	myEmployee := Employee{

		Person:         myPerson,
		EmployeeNumber: 1,
	}

	fmt.Println(myEmployee.FirstName)
	fmt.Println(myEmployee.EmployeeNumber)

}
