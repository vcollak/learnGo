package main

import "fmt"

//function that takes a channel and some number
func foo(c chan int, num int) {

	//sending num to the channel
	c <- num
}

func main() {

	//making a channel
	fooChannel := make(chan int)

	//go routines to call foo and send them numbers
	go foo(fooChannel, 10)
	go foo(fooChannel, 15)

	//receive values from the channels. blocking call so no need to sync
	//only executed after go routines completed
	v1 := <-fooChannel
	v2 := <-fooChannel

	//print hte values
	fmt.Println(v1, v2)

}
