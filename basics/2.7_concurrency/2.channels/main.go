package main

import "fmt"

func main() {

	fmt.Println("### Concurrency ###")

	//makes an channel with buffer size of 2
	c := make(chan int, 2)

	c <- 1 //put a number in the channel
	c <- 2 //put a number in the channel

	fmt.Println(<-c) //consume the channel
	fmt.Println(<-c) //consume the channel

}
