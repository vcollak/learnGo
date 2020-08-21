package main

import (
	"errors"
	"fmt"
)

func SayName(number int) error {

	if number == 10 {
		//generate a new error and return it
		return errors.New("Got name error")
	} else {
		fmt.Println(number)
		return nil
	}

}

func main() {

	//will be fine
	if SayName(111) != nil {
		fmt.Println("got an error")
	}

	//will return an error
	if SayName(10) != nil {
		fmt.Println("got an error")
	}

}
