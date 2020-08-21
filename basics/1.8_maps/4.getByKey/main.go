package main

import (
	"fmt"
)

func main() {

	fmt.Println("### Maps 1 ###")

	person := map[string]string{

		"name": "Bobbyimir Peters",
		"ssn":  "332-33-2323",
	}

	fmt.Printf("%#v\n", person)
	fmt.Println(person["name"])

	/////////////////////////////
	ages := map[string]int{

		"abbie":  7,
		"austin": 10,
		"peter":  13,
		"liz":    39,
		"bobby":  39,
	}

	fmt.Println("Abbie is ", ages["abbie"])

}
