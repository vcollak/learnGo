package main

import "fmt"

func main() {

	name := map[string]string{
		"first": "Bobby",
		"last":  "Peters",
	}

	//if the key is found the map
	//returns tuple with value and true
	value, ok := name["first"]
	if ok {
		fmt.Printf("The first name is %s\n", value)
	} else {
		fmt.Println("Unable to fund the key")
	}

	//if the key is NOT found the map
	//returns tuple with value and false
	valueDude, okDude := name["dude"]
	if okDude {
		fmt.Printf("The dude name is %s\n", valueDude)
	} else {
		fmt.Println("Unable to fund the key 'dude'")
	}

}
