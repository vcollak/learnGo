package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("### Errors3 ###")

	err := errors.New("New error")
	if err != nil {
		fmt.Print(err)
	}
}
