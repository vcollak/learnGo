package main

import (
	"fmt"
)

func main() {

	//set array and loop over it
	c := [...]int{1, 2, 3, 4, 5}
	for i := 0; i <= len(c)-1; i++ {
		fmt.Println(i)
	}

}
