package main

import (
	"fmt"
)

//Weekday is an enum that defines the weekd days
type Weekday int

//defines the constants for each day
const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

//method for Weekday that given a Weekday will return the weekday name
func (day Weekday) String() string {

	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	if day < Sunday || day > Saturday {
		return "Unknown"
	}

	return names[day]

}

//Weekend returns true if a day is a weekend
func (day Weekday) Weekend() bool {

	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}

}

func main() {

	//prints the day value
	fmt.Println(Sunday)

	//prints the day name
	fmt.Printf("Which day is it? %s\n", Sunday)

	//prints true if Saturday is a weekend
	fmt.Printf("Is Saturday a weekend? %t\n", Saturday.Weekend())

}
