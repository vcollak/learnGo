package main

import "fmt"

//function that returns a function
//the returned function takes int and returns int
//the effect of this is that when the function gets executed 
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	//function gets executed
	//sum is set to 0
	//function is returned
	sum := adder()

	for i := 0; i < 10; i++ {
		//each time the function is called it will 
		//iterate inside of the function return
		//in this case:
		//	sum += x
		//  return sum
		fmt.Println(sum(i))
	}

}
