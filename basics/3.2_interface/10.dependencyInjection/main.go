/*
Showcases how to use dependency injection leveraging interfaces.
main.go imports 3 packages:

fancyprinter - concrete implementation of the printer interface ("fancy")
simpleprinter - concrete implementation of the printer interface ("simple"")
printer - printer interface that both concrete packages implement

*/

package main

import (
	"learnGo/basics/3.2_interface/10.dependencyInjection/fancyprinter"
	"learnGo/basics/3.2_interface/10.dependencyInjection/printer"
	"learnGo/basics/3.2_interface/10.dependencyInjection/simpleprinter"
)

func main() {

	//what type of printer to use?
	printerType := "fancy"

	//define printer var as printer interface type
	var printer printer.Printer

	//if fancy use the fancy concrete package
	if printerType == "fancy" {
		printer = fancyprinter.NewPrinter()
	} else {
		printer = simpleprinter.NewPrinter()
	}

	//print (this will execute the method PrintMe() method
	//attached to package
	printer.PrintMe()

}
