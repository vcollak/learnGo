package main

import (
	"fmt"
)

func main() {

	fmt.Println("### Slices ###")

	//slices using literals
	// [2 3 5 7 11 13]
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)

	// [3 5 7]
	fmt.Println(mySlice[1:4])

	// missing low index implies 0
	// [2 3 5]
	fmt.Println(mySlice[:3])

	// missing high index implies len(s)
	// [11 13]
	fmt.Println(mySlice[4:])

	//not a slice. This is actually an array. The [...] insures
	//the compiler counts the number of elements
	myName := [...]string{"Bobbyimir", "Bobbyimirovic", "Peters"}
	fmt.Println(myName[0])
	fmt.Println(myName[1:])
	fmt.Println(myName[:2])

	//make slices - slice of 3 strings
	cities := make([]string, 3)
	cities[0] = "Houston"
	cities[1] = "Austin"
	cities[2] = "San Antonio"

	//this will fail - as index out of range
	//cities[4] = "Corpus Christi"

	//we can add the 4th element by appending
	cities = append(cities, "Corpus Christi")
	fmt.Printf("%q", cities)

	//appending to s a slice
	citiesA := []string{}
	citiesA = append(citiesA, "LA")
	citiesA = append(cities, "SF, NY")
	fmt.Printf("%q", citiesA)

	citiesB := []string{"San Diego", "Mountain View"}
	otherCities := []string{"Santa Monica", "Venice"}
	citiesB = append(citiesB, otherCities...)
	fmt.Printf("%q", citiesB)
	fmt.Println(len(citiesB))

	var z []int
	fmt.Println(z, len(z), cap(z))
	// [] 0 0
	if z == nil {
		fmt.Println("nil!")
	}
	// nil!

	//convert myA array to myS slice
	myA := [4]string{"Paris", "London", "Prague", "Bratislava"}
	myS := myA[:]
	fmt.Println(myS)

}
