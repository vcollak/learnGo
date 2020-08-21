package main

import "fmt"

type Bootcamp1 struct {
	Lat float64
	Lon float64
}

func main() {

	fmt.Println("### Initializing ###")

	//initializes and return a pointer
	x := new(Bootcamp1)

	//same as above through a reference
	y := &Bootcamp1{}

	//initialize and set the properties
	z := Bootcamp1{Lat: 122.2, Lon: 1212.22}

	//these are the same thing
	fmt.Println(*x == *y)

	fmt.Println(z.Lat)

}
