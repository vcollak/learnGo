package main

import (
	"errors"
	"fmt"
)

func main() {

	//wrap the printNumber function in must. If the
	//printNumber returns an error, then must will panic
	//in this case it will NOT panic since err1 is nil and not error
	var err1 error
	err1 = nil
	fmt.Println("1: start")
	must(printNumber(err1))
	fmt.Println("1: done")

	//wrap the printNumber function in must. If the
	//printNumber returns an error, then must will panic
	//in this case it will panic since err is an Error
	var err error
	err = errors.New("this is an error")
	fmt.Println("2: start")
	must(printNumber(err))
	fmt.Println("2: done")

}

func printNumber(err error) error {
	return err
}

//throws a panic if there is an error
func must(err error) {
	if err != nil {
		panic(err)
	}
}
