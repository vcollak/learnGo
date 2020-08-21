package main

import "fmt"

//Concatenate takes a list of strings and concatenates them
func Concatenate(strings ...string) string {

	finalStr := ""

	for _, str := range strings {
		finalStr = finalStr + str
	}
	return finalStr

}

func main() {

	fmt.Println(Concatenate("Life", " ", "is", " ", "good"))

}
