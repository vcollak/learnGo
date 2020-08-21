package main

import (
	"LearnGo/basics/2.5_constructors/writer"
)

func main() {

	//create a new writer object
	w := writer.New("Bobbyimir Peters")
	w.PrintName()

	//updates the name
	w.Update("Dude")
	w.PrintName()

}
