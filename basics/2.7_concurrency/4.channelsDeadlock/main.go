package main

import "fmt"

func main() {

	//without the buffer (1), the go channel will get a deadlock error
	//that's because we're sending something to the channel, but that's a blocking call
	//the next line that receives is not executed until c <- "hello" is done
	//as a result there is nothing to receive the channel message
	//to fix that, we just add buffering. that's the "1" in channel make
	c := make(chan string, 1)
	c <- "hello"

	msg := <-c
	fmt.Println(msg)

}
