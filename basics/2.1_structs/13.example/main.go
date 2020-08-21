package main

import (
	"fmt"
	"strconv"
)

//Person struct
type Person struct {
	first  string
	last   string
	gender string
	age    int
}

//not a pointer so this will be copy of Person
func (p Person) greet() string {
	return "My name is " + p.first + " " + p.last + " My age is " + strconv.Itoa(p.age)
}

//hasBirthday increses the age by 1
func (p *Person) hasBirthday() {
	p.age++
}

//gotMarried
func (p *Person) gotMarried(marriedName string) {

	//chaneg the maiden name if female
	if p.gender == "F" {
		p.last = marriedName
	}

}

func main() {

	man := Person{
		first:  "Bobby",
		last:   "Peters",
		gender: "M",
		age:    35,
	}

	woman := Person{"Liz", "Longoria", "F", 35}

	fmt.Println(man)
	fmt.Println(woman)

	man.hasBirthday()
	woman.gotMarried("Peters")

	fmt.Println(woman.greet())

}
