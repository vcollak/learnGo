package main

import "fmt"

////////////////////
// Human
////////////////////
type Human struct {
	Name string
}

func (h Human) SayName() {
	fmt.Println(h.Name)
}

////////////////////
// Car
////////////////////
type Car struct {
	Name string
}

func (c Car) SayName() {
	fmt.Println(c.Name)
}

////////////////////
// Interface
////////////////////
type Names interface {
	SayName()
}

func main() {

	h := Human{Name: "Bobby"}
	c := Car{Name: "Mercedes"}

	//variable for interface
	var n Names

	//interface to become Human
	n = h
	n.SayName() //call the interface method

	//interface to become Car
	n = c
	n.SayName() //call the interface method

}
