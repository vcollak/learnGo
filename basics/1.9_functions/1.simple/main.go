// Packages
package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

/*  Cannot redeclare like this. function name has to be unique
func add(x, y, z int) int {
	return x + y + z
}*/

//only declare param type once and apply to both params
func addA(x, y int) int {
	return x + y
}

/* function returns multiple parameters */
func location(city string) (string, string) {

	var region string
	var country string

	switch city {

	case "LA", "SF":
		region, country = "California", "USA"
	case "NYC":
		region, country = "Ney York", "USA"
	default:
		region, country = "Unknown", "Unknown"
	}

	return region, country

}

/* returns multiple. they are named now
I personally recommend against using named return
parameters because they often cause more confusion
than they save time or help clarify your code.
*/
func locationA(city string) (region, country string) {

	/* no need to redeclare
	var region string
	var country string
	*/

	switch city {

	case "LA", "SF":
		region, country = "California", "USA"
	case "NYC":
		region, country = "New York", "USA"
	default:
		region, country = "Unknown", "Unknown"
	}

	//return all params
	return

}

func main() {

	fmt.Println("### Functions ###")

	fmt.Println(add(42, 13))
	fmt.Println(addA(10, 13))
	///fmt.Println(add(1, 2, 3))

	//function returns multiple parameters
	region, country := location("LA")
	fmt.Printf("%s, %s", region, country)
	fmt.Println("")

	regionA, countryA := locationA("NYC")
	fmt.Printf("%s, %s", regionA, countryA)
	fmt.Println("")

}
