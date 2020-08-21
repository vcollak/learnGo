package main

import "fmt"

//part of the main package
//this will be called first if when starting the program a.go goes firs like: go run a.go version.go main.go
func init() {
	fmt.Println("Package Main (a.go): Init called")
}
