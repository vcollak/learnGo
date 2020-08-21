package main

import (
	"fmt"
	"runtime"
	"time"
)

func say1(s string) {
	for i := 0; i < 5; i++ {

		//yield the processor time to allow other go routines to run
		//without this, the program would print world 5 times
		//and then hello 5 times. this enables it to alternate 
		//since it gives some processor time to other go routines 
		runtime.Gosched()

		fmt.Println(s)
	}
}
func main() {

	fmt.Println("### Concurrency2 ###")

	go say1("world") // create a new goroutine
	go say1("hello") // current goroutine
	
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
