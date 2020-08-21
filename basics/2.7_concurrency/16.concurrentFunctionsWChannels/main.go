package main

import (
	"fmt"
	"time"
)

func Counter(c chan int) {

	//loop 5 times and add number to a channel
	//sleep after each iteration
	for i := 0; i <= 5; i++ {
		c <- i
		time.Sleep(time.Millisecond * 500)
	}

	//once done, we don't have anything else to send
	//so close the channel. only senders and never the
	//receivers should be closing the channels
	close(c)
}

func main() {

	//make channel and call the go routine
	c := make(chan int)
	go Counter(c)

	//loop for as long the channel is open
	for {

		//get value from the channel and also if it's still open
		i, open := <-c

		//break out of the loop if no longer open
		if !open {
			break
		}

		fmt.Println("Counter ", i)
	}

}
