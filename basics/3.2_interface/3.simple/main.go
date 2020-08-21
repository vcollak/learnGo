/*

Tutorial: http://golangtutorials.blogspot.com/2011/06/interfaces-in-go.html

*/

package main

import "fmt"

//Shaper is an interface and has a single function Area that returns an int.
type Shaper interface {
	Area() int
}

//define a Rectangle struct that has a length and a width
type Rectangle1 struct {
	length, width int
}

//write a function Area that can apply to a Rectangle type
func (r Rectangle1) Area() int {
	return r.length * r.width
}

type Square1 struct {
	side int
}

func (sq Square1) Area() int {
	return sq.side * sq.side
}

func main() {

	fmt.Println("### Interfaces7 ###")

	r := Rectangle1{length: 5, width: 3}
	fmt.Println("Rectangle details are: ", r)

	var s Shaper
	s = r //equivalent to "s = Shaper(r)" since Go identifies r matches the Shaper interface
	fmt.Println("Area of Shaper(Rectangle): ", s.Area())
	fmt.Println()

	q := Square1{side: 5}
	fmt.Println("Square details are: ", q)
	s = q //equivalent to "s = Shaper(q)
	fmt.Println("Area of Shaper(Square): ", s.Area())
	fmt.Println()

}
