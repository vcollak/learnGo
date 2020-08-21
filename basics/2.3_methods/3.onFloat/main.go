package main

import (
	"fmt"
	"math"
)

//create an alias to float
type MyFloat float64

//added a method to it
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {

	fmt.Println("### Methods2 ###")

	//created negative float
	f := MyFloat(-math.Sqrt2)

	//printed an ABSed value
	fmt.Println(f.Abs())

}
