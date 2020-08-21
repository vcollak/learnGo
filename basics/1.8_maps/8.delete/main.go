package main

import "fmt"

func main() {

	phoneNumbers := make(map[string]string)
	phoneNumbers["Bobby"] = "555-555-5555"
	phoneNumbers["Liz"] = "555-555-5555"
	phoneNumbers["random"] = "111-111-1111"

	fmt.Println(phoneNumbers)

	//delete from a map
	delete(phoneNumbers, "random")

	fmt.Println(phoneNumbers)

	phoneNumbers["allOnes"] = "111-111-1111"

	fmt.Println(phoneNumbers)

}
