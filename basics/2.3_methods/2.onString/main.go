package main

import (
	"fmt"
	"strings"
)

//alias to string
type myString string

func (s myString) Uppercase() string {
	return strings.ToUpper(string(s))
}

func main() {

	fmt.Println("### Methods1 ###")
	fmt.Println(myString("test").Uppercase())

}
