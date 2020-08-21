package main

import "fmt"

func main() {

	const Pi = 3.14
	const (
		StatusOK      = 200
		StatusCreated = 201
	)

	fmt.Println(Pi)
	fmt.Println(StatusOK)
	fmt.Println(StatusCreated)

	//will not compile. cannot change constants
	//Pi = 1

	const (
		a = 1 //assign 1
		b     //assign the same as previous statement
		c     //assign the same as previous statement
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const (
		aa = iota //enum that starts a 0 and counts up
		bb        //same as above +1
		cc        //same as above +1
	)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	const (
		aaa = iota +1  	//enum that starts a 1 and counts up
		bbb        		//same as above +1
		ccc        		//same as above +1
	)

	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)


}
