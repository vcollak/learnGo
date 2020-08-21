package main

import (
	"fmt"
)

type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

func main() {

	fmt.Println("### Maps 2 ###")

	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{40.3333, -74.6545}
	fmt.Println(m["Bell Labs"])

}
