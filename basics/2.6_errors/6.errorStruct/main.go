package main

import (
	"fmt"
)

type ErrorStruct struct {
	Name        string
	Description string
	ID          int
}

func (e *ErrorStruct) Error() string {
	return fmt.Sprintf("%d: %s Desc: %s   ", e.ID, e.Name, e.Description)
}

func SayName(number int) error {

	if number == 10 {

		errorStruct := &ErrorStruct{Name: "Display Error", Description: "Error saying the name", ID: 404}
		return errorStruct

	} else {
		fmt.Println(number)
		return nil
	}

}

func main() {

	//will be fine
	if e := SayName(111); e != nil {
		fmt.Println("got an error")
	}

	//will return an error
	if e := SayName(10); e != nil {
		fmt.Println(e)
	}

}
