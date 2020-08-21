package main

import (
	"fmt"
)

func main() {

	//slices use make - make is a special built-in function
	// that is used to initialize slices, maps, and channels
	cities := make([]string, 2)
	cities[0] = "Houston"
	cities[1] = "Austin"

	//iterate over slice using index and value
	for index, value := range cities {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}

	//just show value "_" ignores the index returned from slice
	for _, value := range cities {
		fmt.Printf("%s\n", value)
	}

}
