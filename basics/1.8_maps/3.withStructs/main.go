package main

import (
	"fmt"
)

func main() {

	family := map[string]int{
		"Liz":    43,
		"Bobby":  43,
		"Grace":  18,
		"Austin": 14,
		"Johnny": 11,
	}

	fmt.Println(family["Liz"])

	type Locations struct {
		lat  float64
		long float64
	}

	locations := map[string]Locations{
		"houston": Locations{40.222, -74.000},
		"katy":    Locations{40.111, -74.111},
	}

	fmt.Println(locations["houston"])
	fmt.Println(locations["houston"].lat)

}


