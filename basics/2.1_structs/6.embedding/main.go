package main

import "fmt"

//Organization holds the organizational data
type Organization struct {
	Name      string
	Employees int
	Locations []string
}

//Employee defines the employee info
//this embeds (composition) the Organization struct under the Organization property
type Employee struct {
	Name         string
	Age          int
	Organization Organization
}

//UpdateName is a method that allows us to update the Name
func (e *Employee) UpdateName(name string) {

	//only update if the name is not empty
	if len(name) > 0 {
		e.Name = name
		fmt.Printf("Name updated to %s\n", e.Name)
	} else {
		fmt.Printf("Empty name. Not updating. Name still %s\n", e.Name)
	}

}

func main() {

	company := Organization{
		Name:      "Ceremity",
		Employees: 44,
		Locations: []string{"Houston", "Kosice"}, //remember that have to specify type. cannot just do Locations: {"Houston", "Kosice"}
	}

	employee := Employee{
		Name:         "Bobby",
		Age:          43,
		Organization: company,
	}

	fmt.Printf("Employee Name: %s\n", employee.Name)
	fmt.Printf("Age: %d\n", employee.Age)
	fmt.Printf("Works for: %s\n", employee.Organization.Name)

	fmt.Printf("\n### Organizational Info ###\n")
	fmt.Printf("Name: %s", company.Name)
	fmt.Println()

	fmt.Printf("Employees: %d", company.Employees)
	fmt.Println()

	for _, loc := range company.Locations {
		fmt.Printf("Location: %s", loc)
		fmt.Println()
	}

	employee.UpdateName("")
	employee.UpdateName("Bobbyko")

}
