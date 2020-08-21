package main

import "fmt"

//car with some attributes and ability to start
type Car struct {
	MaxSpeed int
	State    string
}

//notice the pointer on c. We will be changing values so need a pointer
func (c *Car) Startup() {
	c.State = "running"
}

//motorcycle with some attributes and ability to start
type Motorcycle struct {
	MaxSpeed int
	State    string
}

//notice the pointer on m. We will be changing values so need a pointer
func (m *Motorcycle) Startup() {
	m.State = "running"
}

//Interface that implements Startup() method
type Start interface {
	Startup()
}

//we passed the interface to this function and called startup
//on whatever object was passed
func StartVehicle(s Start) {
	s.Startup()
}

func main() {

	//created car and motorcycle
	car := Car{MaxSpeed: 150, State: "Stopped"}
	motorcycle := Motorcycle{MaxSpeed: 200, State: "Stopped"}

	//start vehicle. notice the pointer we pass - not calue
	//instead of this, we also could have:
	StartVehicle(&car)

	fmt.Println("Car is", car.State)               //running
	fmt.Println("Motorcycle is", motorcycle.State) //stopped

}
