package main

import "fmt"

//Person struct
type Person struct {
	firstName string
	lastName  string
}

/*Employee struct. This is embedding the person object
in the employee struct. if we want to call person inside of employee we'll
have to say Person.Employee
*/
type Employee struct {
	Person Person
	ID     uint
}

/*EmployeeA struct. This is composition. It's basically like inheritance
where the EmployeeA automatically inherits the atributes from Person.Employee
For instance, we can call EmployeeA.firstName
*/
type EmployeeA struct {
	Person
	ID uint
}

func main() {

	//create the person object
	person := Person{
		firstName: "Bobbyimir",
		lastName:  "Peters",
	}

	//create the employee object
	employee := Employee{
		Person: person,
		ID:     123343243124234,
	}

	//create the employeeA object
	employeeA := EmployeeA{
		Person: person,
		ID:     123343243124234,
	}

	fmt.Println(person.firstName)
	fmt.Println(employee.Person.firstName, employee.ID) //embedding. have to call employee.Person.firstName
	fmt.Println(employeeA.firstName, employeeA.ID)      //composition can call employeeA.firstName

}
