package main

import (
	"fmt"
)

func main() {

	fmt.Println("### Slices 3 ###")

	//this is a slice literal
	//myStringB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//we can also make slices this way
	//make a slice with 3 elements
	cities := make([]string, 3)
	cities[0] = "houston"
	cities[1] = "austin"
	cities[2] = "san francisco"
	fmt.Printf("%q\n", cities)

	//appendign to a slice
	//empty slice
	countries := []string{}
	countries = append(countries, "United States")
	countries = append(countries, "Slovak Republic")

	//appending several strings to a slice
	countries = append(countries, "France", "England")
	fmt.Printf("%q\n", countries)

	//addinf slices together
	names := []string{"Bobby", "Rad"}
	otherNames := []string{"Liz", "Grace"}

	//we just added a slice to a slice. NOTE the three "..."
	//get an error if we skip the three "..."
	names = append(cities, otherNames...)
	fmt.Printf("%q\n", names)

	//check the length of a slice
	fmt.Println(len(names))

	//empty slices
	var myEmptySlice []string
	//myEmptySlice = append(myEmptySlice, "test")

	if myEmptySlice == nil {
		fmt.Println("Empty Slice")
	} else {
		fmt.Println("NOT Empty Slice")
	}

}
