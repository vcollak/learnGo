package main

import (
	"fmt"
)

type Person2 struct {
	FirstName, LastName, SSN string
	age                      int
}

var person2 map[string]Person2

func main() {

	fmt.Println("### Maps 3 ###")

	//created a map
	person2 = make(map[string]Person2)

	//Bobby - assigned person literal
	person2["Bobby"] = Person2{

		"Bobbyimir",
		"Peters",
		"434-33-3333",
		39,
	}

	//rad - created empty Person then set properties
	r := Person2{}
	r.FirstName = "Johnny"
	r.LastName = "Peters"
	r.SSN = "333-54-7767"
	r.age = 7
	person2["rad"] = r

	//ash - created person, then added to map
	a := Person2{

		FirstName: "Austin",
		LastName:  "Peters",
		SSN:       "323-44-5454",
		age:       10,
	}
	person2["Austin"] = a

	fmt.Println(person2)

}
