package main

import (
	"fmt"
)

type Person3 struct {
	FirstName, LastName, SSN string
	age                      int
}

var person3 map[string]Person3

func main() {

	fmt.Println("### MutatingMaps ###")

	//created a map
	person3 = make(map[string]Person3)

	//Bobby - assigned person literal
	person3["Bobby"] = Person3{

		"Bobbyimir",
		"Peters",
		"434-33-3333",
		39,
	}

	//rad - created empty Person then set properties
	r := Person3{}
	r.FirstName = "Johnny"
	r.LastName = "Peters"
	r.SSN = "333-54-7767"
	r.age = 7
	person3["rad"] = r

	//ash - created person, then added to map
	a := Person3{

		FirstName: "Austin",
		LastName:  "Peters",
		SSN:       "323-44-5454",
		age:       10,
	}
	person3["Austin"] = a

	fmt.Println(person3["rad"])

	//set Johnny to Austin and print
	person3["Austin"] = person3["rad"]
	fmt.Println(person3["Austin"])

	//delete Johnny
	delete(person3, "rad")

	//test if a key in a map is present
	var elem, ok = person3["Austin"]
	if ok {
		fmt.Println("Is present:", elem)
	} else {
		fmt.Println("Not there")
	}

}
