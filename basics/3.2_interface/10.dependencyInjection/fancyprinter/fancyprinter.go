//Package fancyprinter implements a printer interface
package fancyprinter

import (
	"fmt"
	"learnGo/basics/3.2_interface/10.dependencyInjection/printer"
)

//NewPrinter returns the printer interface
//concrete implementation. In this case FancyPrinter
func NewPrinter() printer.Printer {
	return &FancyPrinter{}
}

//FancyPrinter holds a PrintMe method
type FancyPrinter struct{}

//PrintMe prints something to the console
func (f *FancyPrinter) PrintMe() {
	printerType := "fancy"
	fmt.Printf("This is a %s printer", printerType)
}
