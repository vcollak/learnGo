package main

import "fmt"

func main() {

	fmt.Println("### Concurrency2 ###")

	//makes an channel with buffer of 2
	c := make(chan int, 2)

	//put things in the channel
	c <- 1
	c <- 2
	c <- 3

	//consume first 2 things form teh channel
	fmt.Println(<-c)
	fmt.Println(<-c)

	//this will cause a deadlock because the channe size is only 2
	//we overfilled the channel buffer
	fmt.Println(<-c)
}
