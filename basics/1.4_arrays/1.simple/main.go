// In Go, an _array_ is a numbered sequence of elements of a
// specific length.
package main

import (
	"fmt"
)

func main() {

	//array with 2 items
	fmt.Println("### array with 2 items ###")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//array of strings
	fmt.Println("### array of strings ###")
	myArray := [2]string{"hello", "world"}
	fmt.Println(myArray)

	//implicit length
	fmt.Println("### implicit length ###")
	a1 := [...]string{"hello", "world"}
	fmt.Println(a1)        // [hello world!]
	fmt.Printf("%s\n", a1) // [hello world!]
	fmt.Printf("%q\n", a1) // ["hello" "world!"]

	//multidimensional arrays
	fmt.Println("### multidimensional arrays ###")
	var ma [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			ma[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Printf("%q", ma)


	//will not compile. using out of bounds
	//var outOfBOunds [2]string
	//a[3] = "Hello"

	//multi dimensional array
	fmt.Println("### multidimensional arrays - another ###")
	doubleArray := [2][4]int{
		[4]int{1, 2, 3, 4},
		[4]int{5, 6, 7, 8},
	}

	fmt.Println(doubleArray)

}
