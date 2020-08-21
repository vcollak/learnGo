package main

import (
	"fmt"
)

type User2 struct {
	FirstName, LastName string
}

func (user *User2) Greeting() string {
	return fmt.Sprintf("Dear %s %s", user.FirstName, user.LastName)
}

func main() {

	fmt.Println("### Methods ###")

	//can create User this way
	//user := User{"Matt", "Duder"}

	//or this way
	/*
		user := User{
			FirstName: "Matt",
			LastName:  "Duder",
		}
	*/

	//or this way
	user := new(User2)
	user.FirstName = "Bobby"
	user.LastName = "Peters"

	//print greeting
	fmt.Println(user.Greeting())

}
