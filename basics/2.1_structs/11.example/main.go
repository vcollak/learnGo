package main

import "fmt"

//Organization holds the organizational data
type Organization struct {
	Name      string
	Employees int
	Locations []string
}

func main() {

	company := Organization{
		Name:      "Ceremity",
		Employees: 44,
		Locations: []string{"Houston", "Kosice"}, //remember that have to specify type. cannot just do Locations: {"Houston", "Kosice"}
	}

	fmt.Printf("Name: %s", company.Name)
	fmt.Println()

	fmt.Printf("Employees: %d", company.Employees)
	fmt.Println()

	for _, loc := range company.Locations {
		fmt.Printf("Location: %s", loc)
		fmt.Println()
	}

}
