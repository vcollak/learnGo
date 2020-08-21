package simpleprinter

import (
	"fmt"
	"learnGo/basics/3.2_interface/10.dependencyInjection/printer"
)

func NewPrinter() printer.Printer {
	return &SimplePrinter{}
}

type SimplePrinter struct{}

func (f *SimplePrinter) PrintMe() {
	fmt.Println("This is a simple printer")
}
