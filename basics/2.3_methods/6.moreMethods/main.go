package main

import (
	"fmt"
)

type Person6 struct {
	Name    string
	Age     int
	Address Address
	Human6
}

type Human6 struct {
	SpeciesName string
}

type Address struct {
	Address string
	City    string
	State   string
	Country string
}

func (p *Person6) GetName() string {
	return p.Name
}

func (a *Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Address, a.City, a.State, a.Country)
}

func (h *Human6) GetSpeciesName() string {

	return h.SpeciesName

}

func main() {

	fmt.Println("### Methods5 ###")

	person := new(Person6)
	person.Name = "Bobby"
	person.Address.Address = "20515 Cottage"
	person.Address.City = "Richond"
	person.Address.State = "TX"
	person.Address.Country = "USA"
	person.SpeciesName = "Human"

	fmt.Println("My name is ", person.Name)
	fmt.Println("I am ", person.GetSpeciesName())
	fmt.Println("My Address is", person.Address.FullAddress())

}
