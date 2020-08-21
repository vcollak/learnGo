package main

import (
	"fmt"
)

func main() {

	capitals := map[string]string{
		"texas":      "Austin",
		"california": "Sacramento",
		"nebraska":   "Lincoln",
	}

	fmt.Println("-- loop over a structure using key --")
	for key, value := range capitals {
		fmt.Printf("%s  %s\n", key, value)
	}

	fmt.Println("-- loop over a structure ignoring keys --")
	for _, valueA := range capitals {
		fmt.Println(valueA)
	}

	fmt.Println("-- loop over a structure ignoring values --")
	for keyA := range capitals {
		fmt.Println(keyA)
	}

}
