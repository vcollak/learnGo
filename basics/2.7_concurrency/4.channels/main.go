package main

import "fmt"

//pass the slice and channel to the function
func sum(a []int, c chan int) {

	//add up the numbers
	sum := 0
	for _, v := range a {
		sum += v
	}

	//save the value in the channel
	c <- sum // send sum to c
}

func main() {

	fmt.Println("### Concurrency ###")

	//create a slice with numbers
	a := []int{7, 2, 8, -9, 4, 0}

	//make a channel
	c := make(chan int)

	//divide the slice. give the first go routine the
	//first half and the other go routine the other half
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	//get data (sum) from the channels
	x, y := <-c, <-c // receive from c

	//print the individual sums and also sum them
	//together.
	//this is a blocking call so no need for sync to make sure all
	//go routines exited
	fmt.Printf("%d + %d = %d\n",x, y, x+y)
}
