package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {

	fmt.Println("### Maps ###")

	//define and instantiate 
	celebs := map[string]int{
		"Nicolas Cage": 50,
		"Selena Gomez": 21,
	}
	fmt.Printf("%#v\n", celebs)

	

	var m map[string]Vertex	//define  
	m = make(map[string]Vertex) //intantiate separately
	m["Bell Labs"] = Vertex{40.2222, -74.222}
	fmt.Println(m["Bell Labs"])

}
