package main

import (

	//Packages are pathed to GOROOT or GOPATH.
	//If these were in GOPATH/src then we just reference "appName/de" or "appName/en"
	"LearnGo/basics/2.8_packages/5.packagesWithSubpackages/greet/de"
	"LearnGo/basics/2.8_packages/5.packagesWithSubpackages/greet/en"

	"fmt"
)

var integers [10]int

//this init is called first since it appears first
//use init to initialize types. in this case we're
//initializing the intergers array
func init() {
	fmt.Println("Package Main (main.go): Init called first")

	//init is used here to initialize the integers
	for i := 0; i < 10; i++ {
		integers[i] = i
	}

}

//called second
func init() {
	fmt.Println("Package Main (main.go): Init called second")
}

//called third
func init() {
	fmt.Println("Package Main (main.go): Init called third")
}

func main() {

	//prints from en package
	fmt.Println(en.Morning)

	//prints from de package
	fmt.Println(de.Morning)

	//prints version from the package main version var
	//lowercase since this is in the main package
	fmt.Println("Version:" + version)

	//prints the integers we initialized in init()
	fmt.Println(integers)

}
