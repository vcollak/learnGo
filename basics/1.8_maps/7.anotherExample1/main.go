package main

import (
	"fmt"
)

func main() {

	//can define then initialize empty one
	var birthYear map[string]string
	birthYear = make(map[string]string)
	birthYear["Bobby"] = "1975"
	birthYear["Liz"] = "1975"
	fmt.Println(birthYear["Liz"])

	//can define and initialize at the same time
	birthMonth := map[string]string{
		"Liz":   "October",
		"Bobby": "May",
	}
	birthMonth["Liz"] = "October"
	fmt.Println(birthMonth["Liz"])

}
