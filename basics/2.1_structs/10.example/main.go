package main

import "fmt"

type House struct {
	Manufacturer  string
	YearBuilt     string
	NumberOfRooms int
}

func main() {

	//create a new house and print info
	house := House{Manufacturer: "Perry", YearBuilt: "2003", NumberOfRooms: 4}
	fmt.Printf("Manufacturer: %s", house.Manufacturer)
	fmt.Printf("\nYear built: %s", house.YearBuilt)
	fmt.Printf("\nNumber of rooms: %d", house.NumberOfRooms)

	fmt.Print("\n\n")

	//create a another house and print info
	houseL := House{Manufacturer: "Lenar", YearBuilt: "2010", NumberOfRooms: 6}
	fmt.Printf("Manufacturer: %s", houseL.Manufacturer)
	fmt.Printf("\nYear built: %s", houseL.YearBuilt)
	fmt.Printf("\nNumber of rooms: %d", houseL.NumberOfRooms)

}
