package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

//this is a pointer to the Vertex struct
//if not a pointer Vertex would be passed as a value
//and count not modify the vertex struct
func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//this is a pointer to the Vertex struct
//if not a pointer Vertex would be passed as a value
//and cound not modify the vertex struct
func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	v := Vertex2{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())

}
