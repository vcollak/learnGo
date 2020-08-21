package main

import (
	"fmt"
	"time"
)

//Bootcamp struct. Anything with first capital letter gets
//exported outside of the package
type Bootcamp struct {

	//capitalize the names to export outside of the package
	Lat  float64
	Lon  float64
	Date time.Time
}

//Point struct
type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}  //type is Point
	q = &Point{1, 2} //type is *Point
	r = Point{X: 1}  //Y:0 implicitly
	s = Point{}      //X:0 and Y:0

)

func main() {

	fmt.Println("### Structs ###")

	fmt.Println(Bootcamp{
		Lat:  43.5,
		Lon:  -118.11,
		Date: time.Now(),
	})

	event := Bootcamp{Lat: 333.33, Lon: 323.233}
	event.Date = time.Now()
	fmt.Printf("Event is on %s, location (%f, %f)", event.Date, event.Lat, event.Lon)
	fmt.Println()
	fmt.Println(p, q, r, s)

}
