package main

import "fmt"

func main() {

	//anonymous function - define and then use it by
	//passing ("Hello World)"
	func(msg string) {
		fmt.Println(msg)
	}("Hello world")

}
