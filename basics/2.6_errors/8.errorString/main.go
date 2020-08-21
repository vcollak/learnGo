package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	//ErrNotFound is a constant of the modelError type
	ErrNotFound modelError = "models: resource not found"
)

//model error type of a type string
type modelError string

//method that implements the error interface methos Error
func (e modelError) Error() string {
	return string(e)
}

//if we wanted to manipulate the way
//the error is shown we can create a method for that
//this one just removes the word "models:" and capitalizes
//the first word
func (e modelError) Public() error {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return errors.New(strings.Join(split, " "))
}

//return just the error
func doSomething() error {
	return ErrNotFound
}

//return the erro, but with text altered
func doSomethingPublic() error {
	return ErrNotFound.Public()
}

func main() {

	//show the error. calls the Error() method, which
	//returns the error string
	if err := doSomething(); err != nil {
		fmt.Println(err)
	}

	//Public error we can show to teh user.
	if err := doSomethingPublic(); err != nil {
		fmt.Println(err)
	}
}
