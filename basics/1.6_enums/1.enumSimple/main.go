package main

import (
	"fmt"
)

//Weekday is an enum that defines the weekdays
type Weekday int

//constants int constant for each day of the week of type Weekday
const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

func main() {

	//prints the value of Sunday (0)
	fmt.Println(Sunday)

}
