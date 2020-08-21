package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {

	x, y := 0, 1

	//loop for n
	for i := 0; i < n; i++ {

		//put x in the channel
		c <- x

		//place y into x; place x+y into y
		x, y = y, x+y
	}

	//close the channel once done
	//Note: Only the sender should close a channel,
	//never the receiver. Sending on a closed channel will cause a panic.
	close(c)
}

func main() {

	fmt.Println("### Concurrency3 ###")

	//create a channel with buffer of 10
	c := make(chan int, 10)

	//call function with the size of the channel and channel itself
	go fibonacci(cap(c), c)

	//loop for each value from the channel
	for i := range c {
		fmt.Println(i)
	}
}
