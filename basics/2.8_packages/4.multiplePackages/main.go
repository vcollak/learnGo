package main

import (
	"LearnGo/basics/2.8_packages/4.multiplePackages/cat"
	"LearnGo/basics/2.8_packages/4.multiplePackages/dog"
	"fmt"
)

func main() {

	fmt.Println(cat.GetName())
	fmt.Println(dog.GetName())

}
