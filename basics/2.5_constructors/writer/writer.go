package writer

import "fmt"

//Writer holds the name of the writer
type Writer struct {
	name string
}

//New creates a new writer object
//by creating a new struct and assigning
//value to name variable
func New(name string) *Writer {

	w := Writer{}
	w.name = name
	return &w
}

//Update updates the name of the writer
func (w *Writer) Update(name string) {
	w.name = name
}

//PrintName prints the name of the writer
func (w *Writer) PrintName() {
	fmt.Println("Name:", w.name)
}
