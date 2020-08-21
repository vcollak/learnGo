package main

import "fmt"

//Person struct
type Person struct {
	first  string
	last   string
	gender string
	age    int
}

func main() {

	personA := Person{
		first:  "Bobby",
		last:   "Peters",
		gender: "M",
		age:    35,
	}
	fmt.Println(personA)

	personB := Person{"Liz", "Peters", "F", 35}
	fmt.Println(personB)

}
